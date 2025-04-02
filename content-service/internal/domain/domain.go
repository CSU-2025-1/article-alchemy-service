package domain

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
)

type ContentService interface {
	ExtractContentPreview(ctx context.Context, dto model.ExtractContentPreviewDTO) error
	ExtractContent(ctx context.Context, dto model.ExtractContentDTO) (uint64, error)
	GetContentsByUserId(ctx context.Context, userId uint64) ([]model.Content, error)
}
