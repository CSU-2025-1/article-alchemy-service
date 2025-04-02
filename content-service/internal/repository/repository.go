package repository

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/entity"
)

type ContentRepository interface {
	Create(ctx context.Context, entity entity.Content) (uint64, error)
	Update(ctx context.Context, entity entity.Content) (uint64, error)
	GetContentsByUserId(ctx context.Context, userId uint64) ([]entity.Content, error)
}
