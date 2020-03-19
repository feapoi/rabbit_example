package topic
import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError3(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
// 只能在安装 rabbitmq 的服务器上操作
func Recv2() {
	conn, err := amqp.Dial("amqp://guest:guest@47.100.228.3:5672/")
	failOnError3(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError3(err, "Failed to open a channel")
	defer ch.Close()
	err = ch.Qos(1,0,false)   //这样RabbitMQ就会使得每个Consumer在同一个时间点最多处理一个Message。换句话说，在接收到该Consumer的ack前，他它不会将新的Message分发给它。
	q, err := ch.QueueDeclare(
		"topic_info", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError3(err, "Failed to declare a queue")

	//绑定交换器和队列
	err = ch.QueueBind(q.Name, "*.orange.*", "topic_logs", false, nil)
	failOnError3(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack 收到消息即返回ACK删除，不管接收方有没有处理完
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError3(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}