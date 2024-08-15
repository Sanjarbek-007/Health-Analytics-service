package mongodb

import (
	"context"
	"fmt"
	pb "health-service/genproto/health"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MedicalRecordRepository interface {
	AddMedicalReport(ctx context.Context, Req *pb.AddMedicalReportReq) (*pb.AddMedicalReportRes, error)
	GetMedicalReport(ctx context.Context, req *pb.GetMedicalReportReq) (*pb.GetMedicalReportRes, error)
	GetMedicalReportById(ctx context.Context, req *pb.GetMedicalReportByIdReq) (*pb.GetMedicalReportByIdRes, error)
	UpdateMedicalReport(ctx context.Context, req *pb.UpdateMedicalReportReq) (*pb.UpdateMedicalReportRes, error)
	DeleteMedicalReport(ctx context.Context, req *pb.DeleteMedicalReportReq) (*pb.DeleteMedicalReportRes, error)
}

type medicalRecordRepositoryImpl struct {
	coll *mongo.Database
}

func NewMedicalRecordRepository(db *mongo.Database) MedicalRecordRepository {
	return &medicalRecordRepositoryImpl{coll: db}
}

func (r *medicalRecordRepositoryImpl) AddMedicalReport(ctx context.Context, req *pb.AddMedicalReportReq) (*pb.AddMedicalReportRes, error) {
	coll := r.coll.Collection("medical_record")
	id := uuid.NewString()

	_, err := coll.InsertOne(ctx, bson.M{
		"_id":         id,
		"userId":      req.UserId,
		"recordType":  req.RecordType,
		"recordDate":  time.Now().Format("2006/01/02"),
		"description": req.Description,
		"doctorId":    req.DoctorId,
		"attachments": req.Attachments,
		"createdAt":   time.Now().Format("2006/01/02"),
		"updatedAt":   time.Now().Format("2006/01/02"),
		"deletedAt":   0,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddMedicalReportRes{
		Id: id,
	}, err
}

// 1152345678902345678909876543456789876543456789876543234567898765433456789987654323456789876543234567887654
func (r *medicalRecordRepositoryImpl) GetMedicalReport(ctx context.Context, req *pb.GetMedicalReportReq) (*pb.GetMedicalReportRes, error) {
    coll := r.coll.Collection("medical_record")
    filter := bson.M{
        "userId":    req.UserId,
        "deletedAt": 0,
    }

    options := options.Find()
    if req.Limit > 0 {
        options.SetLimit(int64(req.Limit))
    }
    if req.Offset > 0 {
        options.SetSkip(int64(req.Offset))
    }

    cursor, err := coll.Find(ctx, filter, options)
    if err != nil {
        return nil, fmt.Errorf("failed to execute query: %v", err)
    }
    defer cursor.Close(ctx)

    var records []*pb.MedicalReport
    for cursor.Next(ctx) {
        var record pb.MedicalReport
        if err := cursor.Decode(&record); err != nil {
            return nil, fmt.Errorf("failed to decode record: %v", err)
        }
        records = append(records, &pb.MedicalReport{
            Id:          record.Id,
            RecordType:  record.RecordType,
            RecordDate:  record.RecordDate,
            Description: record.Description,
            DoctorId:    record.DoctorId,
            Attachments: record.Attachments,
            FirstName:   record.FirstName,
            LastName:    record.LastName,
        })
    }

    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("cursor error: %v", err)
    }

    return &pb.GetMedicalReportRes{
        MedicalReport: records,
    }, nil
}

func (r *medicalRecordRepositoryImpl) GetMedicalReportById(ctx context.Context, req *pb.GetMedicalReportByIdReq) (*pb.GetMedicalReportByIdRes, error) {
	var record pb.GetMedicalReportByIdRes
	coll := r.coll.Collection("medical_record")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"_id": req.Id}, {"deletedAt": 0}}}).Decode(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *medicalRecordRepositoryImpl) UpdateMedicalReport(ctx context.Context, req *pb.UpdateMedicalReportReq) (*pb.UpdateMedicalReportRes, error) {
	coll := r.coll.Collection("medical_record")
	id := req.GetId()

	filter := bson.M{"_id": id}

	update := bson.M{}

	if req.RecordType != "" {
		update["recordType"] = req.RecordType
	}
	if req.Description != "" {
		update["description"] = req.Description
	}
	if req.DoctorId != "" {
		update["doctorId"] = req.DoctorId
	}
	if len(req.Attachments) > 0 {
		update["attachments"] = req.Attachments
	}

	if len(update) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	update["updatedAt"] = time.Now()

	update = bson.M{"$set": update}

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update medical record: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no medical record found with id: %s", req.GetId())
	}

	return nil, nil

}

func (r *medicalRecordRepositoryImpl) DeleteMedicalReport(ctx context.Context, req *pb.DeleteMedicalReportReq) (*pb.DeleteMedicalReportRes, error) {
	coll := r.coll.Collection("medical_record")

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deletedAt", Value: time.Now().Unix()}}}}

	filter := bson.D{{Key: "_id", Value: req.Id}, {Key: "deletedAt", Value: 0}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return &pb.DeleteMedicalReportRes{
			Message: false,
		}, err
	}

	return &pb.DeleteMedicalReportRes{
		Message: true,
	}, nil
}
