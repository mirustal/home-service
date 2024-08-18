package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	dbErr "home-service/internal/adapters/db"
	"home-service/internal/models"
)

func (pg *DbPostgres) Close() {
	pg.db.Close()
}

func (pg *DbPostgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *DbPostgres) CreateHouse(ctx context.Context, address string, year int, developer string) (int, error) {
	const op = "postgres.CreateHouse"

	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to begin tx: %w", op, err)
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO houses (address, year, developer, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id;
	`

	var houseID int
	err = tx.QueryRow(ctx, query, address, year, developer).Scan(&houseID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("%s: failed to commit tx: %w", op, err)
	}

	return houseID, nil
}

func (pg *DbPostgres) CreateFlat(ctx context.Context, houseID, price, rooms int) (int, error) {
	const op = "postgres.CreateFlat"

	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to begin tx: %w", op, err)
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO flats (house_id, price, rooms, status, created_at, updated_at)
		VALUES ($1, $2, $3, 'created', NOW(), NOW())
		RETURNING id;
	`

	var flatID int
	err = tx.QueryRow(ctx, query, houseID, price, rooms).Scan(&flatID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	updateHouseQuery := `
		UPDATE houses
		SET updated_at = NOW()
		WHERE id = $1;
	`

	_, err = tx.Exec(ctx, updateHouseQuery, houseID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("%s: failed to commit tx: %w", op, err)
	}

	return flatID, nil
}

func (pg *DbPostgres) UpdateFlatStatus(ctx context.Context, flatID int, status string) (models.Flat, error) {
	const op = "postgres.UpdateFlatStatus"

	var flat models.Flat

	tx, err := pg.db.Begin(ctx)
	if err != nil {
		pg.log.Error("%s: %v", op, err)
		return flat, fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE flats
		SET status = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, house_id, price, rooms, status, created_at, updated_at;
	`

	err = tx.QueryRow(ctx, query, status, flatID).Scan(
		&flat.ID,
		&flat.HouseID,
		&flat.Price,
		&flat.Rooms,
		&flat.Status,
		&flat.CreatedAt,
		&flat.UpdatedAt,
	)
	if err != nil {
		pg.log.Error("%s: %v", op, err)
		return flat, fmt.Errorf("failed to update flat status: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		pg.log.Error("%s: %v", op, err)
		return flat, fmt.Errorf("failed to commit tx: %w", err)
	}

	return flat, nil
}

func (pg *DbPostgres) GetFlatByID(ctx context.Context, flatID int) (models.Flat, error) {
	const op = "postgres.GetFlatByID"

	var flat models.Flat

	query := `
		SELECT id, house_id, price, rooms, status, created_at, updated_at
		FROM flats
		WHERE id = $1;
	`

	err := pg.db.QueryRow(ctx, query, flatID).Scan(
		&flat.ID,
		&flat.HouseID,
		&flat.Price,
		&flat.Rooms,
		&flat.Status,
		&flat.CreatedAt,
		&flat.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return flat, fmt.Errorf("%s: flat not found: %w", op, err)
		}
		return flat, fmt.Errorf("%s: failed to get flat by ID: %w", op, err)
	}

	return flat, nil
}

func (pg *DbPostgres) GetHouse(ctx context.Context, houseID int) (models.House, error) {
	const op = "postgres.GetHouse"

	query := `
		SELECT id, address, year, developer, created_at, updated_at
		FROM houses
		WHERE id = $1
		LIMIT 1;
	`

	var house models.House
	err := pg.db.QueryRow(ctx, query, houseID).Scan(
		&house.ID,
		&house.Address,
		&house.Year,
		&house.Developer,
		&house.CreatedAt,
		&house.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return house, dbErr.ErrHouseNotFound
	} else if err != nil {
		return house, fmt.Errorf("%s: %w", op, err)
	}

	return house, nil
}

func (pg *DbPostgres) GetFlatsByHouseID(ctx context.Context, houseID int, includeAll bool) ([]models.Flat, error) {
	const op = "postgres.GetFlatsByHouseID"

	query := `
		SELECT id, house_id, price, rooms, status, created_at, updated_at
		FROM flats
		WHERE house_id = $1
	`
	if !includeAll {
		query += " AND status = 'approved'"
	}

	rows, err := pg.db.Query(ctx, query, houseID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var flats []models.Flat
	for rows.Next() {
		var flat models.Flat
		err := rows.Scan(
			&flat.ID,
			&flat.HouseID,
			&flat.Price,
			&flat.Rooms,
			&flat.Status,
			&flat.CreatedAt,
			&flat.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		flats = append(flats, flat)
	}

	return flats, nil
}

func (pg *DbPostgres) SubscribeToHouse(ctx context.Context, houseID int, email string) (int, error) {
	const op = "postgres.SubscribeToHouse"

	query := `
		INSERT INTO subscriptions (house_id, email, created_at)
		VALUES ($1, $2, NOW())
		RETURNING id;
	`

	var subscriptionID int
	err := pg.db.QueryRow(ctx, query, houseID, email).Scan(&subscriptionID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return subscriptionID, nil
}

func (pg *DbPostgres) GetSubscribers(ctx context.Context, houseID int) ([]string, error) {
	const op = "postgres.GetSubscribers"

	query := `
		SELECT email
		FROM subscriptions
		WHERE house_id = $1;
	`

	rows, err := pg.db.Query(ctx, query, houseID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		emails = append(emails, email)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return emails, nil
}
