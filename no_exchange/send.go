package no_exchange

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
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := "error"
	err = ch.Publish(
		"",     // exchange
		"", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,    //消息持久化
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
