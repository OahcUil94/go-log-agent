package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./log-agent.conf")
	if err != nil {
		fmt.Println("new config failed, err: ", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err: ", err)
		return
	}

	fmt.Println("Port: ", port)
	log_level := conf.String("logs::log_level")
	//if err != nil {
	//	fmt.Println("read log_level failed, ", err)
	//	return
	//}

	fmt.Println("log_level: ", log_level)

	log_path := conf.String("logs::log_path")
	fmt.Println("log_path: ", log_path)
}
