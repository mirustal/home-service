package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"

	dbErr "auth-service/internal/adapters/db"
	"auth-service/internal/models"
)

func (pg *DbPostgres) Close() {
	pg.db.Close()
}

func (pg *DbPostgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *DbPostgres) GetUser(ctx context.Context, userID string) (models.User, error) {
	const op = "postgres.GetUser"

	query := `
        SELECT users.id,
               users.email,
               users.password_hashed,
               users.user_type,
               users.created_at,
               users.updated_at
        FROM users
        WHERE id = $1
        LIMIT 1;
    `
	var user models.User

	err := pg.db.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.HashPass,
		&user.Type,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return user, dbErr.ErrUserNotFound
	} else if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (pg *DbPostgres) SaveUser(ctx context.Context, email string, passHash []byte, userType string) (string, error) {
	const op = "postgres.SaveUser"
	var id uuid.UUID

	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	query := `
        INSERT INTO users (email, password_hashed, user_type) 
        VALUES ($1, $2, $3) 
        RETURNING id;
    `
	err = tx.QueryRow(ctx, query, email, passHash, userType).Scan(&id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
			return "", dbErr.ErrUserExists
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err = tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id.String(), nil
}

func (pg *DbPostgres) SaveRefreshToken(ctx context.Context, refreshToken string, uid string) error {
	const op = "postgres.SaveRefreshToken"


	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	query := `
        UPDATE refresh_tokens
        SET is_valid = false
        WHERE user_id = $1;
    `
	_, err = tx.Exec(ctx, query, uid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	query = `
        INSERT INTO refresh_tokens (token, user_id)
        VALUES ($1, $2);
    `
	_, err = tx.Exec(ctx, query, refreshToken, uid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *DbPostgres) App(ctx context.Context, appID int32) (models.App, error) {
	const op = "postgres.App"

	query := `
        SELECT apps.name, apps.secret, apps.id
        FROM apps
        WHERE id = $1
        LIMIT 1;
    `
	var app models.App

	err := pg.db.QueryRow(ctx, query, appID).Scan(&app.Name, &app.Secret, &app.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return app, dbErr.ErrAppNotFound
		}
		return app, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
