package data

import (
	"context"
	"time"

	"review_service/internal/biz"

	"review_service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

type CustomReview struct {
	review   *model.ReviewInfo
	DeleteAt *time.Time `gorm:"column:delete_at" json:"delete_at"`
}

// NewReviewRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	// review.DeleteAt = time.Now()
	// customReview := &CustomReview{
	// 	review:   review,
	// 	DeleteAt: nil,
	// }
	review.DeleteAt = time.Now()
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	return review, err
}
