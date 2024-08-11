package service

import (
	"context"
	pb "health-service/genproto/health"
	"health-service/storage"
	"log/slog"
)

type WearableService struct {
	pb.UnimplementedWearableServer
	Logger  *slog.Logger
	Storage storage.IStorage
}

func NewWearableService(logger *slog.Logger, storage storage.IStorage) *WearableService {
	return &WearableService{Logger: logger, Storage: storage}
}

func (s *WearableService) AddWearableData(ctx context.Context, req *pb.AddWearableDataReq) (*pb.AddWearableDataRes, error) {
	s.Logger.Info("Received AddWearableData request", "user_id", req.UserId, "device_type", req.DeviceType, "data_type", req.DataType, "data_value", req.DataValue)
    res, err := s.Storage.WearableRepository().AddWearableData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *WearableService) GetWearableData(ctx context.Context, req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error) {
	s.Logger.Info("Received GetWearableData request", "user_id", req.UserId)
    res, err := s.Storage.WearableRepository().GetWearableData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *WearableService) GetWearableDataById(ctx context.Context, req *pb.GetWearableDataByIdReq) (*pb.GetWearableDataByIdRes, error) {
	s.Logger.Info("Received GetWearableDataById request", "id", req.Id)
    res, err := s.Storage.WearableRepository().GetWearableDataById(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *WearableService) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataReq) (*pb.UpdateWearableDataRes, error) {
	s.Logger.Info("Received UpdateWearableData request", "user_id", req.Id, "device_type", req.DeviceType, "data_type", req.DataType, "data_value", req.DataValue)
    res, err := s.Storage.WearableRepository().UpdateWearableData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *WearableService) DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataReq) (*pb.DeleteWearableDataRes, error) {
	s.Logger.Info("Received DeleteWearableData request", "user_id", req.Id, "id", req.Id)
    res, err := s.Storage.WearableRepository().DeleteWearableData(ctx, req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}