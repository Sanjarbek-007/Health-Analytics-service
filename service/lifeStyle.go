package service

import (
	"context"
	pb "health-service/genproto/health"
	"health-service/storage"
	"log/slog"
)

type LifeStyleService struct {
	pb.UnimplementedLifeStyleServer
	Logger  *slog.Logger
	Storage storage.IStorage
}

func NewLifeStyleService(logger *slog.Logger, storage storage.IStorage) *LifeStyleService {
	return &LifeStyleService{Logger: logger, Storage: storage}
}

func (s *LifeStyleService) AddLifeStyleData(ctx context.Context, req *pb.AddLifeStyleDataReq) (*pb.AddLifeStyleDataRes, error) {
	s.Logger.Info("Received AddLifeStyleData request", "user_id", req.UserId, "data_type", req.DataType, "data_value", req.DataValue)
    res, err := s.Storage.LifeStyleRepository().AddLifeStyleData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) GetLifeStyleData(ctx context.Context, req *pb.GetLifeStyleDataReq) (*pb.GetLifeStyleDataRes, error) {
	s.Logger.Info("Received GetLifeStyleData request", "user_id", req.UserId)
    res, err := s.Storage.LifeStyleRepository().GetLifeStyleData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) GetLifeStyleDataById(ctx context.Context, req *pb.GetLifeStyleDataByIdReq) (*pb.GetLifeStyleDataByIdRes, error) {
	s.Logger.Info("Received GetLifeStyleDataById request", "user_id", req.Id, "id", req.Id)
    res, err := s.Storage.LifeStyleRepository().GetLifeStyleDataById(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) UpdateLifeStyleData(ctx context.Context, req *pb.UpdateLifeStyleDataReq) (*pb.UpdateLifeStyleDataRes, error) {
	s.Logger.Info("Received UpdateLifeStyleData request", "user_id", req.Id, "data_type", req.DataType, "data_value", req.DataValue, "id", req.Id)
    res, err := s.Storage.LifeStyleRepository().UpdateLifeStyleData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *LifeStyleService) DeleteLifeStyleData(ctx context.Context, req *pb.DeleteLifeStyleDataReq) (*pb.DeleteLifeStyleDataRes, error) {
	s.Logger.Info("Received DeleteLifeStyleData request", "user_id", req.Id, "id", req.Id)
    res, err := s.Storage.LifeStyleRepository().DeleteLifeStyleData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}