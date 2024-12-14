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
	AppealReview(ctx context.Context, appealParam *model.ReviewAppealInfo) (*model.ReviewAppealInfo, error)
	AuditReview(ctx context.Context, param *AuditParam) error
	AuditAppeal(ctx context.Context, param *AuditAppealParam) error
	ListReviewByUserID(ctx context.Context, userID int64, offset int32, limit int32) ([]*model.ReviewInfo, error)
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

// 获取商家回复
func (uc *ReviewUsecase) GetReply(ctx context.Context, replyID int64) (*model.ReviewReplyInfo, error) {
	return uc.repo.GetReviewReply(ctx, replyID)
}

// 商家对用户评价进行申诉
func (uc *ReviewUsecase) AppealReview(ctx context.Context, param *ReviewAppealParam) (*model.ReviewAppealInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz ]AppealReview param: %v", param)
	return uc.repo.AppealReview(ctx, &model.ReviewAppealInfo{
		AppealID:  snowflake.GenID(),
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		Reason:    param.Resaon,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
		DeleteAt:  time.Now(),
	})
}

func (uc *ReviewUsecase) AuditReview(ctx context.Context, param *AuditParam) error {
	uc.log.WithContext(ctx).Debugf("[biz ]AuditReview param: %v", param)
	return uc.repo.AuditReview(ctx, param)
}

func (uc *ReviewUsecase) AuditAppeal(ctx context.Context, param *AuditAppealParam) error {
	uc.log.WithContext(ctx).Debugf("[biz ]AuditAppealReview param: %v", param)
	return uc.repo.AuditAppeal(ctx, param)
}

func (uc *ReviewUsecase) ListReviewByUserID(ctx context.Context, userID int64, page int32, size int32) ([]*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz ]ListReviewByUserID userID: %v", userID)
	if page <= 0 {
		page = 1
	}

	if size <= 0 || size > 50 {
		size = 10
	}

	offset := (page - 1) * size
	limit := size
	return uc.repo.ListReviewByUserID(ctx, userID, offset, limit)
}
