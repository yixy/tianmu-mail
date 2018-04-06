package main

import (mail "github.com/yixy/tianmu-mail"
	"flag"
	"github.com/larspensjo/config"
	"fmt"
)

var (
	configFile = flag.String("conf", "conf.ini", "General configuration file")
	msg = flag.String("msg", "hello,world!", "mail msg content")
)

func main() {
	flag.Parse()

	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		fmt.Println("Fail to find",*configFile,err)
		return
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("smtp") {
		_, err := cfg.SectionOptions("smtp")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	mycontent := *msg
	to,err:=cfg.String("smtp","TO")
	if err != nil {
		fmt.Println(err.Error())
	}
	subject,err:=cfg.String("smtp","SUBJECT")
	if err != nil {
		fmt.Println(err.Error())
	}
	user,err:=cfg.String("smtp","USER")
	if err != nil {
		fmt.Println(err.Error())
	}
	passwd,err:=cfg.String("smtp","PASSWORD")
	if err != nil {
		fmt.Println(err.Error())
	}
	host,err:=cfg.String("smtp","HOST")
	if err != nil {
		fmt.Println(err.Error())
	}
	serverAddr,err:=cfg.String("smtp","SERVER_ADDR")
	if err != nil {
		fmt.Println(err.Error())
	}
	email := mail.NewEmail(to,
		subject, mycontent)
	email.SendEmail(user,passwd,host,serverAddr)
}
