package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel *amqp.Channel
	Name string
	exchange string
}

func New(endPoint string) *RabbitMQ {
	conn, err := amqp.Dial(endPoint)
	if err != nil {
		fmt.Println("连接rabbitmq失败,err is", err)
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
		)
	if err != nil {
		panic(err)
	}

	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name
	return mq
}
