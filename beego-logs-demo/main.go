package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logs/logcollect.log"
	config["level"] = logs.LevelDebug

	buf, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err: ", err)
		return
	}

	if err := logs.SetLogger(logs.AdapterFile, string(buf)); err != nil {
		fmt.Println("set logger failed, err: ", err)
		return
	}

	logs.Debug(`this is a test, my name is %s`, "stu01")
	logs.Trace(`this is a test, my name is %s`, "stu02")
	logs.Warn(`this is a test, my name is %s`, "stu03")
}

