package data

import (
	"context"
	v1 "review_o/api/review/v1"
	"review_o/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type operationRepo struct {
	data *Data
	log  *log.Helper
}

// NewOperationRepo .
func NewOperationRepo(data *Data, logger log.Logger) biz.OperationRepo {
	return &operationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o *operationRepo) AuditReview(ctx context.Context, req *biz.AuditReviewParam) error {
	o.log.WithContext(ctx).Infof("[data] AuditReview: %v", req)
	_, err := o.data.rc.AuditReview(ctx, &v1.AuditReviewRequest{
		ReviewID: req.ReviewID,
		Status:   req.Status,
		OpUser:   req.OpUser,
		OpReason: req.OpReason,
	})
	if err != nil {
		return err
	}

	return nil
}

func (o *operationRepo) AuditAppeal(ctx context.Context, req *biz.AuditAppealParam) error {
	o.log.WithContext(ctx).Infof("[data] AuditAppeal: %v", req)
	_, err := o.data.rc.AuditAppeal(ctx, &v1.AuditAppealRequest{
		ReviewID: req.ReviewID,
		AppealID: req.AppealID,
		Status:   req.Status,
		OpUser:   req.OpUser,
		OpReason: &req.OpReason,
	})
	if err != nil {
		return err
	}
	return nil
}
