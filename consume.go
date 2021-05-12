package rabbitmq

import "github.com/streadway/amqp"

func (q *RabbitMQ) Consume() <- chan amqp.Delivery {
	c, err := q.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		panic(err)
	}
	return c
}
