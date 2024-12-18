package data

import (
	"context"
	"errors"
	"time"

	"review_service/internal/biz"

	"review_service/internal/data/model"
	"review_service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

type CustomReview struct {
	ReviewInfo *model.ReviewInfo
	DeleteAt   *time.Time `gorm:"column:delete_at" json:"delete_at"`
}

// NewReviewRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	return review, err
}

func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.
		WithContext(ctx).
		Where(r.data.query.ReviewInfo.OrderID.Eq(orderID)).
		Find()
}

func (r *reviewRepo) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.WithContext(ctx).Where(r.data.query.ReviewInfo.ReviewID.Eq(reviewID)).First()
}

func (r *reviewRepo) SaveReply(ctx context.Context, reply *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	// 查询该评论是否已有回复
	review, err := r.data.query.ReviewInfo.WithContext(ctx).
		Where(r.data.query.ReviewInfo.ReviewID.Eq(reply.ReviewID)).
		First()

	if err != nil {
		return nil, err
	}
	if review.HasReply == 1 {
		return nil, errors.New("该评论已有回复")
	}

	if review.StoreID != reply.StoreID {
		return nil, errors.New("该评论不属于该店铺")
	}

	err = r.data.query.Transaction(func(tx *query.Query) error {
		if err := tx.ReviewReplyInfo.WithContext(ctx).Save(reply); err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply failed: %v")
			return err
		}

		if _, err := tx.ReviewInfo.WithContext(ctx).
			Where(tx.ReviewInfo.ReviewID.Eq(reply.ReviewID)).
			Update(tx.ReviewInfo.HasReply, 1); err != nil {
			r.log.WithContext(ctx).Errorf("UpdateReview failed: %v", err)

			return err
		}
		return nil
	})

	return reply, err
}

func (r *reviewRepo) GetReviewReply(ctx context.Context, replyID int64) (*model.ReviewReplyInfo, error) {
	return r.data.query.ReviewReplyInfo.
		WithContext(ctx).
		Where(r.data.query.ReviewReplyInfo.ReplyID.Eq(replyID)).
		First()
}

func (r *reviewRepo) AppealReview(ctx context.Context, appealParam *model.ReviewAppealInfo) (*model.ReviewAppealInfo, error) {
	err := r.data.query.ReviewAppealInfo.WithContext(ctx).Save(appealParam)
	return appealParam, err
}

func (r *reviewRepo) AuditReview(ctx context.Context, param *biz.AuditParam) error {
	_, err := r.data.query.ReviewInfo.WithContext(ctx).
		Where(r.data.query.ReviewInfo.ReviewID.Eq(param.ReviewID)).
		Updates(map[string]interface{}{
			"status":     param.Status,
			"op_user":    param.OpUser,
			"op_reason":  param.OpReason,
			"op_remarks": param.OpRemarks,
		})
	return err
}

func (r *reviewRepo) AuditAppeal(ctx context.Context, param *biz.AuditAppealParam) error {
	_, err := r.data.query.ReviewAppealInfo.WithContext(ctx).
		Where(r.data.query.ReviewAppealInfo.AppealID.Eq(param.AppealID)).
		Updates(map[string]interface{}{
			"status":  param.Status,
			"op_user": param.OpUser,
			"reason":  param.OpReason,
		})
	return err
}

func (r *reviewRepo) ListReviewByUserID(ctx context.Context, userID int64, offset int32, limit int32) ([]*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.WithContext(ctx).
		Where(r.data.query.ReviewInfo.UserID.Eq(userID)).
		Order(r.data.query.ReviewInfo.ID.Desc()).
		Limit(int(limit)).
		Offset(int(offset)).
		Find()
}
