package service

import (
	"context"

	pb "kratos-timer/api/timer/v1"
	"kratos-timer/internal/biz"
)

type TimerService struct {
	pb.UnimplementedTimerServer
	uc *biz.TimerUsecase
}

func NewTimerService(uc *biz.TimerUsecase) *TimerService {
	return &TimerService{uc: uc}
}

func (s *TimerService) CreateTimer(ctx context.Context, req *pb.TimerInfo) (*pb.TimerInfo, error) {
	return &pb.TimerInfo{}, nil
}

func (s *TimerService) EnableTimer(ctx context.Context, req *pb.EnableTimerRequest) (*pb.TimerInfo, error) {
	return &pb.TimerInfo{}, nil
}

func (s *TimerService) DisableTimer(ctx context.Context, req *pb.DisableTimerRequest) (*pb.TimerInfo, error) {
	return &pb.TimerInfo{}, nil
}

func (s *TimerService) GetTimer(ctx context.Context, req *pb.GetTimerRequest) (*pb.TimerInfo, error) {
	return &pb.TimerInfo{}, nil
}

func (s *TimerService) ListTimer(ctx context.Context, req *pb.ListTimerRequest) (*pb.ListTimerReply, error) {
	return &pb.ListTimerReply{}, nil
}
