package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func (c *Consumer) Consume() (<-chan amqp091.Delivery, error) {
	return c.channel.Consume(
		c.queue, "", true, false, false, false, nil,
	)
}
