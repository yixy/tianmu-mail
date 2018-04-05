package mail

import (
	"net/smtp"
	"strings"
	"fmt"
)

type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func (email Email)SendEmail(user,passwd,host,serverAddr string) (err error) {
	auth := smtp.PlainAuth("", user, passwd, host)
	sendTo := strings.Split(email.to, ";")

	for _, v := range sendTo {
		str := strings.Replace("From: "+user+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg
		err = smtp.SendMail(
			serverAddr,
			auth,
			user,
			[]string{v},
			[]byte(str),
		)
		if err!=nil{
			fmt.Println(err.Error())
		}
	}

	return err
}
