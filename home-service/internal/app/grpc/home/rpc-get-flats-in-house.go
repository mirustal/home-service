package home

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"home-service/pkg/pb"
)

func (s *serverAPI) GetFlatsInHouse(ctx context.Context, req *pb.GetFlatsInHouseRequest) (*pb.GetFlatsInHouseResponse, error) {
	_, userType, err := s.AuthCheck(ctx)
	if err != nil {
		return nil, err
	}

	flats, err := s.home.GetFlatsInHouse(ctx, int(req.GetId()), userType)
	fmt.Println(flats)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get flats in house: %v", err))
	}

	var flatResponses []*pb.Flat
	for _, flat := range flats {
		flatResponses = append(flatResponses, &pb.Flat{
			Id:      int32(flat.ID),
			Price:   int32(flat.Price),
			HouseId: int32(flat.HouseID),
			Rooms:   int32(flat.Rooms),
			Status:  flat.Status,
		})
	}

	return &pb.GetFlatsInHouseResponse{
		Flats: flatResponses,
	}, nil
}
