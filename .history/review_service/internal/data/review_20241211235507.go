package data

import (
	"context"
	"time"

	"review_service/internal/biz"
	"review_service/pkg/snowflake"

	"review_service/internal/data/model"

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
	// var deleteAt *time.Time = nil
	review.DeleteAt = time.Now()
	review.ReviewID = snowflake.GenID()
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	return review, err
}
