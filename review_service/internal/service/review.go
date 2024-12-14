package service

import (
	"context"
	"fmt"

	pb "review_service/api/review/v1"
	"review_service/internal/biz"
	"review_service/internal/data/model"
)

type ReviewService struct {
	pb.UnimplementedReviewServer
	uc *biz.ReviewUsecase
}

func NewReviewService(uc *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{uc: uc}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewReply, error) {
	fmt.Printf("[service] CreateReview, req:%v\n", req)
	var anonymous int32
	if req.Anonymous {
		anonymous = 1
	}
	review, err := s.uc.CreateReview(ctx, &model.ReviewInfo{
		UserID:       req.UserID,
		OrderID:      req.OrderID,
		StoreID:      req.StoreID,
		ServiceScore: req.ServiceScore,
		ExpressScore: req.ExpressScore,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Anonymous:    anonymous,
		Status:       0,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewReply{ReviewID: review.ReviewID}, nil
}

func (c *ReviewService) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.GetReviewReply, error) {
	fmt.Printf("[service] GetReview, req:%v\n", req.ReviewID)
	review, err := c.uc.GetReview(ctx, req.ReviewID)
	if err != nil {
		return nil, err
	}
	return &pb.GetReviewReply{
		Data: &pb.ReviewInfo{
			ReviewID:     review.ReviewID,
			UserID:       review.UserID,
			OrderID:      review.OrderID,
			Score:        review.Score,
			ServiceScore: review.ServiceScore,
			ExpressScore: review.ExpressScore,
			Content:      review.Content,
			PicInfo:      review.PicInfo,
			VideoInfo:    review.VideoInfo,
			Status:       review.Status,
		},
	}, err
}

func (c *ReviewService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	fmt.Printf("[service] ReplyReview, req:%v\n", req)
	reply, err := c.uc.CreateReply(ctx, &biz.ReplyParam{
		ReviewID:  req.GetReviewID(),
		StoreID:   req.GetStoreID(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.ReplyReviewReply{
		ReplyID: reply.ReplyID,
	}, nil
}

func (c *ReviewService) GetReplyReview(ctx context.Context, req *pb.GetReviewReplyRequest) (*pb.ReviewReplyInfo, error) {
	fmt.Printf("[service] GetReviewReply, req:%v\n", req)
	review_reply, err := c.uc.GetReply(ctx, req.ReplyID)
	if err != nil {
		return nil, err
	}

	return &pb.ReviewReplyInfo{
		ReplyID: review_reply.ReplyID,
		StoreID: review_reply.StoreID,
		Content: review_reply.Content,
	}, err

}

func (c *ReviewService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	fmt.Printf("[service] AppealReview, req:%v\n", req)

	appeal_review, err := c.uc.AppealReview(ctx, &biz.ReviewAppealParam{
		ReviewID:  req.ReviewID,
		Content:   req.Content,
		StoreID:   req.StoreID,
		Resaon:    req.Reason,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AppealReviewReply{
		AppealID: appeal_review.AppealID,
	}, nil
}

// ReviewID  int64
// Status    int32
// OpUser    string
// OpReason  string
// OpRemarks string
func (c *ReviewService) AuditReview(ctx context.Context, req *pb.AuditReviewRequest) (*pb.AuditReviewReply, error) {
	fmt.Printf("[service] AuditReview, req:%v\n", req)

	err := c.uc.AuditReview(ctx, &biz.AuditParam{
		ReviewID: req.ReviewID,
		Status:   req.Status,
		OpUser:   req.OpUser,
		OpReason: req.OpReason,
	})
	return &pb.AuditReviewReply{ReviewID: req.ReviewID, Status: req.Status}, err
}

func (c *ReviewService) AuditAppeal(ctx context.Context, req *pb.AuditAppealRequest) (*pb.AuditAppealReply, error) {
	fmt.Printf("[service] AuditAppeal, req:%v\n", req)
	err := c.uc.AuditAppeal(ctx, &biz.AuditAppealParam{
		ReviewID: req.ReviewID,
		AppealID: req.AppealID,
		Status:   req.Status,
		OpUser:   req.OpUser,
	})
	return &pb.AuditAppealReply{}, err
}

func (c *ReviewService) ListReviewByUserID(ctx context.Context, req *pb.ListReviewByUserIDRequest) (*pb.ListReviewByUserIDReply, error) {
	fmt.Printf("[service] ListReviewByUserID, req:%v\n", req)
	reviews, err := c.uc.ListReviewByUserID(ctx, req.UserID, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.ReviewInfo, 0, len(reviews))
	for _, review := range reviews {
		list = append(list, &pb.ReviewInfo{
			ReviewID:     review.ReviewID,
			UserID:       review.UserID,
			OrderID:      review.OrderID,
			Score:        review.Score,
			ServiceScore: review.ServiceScore,
			ExpressScore: review.ExpressScore,
		})
	}
	return &pb.ListReviewByUserIDReply{Data: list}, err
}
