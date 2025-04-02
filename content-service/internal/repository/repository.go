package repository

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
)

type ContentRepository interface {
	Create(ctx context.Context, entity model.Content) (uint64, error)
	Update(ctx context.Context, entity model.Content) error
	GetById(ctx context.Context, contentId uint64) (model.Content, error)
	GetContentsByUserId(ctx context.Context, userId uint64) ([]model.Content, error)
}
