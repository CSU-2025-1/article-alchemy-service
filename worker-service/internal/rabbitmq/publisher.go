package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

func (p *Publisher) GetExchange() string {
	return p.exchange
}

func (p *Publisher) Publish(ctx context.Context, routingKey string, event interface{}) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(ctx, p.exchange, routingKey, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
