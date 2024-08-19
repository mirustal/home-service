package home

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"home-service/pkg/pb"
)

func (s *serverAPI) CreateHouse(ctx context.Context, req *pb.CreateHouseRequest) (*pb.CreateHouseResponse, error) {
	_, userType, err := s.AuthCheck(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token not valid")
	}

	if !IsModerator(userType) {
		return nil, status.Error(codes.Unauthenticated, "not permission")
	}

	if err := s.validateCreateHouseRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid request: %w", err))
	}

	house, err := s.home.CreateHouse(ctx, req.Address, int(req.Year), req.Developer)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create house: %w", err))
	}

	return &pb.CreateHouseResponse{
		Id:        int32(house.ID),
		Address:   house.Address,
		Year:      int32(house.Year),
		Developer: house.Developer,
		UpdatedAt: house.UpdatedAt.Format(time.RFC3339),
		CreatedAt: house.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *serverAPI) validateCreateHouseRequest(req *pb.CreateHouseRequest) error {
	if req.Address == "" {
		return errors.New("address is required")
	}
	if req.Year <= 0 {
		return errors.New("year must be a positive integer")
	}
	return nil
}
