#!/bin/bash

IP=`curl --connect-timeout 60 ipinfo.io/ip`
./main -msg=$IP
curl -X PUT "https://api.godaddy.com/v1/domains/helowrd.net/records/A/%40" -H "accept: application/json" -H "Content-Type: application/json" -H "Authorization: sso-key Aa6rD6mknSV_NxHXEMhzoAk3My6KWgttFk:NxHZNToscuRBR7ZT3RDZia" -d "[ { \"data\": \"$IP\", \"ttl\": 600}]"

while true; do
    currentIp=`curl --connect-timeout 60 ipinfo.io/ip`
    if [ "$IP" == "$currentIp" ]; then
        echo "IP addr not change."
        echo $IP
        sleep 600
    else
	    if [[ "$currentIp" =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
	        FIELD1=$(echo $currentIp|cut -d. -f1)
	        FIELD2=$(echo $currentIp|cut -d. -f2)
	        FIELD3=$(echo $currentIp|cut -d. -f3)
	        FIELD4=$(echo $currentIp|cut -d. -f4)
	        if [ $FIELD1 -le 255 -a $FIELD2 -le 255 -a $FIELD3 -le 255 -a $FIELD4 -le 255 ]; then
	            IP=$currentIp
	            echo "IP addr changed."
	            echo $IP
	            ./main -msg=$IP
		    curl -X PUT "https://api.godaddy.com/v1/domains/helowrd.net/records/A/%40" -H "accept: application/json" -H "Content-Type: application/json" -H "Authorization: sso-key Aa6rD6mknSV_NxHXEMhzoAk3My6KWgttFk:NxHZNToscuRBR7ZT3RDZia" -d "[ { \"data\": \"$IP\", \"ttl\": 600}]"
	            sleep 600
	        else
	            echo "IP $currentIp not available!"
	        	sleep 60
	        fi
		else
		    echo "IP format error!"
			sleep 60
		fi
    fi
done
