package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	// 使用的是生产者模块，向kafka发送一些日志信息
	// 初始化一些配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区随机
	config.Producer.Return.Successes = true // 需要返回success，true

	// 实例化一个生产者对象
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err: ", err)
		return
	}

	defer client.Close()

	msg := &sarama.ProducerMessage{
		Topic: "nginx_log", // topic如果不存在，会自动创建
		Value: sarama.StringEncoder("this is a good test, my message is good"),
	}

	for {
		// pid: 分区id
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed, ", err)
			return
		}

		fmt.Printf("pid:%v, offset:%v\n", pid, offset)

		time.Sleep(time.Second)
	}
}