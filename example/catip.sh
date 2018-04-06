#!/bin/bash

IP=`curl ipinfo.io/ip`
./main -msg=$IP

while true; do
    currentIp=`curl ipinfo.io/ip`
    if [ "$IP" == "$currentIp" ]; then
        echo "IP addr not change."
        echo $IP
    else
        IP=$currentIp
        echo "IP addr changed."
        echo $IP
        ./main -msg=$IP
    fi
    sleep 60
done

