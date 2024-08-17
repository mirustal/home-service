package homegrpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	GetFlatsInHouse(ctx context.Context, id int) (flat []models.Flat, err error)
	SubscribeToHouse(ctx context.Context, id int, email string) (err error)
	CreateFlat(ctx context.Context, house_id int, price int, rooms int) (flat models.Flat, err error)
	UpdateFlat(ctx context.Context, id int, status string) (models.Flat, error)
}

func (s *serverAPI) CreateHouse(ctx context.Context, req *homegrpc.CreateHouseRequest) (*homegrpc.CreateHouseResponse, error) {
	
	if err := s.validateCreateHouseRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid request: %v", err))
	}

	house, err := s.home.CreateHouse(ctx, req.Address, int(req.Year), req.Developer)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create house: %v", err))
	}


	return &homegrpc.CreateHouseResponse{
		Id:        int32(house.ID),
		Address:   house.Address,
		Year:      int32(house.Year),
		Developer: house.Developer,
	}, nil
}

func (s *serverAPI) GetFlatsInHouse(ctx context.Context, req *homegrpc.GetFlatsInHouseRequest) (*homegrpc.GetFlatsInHouseResponse, error) {
	flats, err := s.home.GetFlatsInHouse(ctx, int(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get flats in house: %v", err))
	}

	var flatResponses []*homegrpc.Flat
	for _, flat := range flats {
		flatResponses = append(flatResponses, &homegrpc.Flat{
			Id:    int32(flat.ID),
			Price: int32(flat.Price),
			Rooms: int32(flat.Rooms),
			Status: stringToStatusEnum(flat.Status),
		})
	}

	return &homegrpc.GetFlatsInHouseResponse{
		Flats: flatResponses,
	}, nil
}
func (s *serverAPI) SubscribeToHouse(ctx context.Context, req *homegrpc.SubscribeToHouseRequest) (*homegrpc.SubscribeToHouseResponse, error) {

	err := s.home.SubscribeToHouse(ctx, int(req.GetId()), req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to subscribe to house: %v", err))
	}

	return &homegrpc.SubscribeToHouseResponse{}, nil
}

func (s *serverAPI) CreateFlat(ctx context.Context, req *homegrpc.CreateFlatRequest) (*homegrpc.CreateFlatResponse, error) {
	if err := s.validateCreateFlatRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid request: %v", err))
	}

	flat, err := s.home.CreateFlat(ctx, int(req.HouseId), int(req.Price), int(req.Rooms))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create flat: %v", err))
	}

	return &homegrpc.CreateFlatResponse{
			Id:    int32(flat.ID),
			Price: int32(flat.Price),
			Rooms: int32(flat.Rooms),
			Status: stringToStatusEnum(flat.Status),
	}, nil
}

func (s *serverAPI) UpdateFlat(ctx context.Context, req *homegrpc.UpdateFlatRequest) (*homegrpc.UpdateFlatResponse, error) {
	statusStr := statusEnumToString(req.Status)

	flat, err := s.home.UpdateFlat(ctx, int(req.GetId()), statusStr)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update flat: %v", err))
	}

	return &homegrpc.UpdateFlatResponse{
			Id:     int32(flat.ID),
			Price:  int32(flat.Price),
			Rooms:  int32(flat.Rooms),
			Status: req.Status,  
	}, nil
}



func (s *serverAPI) validateCreateHouseRequest(req *homegrpc.CreateHouseRequest) error {
	if req.Address == "" {
		return errors.New("address is required")
	}
	if req.Year <= 0 {
		return errors.New("year must be a positive integer")
	}
	if req.Developer == "" {
		return errors.New("developer is required")
	}
	return nil
}

func (s *serverAPI) validateCreateFlatRequest(req *homegrpc.CreateFlatRequest) error {
	if req.HouseId <= 0 {
		return errors.New("house_id must be a positive integer")
	}
	if req.Price <= 0 {
		return errors.New("price must be a positive integer")
	}
	if req.Rooms <= 0 {
		return errors.New("rooms must be a positive integer")
	}
	return nil
}

func statusEnumToString(status homegrpc.Status) string {
	switch status {
	case homegrpc.Status_CREATED:
		return "CREATED"
	case homegrpc.Status_APPROVED:
		return "APPROVED"
	case homegrpc.Status_DECLINED:
		return "DECLINED"
	case homegrpc.Status_ON_MODERATION:
		return "ON_MODERATION"
	default:
		return "UNKNOWN"
	}
}

func stringToStatusEnum(status string) homegrpc.Status {
	switch status {
	case "CREATED":
		return homegrpc.Status_CREATED
	case "APPROVED":
		return homegrpc.Status_APPROVED
	case "DECLINED":
		return homegrpc.Status_DECLINED
	case "ON_MODERATION":
		return homegrpc.Status_ON_MODERATION
	default:
		return homegrpc.Status(0)  // UNKNOWN или другой подходящий статус по умолчанию
	}
}