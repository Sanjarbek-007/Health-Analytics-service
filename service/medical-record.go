package service

import (
	"context"
	pb "health-service/genproto/health"
	"health-service/storage"
	"log/slog"
)

type MedicalService struct {
	pb.UnimplementedMedicalRecordServer
	Logger  *slog.Logger
	Storage storage.IStorage
}

func NewMedicalService(logger *slog.Logger, storage storage.IStorage) *MedicalService {
	return &MedicalService{Logger: logger, Storage: storage}
}

func (s *MedicalService) AddMedicalReport(ctx context.Context, req *pb.AddMedicalReportReq) (*pb.AddMedicalReportRes, error) {
	s.Logger.Info("Received AddMedicalReport request", "user_id", req.UserId, "record_type", req.RecordType, "description", req.Description, "doctor_id", req.DoctorId, "attachments", req.Attachments)
	res, err := s.Storage.MedicalRecordRepository().AddMedicalReport(ctx, req)
	if err!= nil {
        return nil, err
    }

	return res, nil
}

func (s *MedicalService) GetMedicalReport(ctx context.Context, req *pb.GetMedicalReportReq) (*pb.GetMedicalReportRes, error) {
	s.Logger.Info("Received GetMedicalReport request", "user_id", req.UserId)
    res, err := s.Storage.MedicalRecordRepository().GetMedicalReport(ctx, req)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (s *MedicalService) GetMedicalReportById(ctx context.Context, req *pb.GetMedicalReportByIdReq) (*pb.GetMedicalReportByIdRes, error) {
	s.Logger.Info("Received GetMedicalReportById request", "id", req.Id)
    res, err := s.Storage.MedicalRecordRepository().GetMedicalReportById(ctx, req)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (s *MedicalService) UpdateMedicalReport(ctx context.Context, req *pb.UpdateMedicalReportReq) (*pb.UpdateMedicalReportRes, error) {
	s.Logger.Info("Received UpdateMedicalReport request", "user_id", req.Id, "record_type", req.RecordType, "description", req.Description, "doctor_id", req.DoctorId, "attachments", req.Attachments)
    res, err := s.Storage.MedicalRecordRepository().UpdateMedicalReport(ctx, req)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (s *MedicalService) DeleteMedicalReport(ctx context.Context, req *pb.DeleteMedicalReportReq) (*pb.DeleteMedicalReportRes, error) {
	s.Logger.Info("Received DeleteMedicalReport request", "user_id", req.Id, "id", req.Id)
    res, err := s.Storage.MedicalRecordRepository().DeleteMedicalReport(ctx, req)
    if err!= nil {
        return nil, err
    }

    return res, nil
}