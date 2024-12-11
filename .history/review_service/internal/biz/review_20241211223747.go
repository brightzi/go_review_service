package biz

import (
	"context"
	"review_service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

// Review is a Review model.
type Review struct {
	Hello string
}

// ReviewRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*Review, error)
	Update(context.Context, *Review) (*Review, error)
	FindByID(context.Context, int64) (*Review, error)
	ListByHello(context.Context, string) ([]*Review, error)
	ListAll(context.Context) ([]*Review, error)
}

// ReviewUsecase is a Review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewReviewUsecase new a Review usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateReview creates a Review, and returns the new Review.
func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Infof("CreateReview: %v", review)
	return uc.repo.Save(ctx, g)
}
