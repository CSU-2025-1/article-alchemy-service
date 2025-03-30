package worker

import (
	"encoding/json"
	"log"
	"os"

	"article-alchemy-service/internal/service"
	"article-alchemy-service/pkg/models"

	"github.com/streadway/amqp"
)

func StartWorker() {
	log.Printf("RABBITMQ_URL: %s", os.Getenv("RABBITMQ_URL"))
	rabbitURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	queueName := os.Getenv("RABBITMQ_QUEUE")
	responseQueueName := os.Getenv("RABBITMQ_RESPONSE_QUEUE")

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", queueName, err)
	}

	_, err = ch.QueueDeclare(
		responseQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare response queue %s: %v", responseQueueName, err)
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Unable to subscribe to queue: %v", err)
	}

	log.Println("Worker is running, waiting for messages...")

	for msg := range msgs {
		var event models.Event
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Printf("JSON deserialization error: %v", err)
			continue
		}

		go processEvent(event, ch)
	}
}

func processEvent(event models.Event, ch *amqp.Channel) {
	summary, err := service.GetSummary(event.URL)
	if err != nil {
		log.Printf("Error getting summary: %v", err)
		return
	}

	responseEvent := models.EventResponse{
		UserID: event.UserID,
		URL:    event.URL,
		Email:  event.Email,
		Data:   summary,
	}

	sendToQueue(responseEvent, ch)
}

func sendToQueue(response models.EventResponse, ch *amqp.Channel) {
	queueName := os.Getenv("RABBITMQ_RESPONSE_QUEUE")

	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare response queue before sending: %v", err)
		return
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Printf("Response serialization error: %v", err)
		return
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Error sending to RabbitMQ: %v", err)
	} else {
		log.Printf("✅ The answer has been sent to the queue %s", queueName)
	}
}
