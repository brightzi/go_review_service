package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// OperationRepo is a Greater repo.
type OperationRepo interface {
	AuditReview(ctx context.Context, req *AuditReviewParam) error
	AuditAppeal(ctx context.Context, req *AuditAppealParam) error
}

// OperationUsecase is a Operation usecase.
type OperationUsecase struct {
	repo OperationRepo
	log  *log.Helper
}

// NewOperationUsecase new a Operation usecase.
func NewOperationUsecase(repo OperationRepo, logger log.Logger) *OperationUsecase {
	return &OperationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (o *OperationUsecase) AuditReview(ctx context.Context, req *AuditReviewParam) error {
	o.log.WithContext(ctx).Infof("[biz] AuditReview: %v", req)
	err := o.repo.AuditReview(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *OperationUsecase) AuditAppeal(ctx context.Context, req *AuditAppealParam) error {
	o.log.WithContext(ctx).Infof("[biz] AuditAppeal: %v", req)
	err := o.repo.AuditAppeal(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
