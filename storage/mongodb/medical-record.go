package mongodb

import (
	"context"
	"fmt"
	pb "health-service/genproto/health"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	_, err := coll.InsertOne(ctx, bson.M{
		"_id":         uuid.NewString(),
		"user_id":     req.UserId,
		"record_type": req.RecordType,
		"record_date": time.Now().Format("2006/01/02"),
		"description": req.Description,
		"doctor_id":   req.DoctorId,
		"attachments": req.Attachments,
		"created_at":  time.Now().Format("2006/01/02"),
		"updated_at":  time.Now().Format("2006/01/02"),
		"deleted_at":  0,
	})
	if err != nil {
		return &pb.AddMedicalReportRes{
			Message: "Medical report creation failed",
		}, err
	}

	return &pb.AddMedicalReportRes{
		Message: "Medical report created successfully",
	}, err
}

func (r *medicalRecordRepositoryImpl) GetMedicalReport(ctx context.Context, req *pb.GetMedicalReportReq) (*pb.GetMedicalReportRes, error) {
	var record pb.GetMedicalReportRes
	coll := r.coll.Collection("medical_record")
	cursor, err := coll.Find(ctx, bson.M{
		"$and": bson.D{
			{Key: "user_id", Value: req.UserId},
			{Key: "deleted_at", Value: 0},
		},
	})
	if err != nil {
		fmt.Println("dfadfjkasdfhadsfudshfdsgfgdsyf")
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var doc pb.MedicalReport
		err := cursor.Decode(&doc)
		if err != nil {
			fmt.Println("Error decoding document: ", err)
			return nil, err
		}
		record.MedicalReport = append(record.MedicalReport, &doc)
	}
	fmt.Println(&record)
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	fmt.Println(&record)
	return &record, nil
}

func (r *medicalRecordRepositoryImpl) GetMedicalReportById(ctx context.Context, req *pb.GetMedicalReportByIdReq) (*pb.GetMedicalReportByIdRes, error) {
	var record pb.GetMedicalReportByIdRes
	coll := r.coll.Collection("medical_record")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"_id": req.Id}, {"deleted_at": 0}}}).Decode(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *medicalRecordRepositoryImpl) UpdateMedicalReport(ctx context.Context, req *pb.UpdateMedicalReportReq) (*pb.UpdateMedicalReportRes, error) {
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "record_type", Value: req.RecordType},
			{Key: "description", Value: req.Description},
			{Key: "doctor_id", Value: req.DoctorId},
			{Key: "attachments", Value: req.Attachments},
			{Key: "updated_at", Value: time.Now().Format("2006/01/02")},
		}},
	}
	coll := r.coll.Collection("medical_record")

	filter := bson.D{{Key: "_id", Value: req.Id}, {Key: "deleted_at", Value: 0}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return &pb.UpdateMedicalReportRes{
			Message: false,
		}, err
	}

	return &pb.UpdateMedicalReportRes{
		Message: true,
	}, nil
}

func (r *medicalRecordRepositoryImpl) DeleteMedicalReport(ctx context.Context, req *pb.DeleteMedicalReportReq) (*pb.DeleteMedicalReportRes, error) {
	coll := r.coll.Collection("medical_record")

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now().Unix()}}}}

	filter := bson.D{{Key: "_id", Value: req.Id}, {Key: "deleted_at", Value: 0}}

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
