package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "./my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		MustExist: false,
		Poll: true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var (
		msg *tail.Line
		ok bool
	)

	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		fmt.Println("msg: ", msg)
	}
}