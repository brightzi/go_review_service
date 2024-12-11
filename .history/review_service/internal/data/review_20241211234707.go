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
	// review.DeleteAt = time.Now()
	// err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	// return review, err
	// 使用 CustomReview 结构体重写 DeleteAt 字段
	customReview := &model.CustomReview{
		ReviewInfo: review,
		DeleteAt:   nil, // 设置 DeleteAt 为 nil
	}

	// 保存 CustomReview 数据
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(customReview)
	if err != nil {
		return nil, err
	}
	return review, nil
}
