package biz

import (
	"context"
	"errors"
	"review_service/internal/data/model"
	"review_service/pkg/snowflake"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Review is a Review model.
type Review struct {
	Hello string
}

// ReviewRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
	GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error)
	SaveReply(ctx context.Context, reply *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
	GetReviewReply(ctx context.Context, replyID int64) (*model.ReviewReplyInfo, error)
}

// ReviewUsecase is a Review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewReviewUsecase new a Review usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateReview creates a Review, and returns the new Review.
func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Infof("CreateReview: %v", review)
	review.DeleteAt = time.Now()
	// review.ReviewID = snowflake.GenID()
	reviews, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, errors.New("get review by order id failed")
	}
	if len(reviews) > 0 {
		return nil, errors.New("order id already exists")
	}

	review.ReviewID = snowflake.GenID()
	return uc.repo.SaveReview(ctx, review)
}

func (uc *ReviewUsecase) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	return uc.repo.GetReview(ctx, reviewID)
}

func (uc *ReviewUsecase) CreateReply(ctx context.Context, param *ReplyParam) (*model.ReviewReplyInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz ]CreateReply param: %v", param)
	reply := &model.ReviewReplyInfo{
		ReplyID:   snowflake.GenID(),
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	}
	reply.DeleteAt = time.Now()
	return uc.repo.SaveReply(ctx, reply)
}

func (uc *ReviewUsecase) GetReply(ctx context.Context, reviewID int64) (*model.ReviewReplyInfo, error) {
	return uc.repo.GetReviewReply(ctx, reviewID)
}
