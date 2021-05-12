package rabbitmq

func (q *RabbitMQ) Close()  {
	q.channel.Close()
}
