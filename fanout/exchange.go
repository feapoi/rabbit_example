package fanout

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError2(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Exchange() {
	conn, err := amqp.Dial("amqp://guest:guest@47.100.228.3:5672/")
	failOnError2(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError2(err, "Failed to open a channel")
	defer ch.Close()

	//定义交换器
	err = ch.ExchangeDeclare("fanout_logs", amqp.ExchangeFanout, true, false, false, false, nil)

	failOnError2(err, "Failed to declare an exchange")
}
