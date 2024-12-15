package service

import (
	"context"

	pb "review_b/api/business/v1"
	"review_b/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	replyID, err := s.uc.ReplyReview(ctx, &biz.ReplyReviewParam{
		ReviewID:  req.ReviewID,
		StoreID:   req.StoreID,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyID: replyID}, nil
}
func (s *BusinessService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	appealId, err := s.uc.AppealReview(ctx, &biz.AppealReviewParam{
		ReviewID:  req.ReviewID,
		StoreID:   req.StoreID,
		Reason:    req.Reason,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AppealReviewReply{AppealID: appealId}, nil
}
