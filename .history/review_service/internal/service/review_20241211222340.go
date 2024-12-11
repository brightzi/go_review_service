package service

import (
	"context"

	pb "review_service/api/review/v1"
	"review_service/internal/biz"
)

type ReviewService struct {
	pb.UnimplementedReviewServer
	uc *biz.ReviewUsecase
}

func NewReviewService(uca *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{uc: uca}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	//

}
