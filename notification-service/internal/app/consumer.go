package app

import (
    "encoding/json"
    "log"
    "github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/domain/notification"
    "github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/email"
    "fmt"

    amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    queue    amqp.Queue
}

func NewConsumer(amqpURL, queueName string) (*Consumer, error) {
    conn, err := amqp.Dial(amqpURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
    }
    
    channel, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, fmt.Errorf("failed to open channel: %w", err)
    }
    
    queue, err := channel.QueueDeclare(
        queueName,
        true,  
        false, 
        false,   
        false, 
        nil,    
    )
    if err != nil {
        return nil, err
    }
    log.Println("Установлено имя очереди - " + queueName)
    
    return &Consumer{
        conn:    conn,
        channel: channel,
        queue:   queue,
    }, nil
}

func (c *Consumer) StartConsuming() error {
    msgs, err := c.channel.Consume(
        c.queue.Name,
        "", 
        false,  
        false,   
        false,    
        false,  
        nil, 
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
            if err := email.Send(notification.Email, notification.URL, notification.Body, notification.Error); err != nil {
                log.Printf("Failed to send email: %v (ContentID: %d)", err, notification.ContentID)
                _ = d.Nack(false, false)
                continue
            }
            
            if err := d.Ack(false); err != nil {
                log.Printf("Failed to ack message: %v", err)
            }
            log.Printf("Email sent for %s (ContentID: %d)", notification.Email, notification.ContentID)
        }
    }()
    
    <-forever
    return nil
}

func (c *Consumer) Close() {
    _ = c.channel.Close()
    _ = c.conn.Close()
}