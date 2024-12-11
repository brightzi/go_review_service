package data

import (
	"context"

	"review_service/internal/biz"

	"review_service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewReviewRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	review.DeleteAt = nil
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	return review, err
}
