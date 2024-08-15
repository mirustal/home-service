package postgres

import (
	"context"
	"fmt"

	"home-service/internal/models"
)


func (pg *DbPostgres) Close() {
	pg.db.Close()
}

func (pg *DbPostgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}


func (pg *DbPostgres) Insert(ctx context.Context, text string) (string, error) {
	query := `INSERT INTO message (text) VALUES ($1) RETURNING id`
	var id string
	err := pg.db.QueryRow(ctx, query, text).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("unable to insert row: %w", err)
	}
	return id, nil
}


func (pg *DbPostgres) GetUser(ctx context.Context, email string) (models.User, error) {
	return models.User{}, nil
}


func (pg *DbPostgres) IsAdmin(ctx context.Context, id int64) (bool, error) {
	return false, nil
}

func (pg *DbPostgres) SaveUser(ctx context.Context, email string, hashPass []byte) (uid int64, err error) {
	return 0, nil
}
