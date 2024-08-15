package service

import (
	"context"
	pb "health-service/genproto/health"
	"health-service/storage"
	"log/slog"
)

type HealthMonitoringService struct {
	pb.UnimplementedHealthMonitoringServer
	logger *slog.Logger
	repo   storage.IStorage
}

func NewHealthMonitoringService(logger *slog.Logger, repo storage.IStorage) *HealthMonitoringService {
	return &HealthMonitoringService{
		logger: logger,
		repo:   repo,
	}
}

func (s *HealthMonitoringService) GetHealthMonitor(ctx context.Context, req *pb.UserId) (*pb.GetHealthMonitorsRes, error) {
	s.logger.Info("GetHealthMonitor rpc method is working")
	res, err := s.repo.MonitoringRepository().GetHealthMonitor(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetHealthMonitor rpc method finished")
	return res, nil
}
