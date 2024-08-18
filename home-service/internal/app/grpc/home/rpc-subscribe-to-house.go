package home

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"home-service/pkg/pb"
)

func (s *serverAPI) SubscribeToHouse(ctx context.Context, req *pb.SubscribeToHouseRequest) (*pb.SubscribeToHouseResponse, error) {

	err := s.home.SubscribeToHouse(ctx, int(req.GetId()), req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to subscribe to house: %v", err))
	}

	return &pb.SubscribeToHouseResponse{}, nil
}
