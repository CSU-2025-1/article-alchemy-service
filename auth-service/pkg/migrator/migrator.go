package migrator

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

func Migrate(pool *pgxpool.Pool) {
	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln("error applying migrations", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalln("error applying migrations", err)
	}
}
