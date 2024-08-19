package home

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"home-service/pkg/pb"
)

func (s *serverAPI) SubscribeToHouse(ctx context.Context, req *pb.SubscribeToHouseRequest) (*pb.SubscribeToHouseResponse, error) {
	_, _, err := s.AuthCheck(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token not valid")
	}


	err = s.home.SubscribeToHouse(ctx, int(req.GetId()), req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to subscribe to house: %w", err))
	}

	return &pb.SubscribeToHouseResponse{}, nil
}
