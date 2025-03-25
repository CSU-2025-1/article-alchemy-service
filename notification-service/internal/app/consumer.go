package app

import (
    "encoding/json"
    "log"
    "github.com/streadway/amqp"
    "github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/domain/notification"
    "github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/email"
)

type Consumer struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    queue    amqp.Queue
}

func NewConsumer(amqpURL, queueName string) (*Consumer, error) {
    conn, err := amqp.Dial(amqpURL)
    if err != nil {
        return nil, err
    }
    
    channel, err := conn.Channel()
    if err != nil {
        return nil, err
    }
    
    queue, err := channel.QueueDeclare(
        queueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }
    
    return &Consumer{
        conn:    conn,
        channel: channel,
        queue:   queue,
    }, nil
}

func (c *Consumer) StartConsuming() error {
    msgs, err := c.channel.Consume(
        c.queue.Name, // queue
        "",           // consumer
        false,       // auto-ack
        false,       // exclusive
        false,       // no-local
        false,       // no-wait
        nil,         // args
    )
    if err != nil {
        return err
    }
    
    forever := make(chan bool)
    
    go func() {
        for d := range msgs {
            var notification notification.EmailNotification
            if err := json.Unmarshal(d.Body, &notification); err != nil {
                log.Printf("Error decoding message: %v", err)
                _ = d.Nack(false, false) 
                continue
            }
            
            // Отправка email
            if err := email.Send(notification.Email, notification.URL, notification.Data); err != nil {
                log.Printf("Failed to send email: %v", err)
                _ = d.Nack(false, true)
                continue
            }
            
            if err := d.Ack(false); err != nil {
                log.Printf("Failed to ack message: %v", err)
            }
            log.Printf("Email sent to %s", notification.Email)
        }
    }()
    
    <-forever
    return nil
}

func (c *Consumer) Close() {
    _ = c.channel.Close()
    _ = c.conn.Close()
}