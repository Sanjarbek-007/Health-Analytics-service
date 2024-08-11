package service

import (
	"context"
	pb "health-service/genproto/health"
	"health-service/storage"
	"log/slog"
)

type HealthService struct {
	pb.UnimplementedHealthCheckServer
	Logger  *slog.Logger
	Storage storage.IStorage
}

func NewHealthService(logger *slog.Logger, storage storage.IStorage) *HealthService {
	return &HealthService{Logger: logger, Storage: storage}
}

func (s *HealthService) GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error) {
	s.Logger.Info("Received GenerateHealthRecommendations request", "user_id", req.UserId, "recommendation_type", req.RecommendationType, "description", req.Description, "priority", req.Priority)
	res, err := s.Storage.HealthRepository().GenerateHealthRecommendations(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *HealthService) GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error) {
	s.Logger.Info("Received GetRealtimeHealthMonitoring request", "user_id", req.UserId)
	res, err := s.Storage.HealthRepository().GetRealtimeHealthMonitoring(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *HealthService) GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error) {
    s.Logger.Info("Received GetDailyHealthSummary request", "user_id", req.UserId, "start_date", req.Date)
    res, err := s.Storage.HealthRepository().GetDailyHealthSummary(ctx, req)
	if err != nil {
        return nil, err
    }

	return res, nil
}

func (s *HealthService) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error) {
    s.Logger.Info("Received GetWeeklyHealthSummary request", "user_id", req.UserId, "start_date", req.StartDate, "end_date", req.EndDate)
    res, err := s.Storage.HealthRepository().GetWeeklyHealthSummary(ctx, req)
	if err != nil {
        return nil, err
    }

	return res, nil
}