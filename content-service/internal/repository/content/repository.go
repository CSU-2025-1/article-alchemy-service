package content

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/content"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContentRepository struct {
	pool *pgxpool.Pool
}

func NewContentRepository(pool *pgxpool.Pool) *ContentRepository {
	return &ContentRepository{pool}
}

func (c *ContentRepository) Create(ctx context.Context, model content.Content) (uint64, error) {
	sql := `INSERT INTO public.contents (user_id, status, data, error, created_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING contents.content_id
		   `

	row := c.pool.QueryRow(
		ctx,
		sql,
		model.UserID,
		model.Status,
		model.Data,
		model.Error,
		model.CreatedAt)

	var contentId uint64
	if err := row.Scan(&contentId); err != nil {
		return 0, err
	}

	return contentId, nil
}

func (c *ContentRepository) Update(ctx context.Context, model content.Content) (uint64, error) {
	panic("implement me")
}

func (c *ContentRepository) GetContentsByUserId(ctx context.Context, userId uint64) []content.Content {
	panic("implement me")
}
