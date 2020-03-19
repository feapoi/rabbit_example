package main

import (
	_ "rabbit_example/fanout"
	"rabbit_example/topic"
)

func main() {
	topic.Exchange()
	go topic.Recv()
	go topic.Recv2()
	topic.Send()
	forever := make(chan bool)
	<-forever
}
