package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"article-alchemy-service/internal/rabbitmq"
	"article-alchemy-service/internal/service"
	"article-alchemy-service/pkg/models"
)

func StartWorker() {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	exchange := os.Getenv("RABBITMQ_EXCHANGE")
	queue := os.Getenv("RABBITMQ_REQUEST_QUEUE")

	consumer, publisher, err := rabbitmq.NewRabbitMQ(rabbitURL, exchange, queue)

	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer consumer.Close()

	var event models.WorkerEvent
	msgs, err := consumer.Consume()
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to consume messages: %v", err)
		log.Print(errorMsg)
		sendError(publisher, errorMsg, event)
		return
	}

	log.Println("Worker is running, waiting for messages...")

	for msg := range msgs {
		var event models.WorkerEvent
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			errorMsg := fmt.Sprintf("JSON deserialization error: %v", err)
			log.Print(errorMsg)
			sendError(publisher, errorMsg, event)
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
		errorMsg := fmt.Sprintf("Error getting summary: %v", err)
		log.Print(errorMsg)
		sendError(publisher, errorMsg, responseEvent)
		return
	}

	responseEvent.Body = string(summary)

	_ = publisher.Publish(context.Background(), responseQueue, responseEvent)
}

func sendError(publisher *rabbitmq.Publisher, err string, event models.WorkerEvent) {
	responseQueue := os.Getenv("RABBITMQ_RESPONSE_QUEUE")
	responseEvent := models.WorkerEvent{
		ContentID: event.ContentID,
		URL:       event.URL,
		Email:     event.Email,
		Body:      "",
		Error:     err,
	}
	_ = publisher.Publish(context.Background(), responseQueue, responseEvent)
}
