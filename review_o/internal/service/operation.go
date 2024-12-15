package service

import (
	"context"
	"fmt"

	pb "review_o/api/operation/v1"
	"review_o/internal/biz"
)

type OperationService struct {
	pb.UnimplementedOperationServer
	uc *biz.OperationUsecase
}

func NewOperationService(uc *biz.OperationUsecase) *OperationService {
	return &OperationService{uc: uc}
}

func (s *OperationService) AuditReview(ctx context.Context, req *pb.AuditReviewRequest) (*pb.AuditReviewReply, error) {
	fmt.Printf("[service] AuditReview, req: %v\n", req)
	err := s.uc.AuditReview(ctx, &biz.AuditReviewParam{
		ReviewID: req.ReviewID,
		Status:   req.Status,
		OpUser:   req.OpUser,
		OpReason: req.OpReason,
	})

	if err != nil {
		return nil, err
	}

	return &pb.AuditReviewReply{}, nil
}
func (s *OperationService) AuditAppeal(ctx context.Context, req *pb.AuditAppealRequest) (*pb.AuditAppealReply, error) {
	fmt.Printf("[service] AuditAppeal, req: %v\n", req)

	err := s.uc.AuditAppeal(ctx, &biz.AuditAppealParam{
		AppealID: req.AppealID,
		ReviewID: req.ReviewID,
		Status:   req.Status,
		OpUser:   req.OpUser,
		OpReason: req.OpReason,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AuditAppealReply{}, nil
}
