package rabbitmq

import (
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   string
}

type Publisher struct {
	channel  *amqp091.Channel
	exchange string
}

func NewRabbitMQ(url, exchange, queue string) (*Consumer, *Publisher, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	err = ch.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
	)
	if err != nil {
		return nil, nil, err
	}

	err = ch.QueueBind(
		queue,
		queue,
		exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	err = ch.QueueBind(
		os.Getenv("RABBITMQ_RESPONSE_QUEUE"),
		os.Getenv("RABBITMQ_RESPONSE_QUEUE"),
		exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return &Consumer{conn: conn, channel: ch, queue: queue}, &Publisher{channel: ch, exchange: exchange}, nil
}

func (c *Consumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
