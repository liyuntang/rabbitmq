package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

func (q *RabbitMQ) Publish(exchange string, body interface{})  {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	err = q.channel.Publish(exchange,
		"",
		false,
		false,
		amqp.Publishing{
		ReplyTo: q.Name,
		Body: []byte(str),
		})
	if err != nil {
		panic(err)
	}
}
