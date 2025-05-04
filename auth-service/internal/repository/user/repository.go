package user

import (
	"context"
	"errors"
	domainErr "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/errors"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool}
}

func (u *UserRepository) Create(ctx context.Context, usr user.User) (uint64, error) {
	sql := `INSERT INTO public.users (username, email, is_active, password_hash, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`

	row := u.pool.QueryRow(
		ctx,
		sql,
		usr.Username,
		usr.Email,
		usr.IsActive,
		usr.PasswordHash,
		usr.CreatedAt)

	var userID uint64
	if err := row.Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *UserRepository) GetById(ctx context.Context, userId uint64) (user.User, error) {
	sql := `SELECT * FROM public.users WHERE user_id = $1`

	row := u.pool.QueryRow(ctx, sql, userId)

	var usr user.User
	err := row.Scan(
		&usr.UserID,
		&usr.Username,
		&usr.Email,
		&usr.IsActive,
		&usr.PasswordHash,
		&usr.CreatedAt)

	if err != nil {
		return usr, err
	}

	return usr, nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (user.User, error) {
	sql := `SELECT * FROM public.users WHERE email = $1`

	row := u.pool.QueryRow(ctx, sql, email)

	var usr user.User
	err := row.Scan(
		&usr.UserID,
		&usr.Username,
		&usr.Email,
		&usr.IsActive,
		&usr.PasswordHash,
		&usr.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return usr, domainErr.ErrUserNotFound
		}
		return usr, err
	}

	return usr, nil
}
