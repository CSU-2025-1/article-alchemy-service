package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func NewRabbitMQ(url string, exchange string, queue ...string) (*amqp091.Connection, *amqp091.Channel) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ exchange: %s", err)
	}

	for _, item := range queue {
		q, err := ch.QueueDeclare(
			item,
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Fatalf("Failed to declare RabbitMQ queue: %s", err)
		}

		err = ch.QueueBind(
			q.Name,
			item,
			exchange,
			false,
			nil,
		)

		if err != nil {
			log.Fatalf("Failed to bind RabbitMQ queue: %s", err)
		}
	}

	return conn, ch
}
