package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/delivery/rabbitmq"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/repository"
	"log"
	"log/slog"
	"time"
)

type Service struct {
	repository repository.ContentRepository
	publisher  rabbitmq.EventPublisher
	consumer   rabbitmq.EventConsumer
}

func NewService(repository repository.ContentRepository, publisher rabbitmq.EventPublisher, consumer rabbitmq.EventConsumer) *Service {
	return &Service{
		repository: repository,
		publisher:  publisher,
		consumer:   consumer,
	}
}

func (s *Service) ExtractContentPreview(ctx context.Context, dto model.ExtractContentPreviewDTO) error {
	event := model.WorkerEvent{
		ContentID: 0,
		URL:       dto.Url,
		Email:     dto.Email,
		Body:      "",
		Error:     "",
	}

	if err := s.publisher.Publish(ctx, "request_content_parsing", event); err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}

func (s *Service) ExtractContent(ctx context.Context, dto model.ExtractContentDTO) (uint64, error) {
	contentId, err := s.repository.Create(ctx, model.Content{
		UserID:    dto.UserID,
		Status:    model.StatusPending,
		URL:       dto.Url,
		Data:      "",
		Error:     "",
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		return 0, fmt.Errorf("failted to create content: %w", err)
	}

	event := model.WorkerEvent{
		ContentID: contentId,
		URL:       dto.Url,
		Email:     dto.Email,
		Body:      "",
		Error:     "",
	}

	if err = s.publisher.Publish(ctx, "request_content_parsing", event); err != nil {
		return 0, fmt.Errorf("failted to publish event: %w", err)
	}
	return contentId, nil
}

// StartSomeShitListener TODO: need to make limit for goroutines and cancel context
func (s *Service) StartSomeShitListener(ctx context.Context) {
	messages, err := s.consumer.Consume("response_content_parsing")
	if err != nil {
		log.Fatalf("failed to start consumer for response_content_parsing: %v", err)
	}

	for msg := range messages {
		go func(m []byte) {
			var event model.WorkerEvent
			if err = json.Unmarshal(m, &event); err != nil {
				slog.Warn("failed to unmarshal event_data",
					slog.Any("event_data", event),
					slog.Any("error", err),
				)
				return
			}

			if event.ContentID == 0 {
				if err = s.publisher.Publish(ctx, "notification", event); err != nil {
					slog.Warn("failed to publish event to notification queue",
						slog.Any("event_data", event),
						slog.Any("error", err),
					)
				}
				return
			}

			content, err := s.repository.GetById(ctx, event.ContentID)
			if err != nil {
				slog.Warn("failed to get content",
					slog.Uint64("content_id", event.ContentID),
					slog.Any("event_data", event),
					slog.Any("error", err),
				)
				return
			}

			if len(event.Error) != 0 {
				content.Status = model.StatusFailed
				content.Error = event.Error

			} else {
				content.Status = model.StatusCompleted
				content.Data = event.Body
			}

			if err = s.repository.Update(ctx, content); err != nil {
				slog.Warn("failed to update content",
					slog.Uint64("content_id", event.ContentID),
					slog.Any("event_data", event),
					slog.Any("error", err),
				)
				return
			}

			//TODO: should i add sending email for auth user?
		}(msg.Body)
	}
}

func (s *Service) GetContentsByUserId(ctx context.Context, userId uint64) ([]model.Content, error) {
	return s.repository.GetContentsByUserId(ctx, userId)
}
