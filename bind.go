package rabbitmq

func (q *RabbitMQ) Bind(exchange string)  {
	err := q.channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil,
		)
	if err != nil {
		panic(err)
	}
	q.exchange = exchange

}
