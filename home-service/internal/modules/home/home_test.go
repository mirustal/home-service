package home

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"home-service/internal/models"
	"home-service/mocks"
)

func TestHome_CreateHouse(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	t.Run("Success", func(t *testing.T) {
		mockHouseCreater := new(mocks.HouseCreater)
		mockHouseGetter := new(mocks.HouseGetter)

		expectedHouse := models.House{ID: 1, Address: "Test Address", Year: 2020, Developer: "Test Developer"}
		mockHouseCreater.On("CreateHouse", ctx, "Test Address", 2020, "Test Developer").Return(1, nil)
		mockHouseGetter.On("GetHouse", ctx, 1).Return(expectedHouse, nil)

		home := Home{
			log:          logger,
			houseCreater: mockHouseCreater,
			houseGetter:  mockHouseGetter,
		}

		house, err := home.CreateHouse(ctx, "Test Address", 2020, "Test Developer")

		assert.NoError(t, err)
		assert.Equal(t, expectedHouse, house)
		mockHouseCreater.AssertExpectations(t)
		mockHouseGetter.AssertExpectations(t)
	})

	t.Run("CreateHouse Error", func(t *testing.T) {
		mockHouseCreater := new(mocks.HouseCreater)
		mockHouseGetter := new(mocks.HouseGetter)

		mockHouseCreater.On("CreateHouse", ctx, "Test Address", 2020, "Test Developer").Return(0, errors.New("creation failed"))

		home := Home{
			log:          logger,
			houseCreater: mockHouseCreater,
			houseGetter:  mockHouseGetter,
		}

		house, err := home.CreateHouse(ctx, "Test Address", 2020, "Test Developer")

		assert.Error(t, err)
		assert.Empty(t, house)
		mockHouseCreater.AssertExpectations(t)
	})

	t.Run("GetHouse Error", func(t *testing.T) {
		mockHouseCreater := new(mocks.HouseCreater)
		mockHouseGetter := new(mocks.HouseGetter)

		mockHouseCreater.On("CreateHouse", ctx, "Test Address", 2020, "Test Developer").Return(1, nil)
		mockHouseGetter.On("GetHouse", ctx, 1).Return(models.House{}, errors.New("house not found"))

		home := Home{
			log:          logger,
			houseCreater: mockHouseCreater,
			houseGetter:  mockHouseGetter,
		}

		house, err := home.CreateHouse(ctx, "Test Address", 2020, "Test Developer")

		assert.Error(t, err)
		assert.Empty(t, house)
		mockHouseCreater.AssertExpectations(t)
		mockHouseGetter.AssertExpectations(t)
	})
}

func TestHome_GetFlatsInHouse(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	t.Run("Success", func(t *testing.T) {
		mockFlatsByHouseGetter := new(mocks.FlatsByHouseGetter)

		expectedFlats := []models.Flat{
			{ID: 1, HouseID: 1, Price: 100000, Rooms: 2},
			{ID: 2, HouseID: 1, Price: 150000, Rooms: 3},
		}

		mockFlatsByHouseGetter.On("GetFlatsByHouseID", ctx, 1, false).Return(expectedFlats, nil)

		home := Home{
			log:               logger,
			flatByHouseGetter: mockFlatsByHouseGetter,
		}

		flats, err := home.GetFlatsInHouse(ctx, 1, "user")

		assert.NoError(t, err)
		assert.Equal(t, expectedFlats, flats)
		mockFlatsByHouseGetter.AssertExpectations(t)
	})

	t.Run("GetFlatsByHouseID Error", func(t *testing.T) {
		mockFlatsByHouseGetter := new(mocks.FlatsByHouseGetter)

		mockFlatsByHouseGetter.On("GetFlatsByHouseID", ctx, 1, false).Return(nil, errors.New("failed to get flats"))

		home := Home{
			log:               logger,
			flatByHouseGetter: mockFlatsByHouseGetter,
		}

		flats, err := home.GetFlatsInHouse(ctx, 1, "user")

		assert.Error(t, err)
		assert.Nil(t, flats)
		mockFlatsByHouseGetter.AssertExpectations(t)
	})
}


func TestHome_SubscribeToHouse(t *testing.T) {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	t.Run("Success", func(t *testing.T) {
		mockHouseSubscriber := new(mocks.HouseSubscriber)

		mockHouseSubscriber.On("SubscribeToHouse", ctx, 1, "test@example.com").Return(1, nil)

		home := Home{
			log:             logger,
			houseSubscriber: mockHouseSubscriber,
		}

		err := home.SubscribeToHouse(ctx, 1, "test@example.com")

		assert.NoError(t, err)
		mockHouseSubscriber.AssertExpectations(t)
	})

	t.Run("SubscribeToHouse Error", func(t *testing.T) {
		mockHouseSubscriber := new(mocks.HouseSubscriber)

		mockHouseSubscriber.On("SubscribeToHouse", ctx, 1, "test@example.com").Return(0, errors.New("subscription failed"))

		home := Home{
			log:             logger,
			houseSubscriber: mockHouseSubscriber,
		}

		err := home.SubscribeToHouse(ctx, 1, "test@example.com")

		assert.Error(t, err)
		mockHouseSubscriber.AssertExpectations(t)
	})
}
