package database

import (
	"context"
	"log"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	Instance *Postgres
)

func NewPG(ctx context.Context, connString string) (*Postgres, error) {
	log.Print("Create db connections")
	var pgOnce sync.Once
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			log.Printf("unable to create connection pool: %s", err)
		}
		Instance = &Postgres{db}
	})

	return Instance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}

func Migratedb(url string, filePath string) {
	m, err := migrate.New(
		"file://"+filePath,
		url+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		panic(err)
	}
}
