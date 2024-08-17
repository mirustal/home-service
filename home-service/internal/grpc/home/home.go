package homegrpc

import (
	"context"

	"google.golang.org/grpc"

	"home-service/internal/models"
	homegrpc "home-service/pkg/pb"
)

type serverAPI struct {
	homegrpc.UnimplementedHouseServiceServer
	home Home
}

func Register(gRPC *grpc.Server, home Home) {
	homegrpc.RegisterHouseServiceServer(gRPC, &serverAPI{
		home: home,
	})
}

type Home interface {
	CreateHouse(ctx context.Context, address string, year int, developer string) (house models.House, err error)
	GetFlatsInHouse(ctx context.Context, id int) (flat models.Flat, err error)
	SubscribeToHouse(ctx context.Context, id int, email string) (err error)
	CreateFlat(ctx context.Context, house_id int, price int, rooms int) (flat models.Flat, err error)
	UpdateFlat(ctx context.Context, id int, status string) (models.House, error)
}

func (s *serverAPI) CreateHouse(ctx context.Context, req *homegrpc.CreateHouseRequest) (*homegrpc.CreateHouseResponse, error) {

	return &homegrpc.CreateHouseResponse{}, nil
}

func (s *serverAPI) GetFlatsInHouse(ctx context.Context, req *homegrpc.GetFlatsInHouseRequest) (*homegrpc.GetFlatsInHouseResponse, error) {
	return &homegrpc.GetFlatsInHouseResponse{}, nil
}

func (s *serverAPI) SubscribeToHouse(ctx context.Context, req *homegrpc.SubscribeToHouseRequest) (*homegrpc.SubscribeToHouseResponse, error) {

	return &homegrpc.SubscribeToHouseResponse{}, nil
}

func (s *serverAPI) CreateFlat(ctx context.Context, req *homegrpc.CreateFlatRequest) (*homegrpc.CreateFlatResponse, error) {

	return &homegrpc.CreateFlatResponse{}, nil
}

func (s *serverAPI) UpdateFlat(ctx context.Context, req *homegrpc.UpdateFlatRequest) (*homegrpc.UpdateFlatResponse, error) {
	return &homegrpc.UpdateFlatResponse{}, nil
}
