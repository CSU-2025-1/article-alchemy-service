package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
	"github.com/rabbitmq/amqp091-go"
)

type EventPublisher interface {
	Publish(ctx context.Context, routingKey string, event model.WorkerEvent) error
}

type Publisher struct {
	channel  *amqp091.Channel
	exchange string
}

func NewPublisher(ch *amqp091.Channel, exchange string) *Publisher {
	return &Publisher{
		channel:  ch,
		exchange: exchange,
	}
}

func (p *Publisher) Publish(ctx context.Context, routingKey string, event model.WorkerEvent) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.channel.Publish(p.exchange, routingKey, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        bytes,
	})
}
