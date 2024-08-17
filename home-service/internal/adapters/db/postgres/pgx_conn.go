package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"

	"home-service/pkg/config"
)

type DbPostgres struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func New(ctx context.Context, cfg *config.PostgresConfig, log *slog.Logger) (*DbPostgres, error) {
	op := "adapters.db.postgres.New"

	var err error
	connString := createConnString(cfg)

	var pgInstance *DbPostgres
	var pgOnce sync.Once

	pgOnce.Do(func() {
		var db *pgxpool.Pool
		db, err = pgxpool.New(ctx, connString)
		if err == nil {
			pgInstance = &DbPostgres{
				db:  db,
				log: log,
			}
		}
	})

	if err != nil {
		return nil, fmt.Errorf("fail to create connection: %w", err)
	}

	err = pgInstance.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to connect db: %w", err)
	}

	log.With(
		slog.String("where", op),
	)
	log.Info("database connect")


    if err := goose.SetDialect("postgres"); err != nil {
		fmt.Errorf("%s: %w", op, err)
       panic(err)
    }
    db := stdlib.OpenDBFromPool(pgInstance.db)
    if err := goose.Up(db, "migrations"); err != nil {
		fmt.Errorf("%s: %w", op, err)
        panic(err)
    
	}
    if err := db.Close(); err != nil {
		fmt.Errorf("%s: %w", op, err)
        panic(err)
    }


	return pgInstance, nil
}



func createConnString(cfg *config.PostgresConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}