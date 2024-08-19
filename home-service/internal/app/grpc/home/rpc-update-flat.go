package home

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"home-service/pkg/pb"
)

func (s *serverAPI) UpdateFlat(ctx context.Context, req *pb.UpdateFlatRequest) (*pb.UpdateFlatResponse, error) {
	_, userType, err := s.AuthCheck(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token not valid")
	}

	if !IsModerator(userType) {
		return nil, status.Error(codes.Unauthenticated, "not permission")
	}

	statusStr := req.GetStatus()
	flat, err := s.home.UpdateFlat(ctx, int(req.GetId()), statusStr)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update flat: %w", err))
	}

	return &pb.UpdateFlatResponse{
		Id:      int32(flat.ID),
		HouseId: int32(flat.HouseID),
		Price:   int32(flat.Price),
		Rooms:   int32(flat.Rooms),
		Status:  req.Status,
	}, nil
}

func statusEnumToString(status pb.Status) string {
	switch status {
	case pb.Status_created:
		return "created"
	case pb.Status_approved:
		return "approved"
	case pb.Status_declined:
		return "declined"
	case pb.Status_onModeration:
		return "on_moderation"
	default:
		return "UNKNOWN"
	}
}

func stringToStatusEnum(status string) pb.Status {
	switch status {
	case "created":
		return pb.Status_created
	case "approved":
		return pb.Status_approved
	case "declined":
		return pb.Status_declined
	case "on_moderation":
		return pb.Status_onModeration
	default:
		return pb.Status(0) // UNKNOWN или другой подходящий статус по умолчанию
	}
}
