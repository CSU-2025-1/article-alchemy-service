package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"time"
)

const (
	maxRetries = 3
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Close()
}

type pgxClient struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context, dsn string) Client {
	for i := 0; i < maxRetries; i++ {
		pool, err := pgxpool.New(ctx, dsn)
		if err != nil {
			slog.Error("failed to connect to the database", "retry_count", i+1, "error", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if err = pool.Ping(ctx); err != nil {
			pool.Close()
			slog.Error("failed to ping database, retrying...", "retry_count", i+1, "error", err)
			time.Sleep(3 * time.Second)
			continue
		}

		pool.Close()

		return &pgxClient{pool: pool}
	}

	log.Fatalln(fmt.Errorf("failed to connect to the database after %d retries", maxRetries))
	return nil
}

func (p *pgxClient) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return p.pool.Exec(ctx, sql, arguments...)
}

func (p *pgxClient) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.pool.Query(ctx, sql, args...)
}

func (p *pgxClient) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *pgxClient) Close() {
	p.pool.Close()
}
