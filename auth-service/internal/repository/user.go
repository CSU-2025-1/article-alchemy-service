package repository

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool}
}

func (u *UserRepository) Create(ctx context.Context, user entity.User) (uint64, error) {
	sql := `INSERT INTO public.users (username, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING user_id`

	row := u.pool.QueryRow(
		ctx,
		sql,
		user.Username,
		user.Email,
		user.Password)

	var userID uint64
	if err := row.Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *UserRepository) GetById(ctx context.Context, userID uint64) (entity.User, error) {
	sql := `SELECT * FROM public.users WHERE user_id = $1`

	row := u.pool.QueryRow(ctx, sql, userID)

	var user entity.User
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
