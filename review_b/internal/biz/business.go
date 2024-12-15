package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// BusinessRepo is a Greater repo.
type BusinessRepo interface {
	Reply(ctx context.Context, reply *ReplyReviewParam) (int64, error)
	AppealReview(ctx context.Context, appeal *AppealReviewParam) (int64, error)
}

// BusinessUsecase is a Business usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewBusinessUsecase new a Business usecase.
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateBusiness creates a Business, and returns the new Business.
func (uc *BusinessUsecase) ReplyReview(ctx context.Context, reply *ReplyReviewParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] ReplyReview: %v", reply)
	return uc.repo.Reply(ctx, reply)
}

func (uc *BusinessUsecase) AppealReview(ctx context.Context, appeal *AppealReviewParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] AppealReview: %v", appeal)
	return uc.repo.AppealReview(ctx, appeal)
}
