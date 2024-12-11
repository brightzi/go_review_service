package data

import (
	"context"

	"review_service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewReviewRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.reviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) Save(ctx context.Context, g *biz.Review) (*biz.Review, error) {
	return g, nil
}

func (r *reviewRepo) Update(ctx context.Context, g *biz.Review) (*biz.Review, error) {
	return g, nil
}

func (r *reviewRepo) FindByID(context.Context, int64) (*biz.Review, error) {
	return nil, nil
}

func (r *reviewRepo) ListByHello(context.Context, string) ([]*biz.Review, error) {
	return nil, nil
}

func (r *reviewRepo) ListAll(context.Context) ([]*biz.Review, error) {
	return nil, nil
}
