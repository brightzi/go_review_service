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

// func (r *reviewRepo) Save(ctx context.Context, g *biz.Review) (*biz.Review, error) {
// 	return g, nil
// }

func (r *reviewRepo) SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error) {
	return nil, nil
}
