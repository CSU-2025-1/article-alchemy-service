package content

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/repository"
)

type Service struct {
	repository repository.ContentRepository
}

func NewService(repository repository.ContentRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ExtractContentPreview(ctx context.Context, dto ExtractContentPreviewDTO) {

}

func (s *Service) ExtractContent(ctx context.Context, dto ExtractContentDTO) {

}

func (s *Service) GetContentsByUserId(ctx context.Context, userId uint64) {

}
