package home

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"home-service/pkg/pb"
)

func (s *serverAPI) CreateFlat(ctx context.Context, req *pb.CreateFlatRequest) (*pb.CreateFlatResponse, error) {
	_, _, err := s.AuthCheck(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token not valid")
	}

	if err := s.validateCreateFlatRequest(req); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid request")
	}

	flat, err := s.home.CreateFlat(ctx, int(req.HouseId), int(req.Price), int(req.Rooms))
	if err != nil {
		return nil, status.Error(codes.Internal, "fail createFlat")
	}

	return &pb.CreateFlatResponse{
		Id:      int32(flat.ID),
		Price:   int32(flat.Price),
		HouseId: int32(flat.HouseID),
		Rooms:   int32(flat.Rooms),
		Status:  flat.Status,
	}, nil
}

func (s *serverAPI) validateCreateFlatRequest(req *pb.CreateFlatRequest) error {
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
