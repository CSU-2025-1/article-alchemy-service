package worker

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"article-alchemy-service/internal/rabbitmq"
	"article-alchemy-service/internal/service"
	"article-alchemy-service/pkg/models"
)

func StartWorker() {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	exchange := os.Getenv("RABBITMQ_REQUEST_QUEUE")
	queue := os.Getenv("RABBITMQ_REQUEST_QUEUE")

	consumer, publisher, err := rabbitmq.NewRabbitMQ(rabbitURL, exchange, queue)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer consumer.Close()

	msgs, err := consumer.Consume()
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	log.Println("Worker is running, waiting for messages...")

	for msg := range msgs {
		var event models.WorkerEvent
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Printf("JSON deserialization error: %v", err)
			continue
		}

		go processEvent(event, publisher)
	}
}

func processEvent(event models.WorkerEvent, publisher *rabbitmq.Publisher) {
	responseQueue := os.Getenv("RABBITMQ_RESPONSE_QUEUE")
	summary, err := service.GetSummary(event.URL)

	responseEvent := models.WorkerEvent{
		ContentID: event.ContentID,
		URL:       event.URL,
		Email:     event.Email,
	}

	if err != nil {
		log.Printf("Error getting summary: %v", err)
		responseEvent.Error = err.Error()
	} else {
		responseEvent.Body = string(summary)
	}

	_ = publisher.Publish(context.Background(), responseQueue, responseEvent)
}
