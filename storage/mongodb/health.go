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

type HealthRepository interface {
	GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error)
	GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error)
	GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error)
	GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error)
}

type healthRepositoryImpl struct {
	coll *mongo.Database
}

func NewHealthRepository(db *mongo.Database) HealthRepository {
	return &healthRepositoryImpl{coll: db}
}

func (r *healthRepositoryImpl) GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error) {
	coll := r.coll.Collection("health")
	id := uuid.NewString()

	_, err := coll.InsertOne(ctx, bson.M{
		"_id":                id,
		"userId":             req.UserId,
		"recommendationType": req.RecommendationType,
		"description":        req.Description,
		"priority":           req.Priority,
		"createdAt":          time.Now(),
		"updatedAt":          time.Now(),
		"deletedAt":          0,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GenerateHealthRecommendationsRes{
		Id: id,
	}, nil
}

func (r *healthRepositoryImpl) GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error) {
	var user pb.GetRealtimeHealthMonitoringRes
	coll := r.coll.Collection("health")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"userId": req.UserId}, {"deletedAt": 0}, {"createdAt": time.Now()}}}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("realtime health monitoring not found")
	}

	return &user, nil
}

func (r *healthRepositoryImpl) GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error) {
	var summary pb.GetDailyHealthSummaryRes
	coll := r.coll.Collection("health")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"userId": req.UserId}, {"deletedAt": 0}, {"createdAt": req.Date}}}).Decode(&summary)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *healthRepositoryImpl) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error) {
	var summary pb.GetWeeklyHealthSummaryRes
	coll := r.coll.Collection("health")

	cursor, err := coll.Find(ctx, bson.M{
		"$and": []bson.M{
			{"userId": req.UserId},
			{"deletedAt": 0},
			{"createdAt": bson.M{
				"$gte": req.StartDate,
				"$lte": req.EndDate,
			}},
		},
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc pb.HealthRecommendation
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		doc.FirstName = req.FirstName
		doc.LastName = req.LastName
		summary.Recommendations = append(summary.Recommendations, &doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &summary, nil
}
