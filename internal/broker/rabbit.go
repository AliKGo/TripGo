package broker

import amqp "github.com/rabbitmq/amqp091-go"

type Broker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (b *Broker) Publish(exchange, routingKey string, body []byte) error {
	return b.ch.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
