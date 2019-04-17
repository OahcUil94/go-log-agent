# go实现日志收集系统

## tail库

- ReOpen参数：等同于`tail -F`，根据文件名进行追踪，并保持重试，即使该文件被删除或改名后，如果再次创建相同的文件名文件，会继续追踪。
- MustExist参数：如果设置为true，则表示如果文件不存在，就会退出。
- Poll参数：表示使用轮询的方式来监听文件的修改


## kafka的相关操作

### consumer操作

消费者监听生产者发送的消息：

```bash
kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic nginx_log
```

> 注意：--bootstrap-server参数不能少 

### topic操作

查看可用的topic列表：

```bash
kafka-topics --list --zookeeper 127.0.0.1:2181
```

> 注意：--zookeeper参数不能少

删除topic：

```bash
kafka-topics --delete --zookeeper 127.0.0.1:2181 --topic nginx-log
```