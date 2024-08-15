package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	"home-service/pkg/config"
)

type DbPostgres struct {
	db *pgxpool.Pool
	log	*slog.Logger
}

func New(ctx context.Context, cfg *config.DBConfig, log *slog.Logger) (*DbPostgres, error) {
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
				db:	db,
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

	// if err = pgInstance.RunMigrations(ctx); err != nil {
	// 	return nil, fmt.Errorf("fail to run migrations: %w", err)
	// }

	return pgInstance, nil
}

func createConnString(cfg *config.DBConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}

func (pg *DbPostgres) RunMigrations(ctx context.Context) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS message (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		read BOOLEAN NOT NULL DEFAULT false
	);
	`
	_, err := pg.db.Exec(ctx, createTableQuery)
	if err != nil {
		return fmt.Errorf("unable to run migrations: %w", err)
	}
	return nil
}

func (pg *DbPostgres) MarkAsRead(ctx context.Context, id string) error {
	query := `UPDATE message SET read = true WHERE id = $1`
	_, err := pg.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("unable to mark message as read: %w", err)
	}
	return nil
}
