package service

import (
	"context"

	pb "review_service/api/review/v1"
	"review_service/internal/biz"
)

type ReviewService struct {
	pb.UnimplementedReviewServer
	uc *biz.GreeterUsecase
}

func NewReviewService(uc *biz.GreeterUsecase) *ReviewService {
	return &ReviewService{uc: uc}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	return &pb.CreateReviewResponse{}, nil
}