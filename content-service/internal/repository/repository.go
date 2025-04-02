package repository

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/content"
)

type ContentRepository interface {
	Create(ctx context.Context, entity content.Content) (uint64, error)
	Update(ctx context.Context, entity content.Content) (uint64, error)
	GetContentsByUserId(ctx context.Context, userId uint64) []content.Content
}
