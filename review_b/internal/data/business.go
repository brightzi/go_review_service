package data

import (
	"context"

	v1 "review_b/api/review/v1"
	"review_b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *businessRepo) Reply(ctx context.Context, reply *biz.ReplyReviewParam) (int64, error) {
	c.log.WithContext(ctx).Infof("[data] Reply, param:%v", reply)

	ret, err := c.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  reply.ReviewID,
		StoreID:   reply.StoreID,
		Content:   reply.Content,
		PicInfo:   reply.PicInfo,
		VideoInfo: reply.VideoInfo,
	})
	c.log.WithContext(ctx).Infof("[data] Reply, ret:%v, err:%v", ret, err)

	if err != nil {
		return 0, err
	}

	return ret.ReplyID, err
}

func (c *businessRepo) AppealReview(ctx context.Context, param *biz.AppealReviewParam) (int64, error) {
	c.log.WithContext(ctx).Infof("[data] AppealReview, param:%v", param)
	appeal, err := c.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Reason:    param.Reason,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		return 0, err
	}
	return appeal.AppealID, err
}
