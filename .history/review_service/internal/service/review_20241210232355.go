package service

import (
	"context"

	pb "review_service/api/review/v1"
)

type ReviewService struct {
	pb.UnimplementedReviewServer
}

func NewReviewService() *ReviewService {
	return &ReviewService{}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	return &pb.CreateReviewResponse{}, nil
}
