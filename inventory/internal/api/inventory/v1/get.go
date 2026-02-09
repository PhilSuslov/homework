package v1

import (
	"context"

	pb "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *InventoryService) GetPart(ctx context.Context, req *pb.GetPartRequest) (*pb.GetPartResponse, error) {
	part, ok := s.parts[req.Uuid]
	if !ok {
		return nil, status.Error(codes.NotFound, "part not found")
	}

	return &pb.GetPartResponse{
		Part: part,
	}, nil

}