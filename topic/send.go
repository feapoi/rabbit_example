package topic

import (
	"github.com/streadway/amqp"
	"log"
)
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 只能在安装 rabbitmq 的服务器上操作
func Send() {
	conn, err := amqp.Dial("amqp://guest:guest@47.100.228.3:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	bodylist := []string{"quick.orange.rabbit",
		"lazy.orange.elephant",
		"quick.orange.fox",
		"lazy.brown.fox",
		"quick.brown.fox",
		"quick.orange.male.rabbit",
		"lazy.orange.male.rabbit"}
	for _, v := range bodylist {
		err = ch.Publish(
			"topic_logs",     // exchange
			v, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,    //消息持久化
				ContentType: "text/plain",
				Body:        []byte(v),
			})
		log.Printf(" [x] Sent %s", v)
		failOnError(err, "Failed to publish a message")
	}
}
