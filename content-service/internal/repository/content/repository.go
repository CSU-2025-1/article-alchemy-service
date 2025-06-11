package content

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContentRepository struct {
	pool *pgxpool.Pool
}

func NewContentRepository(pool *pgxpool.Pool) *ContentRepository {
	return &ContentRepository{pool}
}

func (c *ContentRepository) Create(ctx context.Context, model model.Content) (uint64, error) {
	sql := `INSERT INTO public.contents (user_id, status, url, data, error, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING contents.content_id`

	row := c.pool.QueryRow(
		ctx,
		sql,
		model.UserID,
		model.Status,
		model.URL,
		model.Data,
		model.Error,
		model.CreatedAt)

	var contentId uint64
	if err := row.Scan(&contentId); err != nil {
		return 0, err
	}

	return contentId, nil
}

func (c *ContentRepository) Update(ctx context.Context, model model.Content) error {
	sql := `UPDATE public.contents SET user_id = $1, status = $2, data = $3, error = $4 WHERE content_id = $5 RETURNING content_id`

	_, err := c.pool.Exec(ctx, sql, model.UserID, model.Status, model.Data, model.Error, model.ContentID)

	return err
}

func (c *ContentRepository) GetById(ctx context.Context, contentId uint64) (model.Content, error) {
	sql := `SELECT * FROM public.contents WHERE content_id = $1`

	row := c.pool.QueryRow(ctx, sql, contentId)

	var content model.Content
	err := row.Scan(
		&content.ContentID,
		&content.UserID,
		&content.Status,
		&content.URL,
		&content.Data,
		&content.Error,
		&content.CreatedAt)

	return content, err
}

func (c *ContentRepository) GetContentsByUserId(ctx context.Context, userId uint64) ([]model.Content, error) {
	sql := `SELECT * FROM public.contents WHERE user_id = $1`

	rows, err := c.pool.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}

	content, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Content])
	if err != nil {
		return nil, err
	}

	return content, nil
}
