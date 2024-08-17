package home

import (
	"context"
	"log/slog"

	"home-service/internal/models"
)

type Home struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Home {
	return &Home{
		log: log,
	}
}

type UserGetter interface {
	GetUser(ctx context.Context, userID string) (models.User, error)
}

func (h *Home) CreateHouse(ctx context.Context, address string, year int, developer string) (house models.House, err error) {
	return models.House{}, nil
}

func (h *Home) GetFlatsInHouse(ctx context.Context, id int) (flat models.Flat, err error) {
	return models.Flat{}, nil
}

func (h *Home) SubscribeToHouse(ctx context.Context, id int, email string) (err error) {
	return nil
}

func (h *Home) CreateFlat(ctx context.Context, house_id int, price int, rooms int) (flat models.Flat, err error) {
	return models.Flat{}, nil
}

func (h *Home) UpdateFlat(ctx context.Context, id int, status string) (models.House, error) {
	return models.House{}, nil
}
