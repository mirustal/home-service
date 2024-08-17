package home

import (
	"context"
	"fmt"
	"log/slog"

	"home-service/internal/models"
)

type Home struct {
	log *slog.Logger
	houseCreater HouseCreater
	houseGetter HouseGetter
	houseSubscriber HouseSubscriber
	flatCreater FlatCreater
	flatGetter FlatGetter
	flatByHouseGetter FlatsByHouseGetter
	flatStatusUpdater FlatsStatusUpdater
}

func New(
	log *slog.Logger,
	houseCreater HouseCreater,
	houseGetter HouseGetter,
	houseSubscriber HouseSubscriber,
	flatCreater FlatCreater,
	flatGetter FlatGetter,
	flatByHouseGetter FlatsByHouseGetter,
	flatStatusUpdater FlatsStatusUpdater) *Home {
	return &Home{
		log: log,
		houseCreater: houseCreater,
	houseGetter: houseGetter,
	houseSubscriber: houseSubscriber,
	flatGetter: flatGetter,
	flatCreater: flatCreater,
	flatByHouseGetter: flatByHouseGetter,
	flatStatusUpdater: flatStatusUpdater,
	}
}

type HouseCreater interface {
	CreateHouse(ctx context.Context, address string, year int, developer string) (int, error)
}

type HouseGetter interface {
	GetHouse(ctx context.Context, houseID int) (models.House, error)
}

type HouseSubscriber interface {
	SubscribeToHouse(ctx context.Context, houseID int, email string) (int, error)
}

type FlatCreater interface {
	CreateFlat(ctx context.Context, houseID, price, rooms int) (int, error)
}

type FlatGetter interface {
	GetFlatByID(ctx context.Context, flatID int) (models.Flat, error)
}

type FlatsStatusUpdater interface {
	UpdateFlatStatus(ctx context.Context, flatID int, status string) (models.Flat, error)
}

type FlatsByHouseGetter interface {
	GetFlatsByHouseID(ctx context.Context, houseID int, includeAll bool) ([]models.Flat, error)
}



func (h *Home) CreateHouse(ctx context.Context, address string, year int, developer string) (models.House, error) {
	const op = "internal.services.home.CreateHouse"
	log := h.log.With(
		slog.String("op", op),
	)

	houseID, err := h.houseCreater.CreateHouse(ctx, address, year, developer)
	if err != nil {
		log.Error("failed to create house")
		return models.House{}, fmt.Errorf("%s: %v", op, err)
	}

	house, err := h.houseGetter.GetHouse(ctx, houseID)
	log.Info("house created successfully", slog.Int("houseID", houseID))

	return house, nil
}

func (h *Home) GetFlatsInHouse(ctx context.Context, houseID int) ([]models.Flat, error) {
	const op = "internal.services.home.GetFlatsInHouse"
	log := h.log.With(
		slog.String("op", op),
	)

	flats, err := h.flatByHouseGetter.GetFlatsByHouseID(ctx, houseID, false)
	if err != nil {
		log.Error("failed to get flats in house")
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	log.Info("flats retrieved successfully", slog.Int("houseID", houseID))

	return flats, nil
}

func (h *Home) SubscribeToHouse(ctx context.Context, houseID int, email string) error {
	const op = "internal.services.home.SubscribeToHouse"
	log := h.log.With(
		slog.String("op", op),
	)

	_, err := h.houseSubscriber.SubscribeToHouse(ctx, houseID, email)
	if err != nil {
		log.Error("failed to subscribe in house")
		return fmt.Errorf("%s: %v", op, err)
	}

	log.Info("user subscribed to house", slog.Int("houseID", houseID), slog.String("email", email))

	return nil
}


func (h *Home) CreateFlat(ctx context.Context, houseID int, price int, rooms int) (models.Flat, error) {
	const op = "internal.services.home.CreateFlat"
	log := h.log.With(
		slog.String("op", op),
	)

	flatID, err := h.flatCreater.CreateFlat(ctx, houseID, price, rooms)
	if err != nil {
		log.Error("failed to create flat")
		return models.Flat{}, fmt.Errorf("%s: %v", op, err)
	}

	flat, err := h.flatGetter.GetFlatByID(ctx, flatID)
	if err != nil {
		log.Error("failed to retrieve flat after creation")
		return models.Flat{}, fmt.Errorf("%s: %v", op, err)
	}

	log.Info("flat created successfully", slog.Int("flatID", flatID))

	return flat, nil
}


func (h *Home) UpdateFlat(ctx context.Context, flatID int, status string) (models.Flat, error) {
	const op = "internal.services.home.UpdateFlat"
	log := h.log.With(
		slog.String("op", op),
	)

	flat, err := h.flatStatusUpdater.UpdateFlatStatus(ctx, flatID, status)
	if err != nil {
		log.Error("failed to update flat status")
		return models.Flat{}, fmt.Errorf("%s: %v", op, err)
	}

	log.Info("flat status updated successfully", slog.Int("flatID", flatID), slog.String("status", status))

	return flat, nil
}

