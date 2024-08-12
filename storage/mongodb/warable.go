package mongodb

import (
	"context"
	pb "health-service/genproto/health"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WearableRepository interface {
	AddWearableData(ctx context.Context, Req *pb.AddWearableDataReq) (*pb.AddWearableDataRes, error)
	GetWearableData(ctx context.Context, req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error)
	GetWearableDataById(ctx context.Context, req *pb.GetWearableDataByIdReq) (*pb.GetWearableDataByIdRes, error)
	UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataReq) (*pb.UpdateWearableDataRes, error)
	DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataReq) (*pb.DeleteWearableDataRes, error)
}

type wearableRepositoryImpl struct {
	coll *mongo.Database
}

func NewWearableRepository(db *mongo.Database) WearableRepository {
	return &wearableRepositoryImpl{coll: db}
}

func (r *wearableRepositoryImpl) AddWearableData(ctx context.Context, req *pb.AddWearableDataReq) (*pb.AddWearableDataRes, error) {
	coll := r.coll.Collection("wareable")

	_, err := coll.InsertOne(ctx, bson.M{
		"_id":    uuid.NewString(),
		"user_id": req.UserId,
		"device_type": req.DeviceType,
		"data_type": req.DataType,
		"data_value": req.DataValue,
		"recorded_timestamp": time.Now().Format("2006-01-02-03T15:04:05Z"),
		"created_at": time.Now().Format("2006/01/02"),
        "updated_at": time.Now().Format("2006/01/02"),
		"deleted_at": 0,
	})
	if err!= nil {
        return &pb.AddWearableDataRes{
            Message: false,
        }, err
    }

	return &pb.AddWearableDataRes{
        Message: true,
    }, nil
}

func (r *wearableRepositoryImpl) GetWearableData(ctx context.Context, req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error) {
    var data pb.GetWearableDataRes
	coll := r.coll.Collection("wareable")

	cursor, err := coll.Find(ctx, bson.M{"$and": []bson.M{{"user_id": req.UserId}, {"deleted_at": 0}}})
	if err!= nil {
        return nil, err
    }

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc pb.Warable
        err := cursor.Decode(&doc)
        if err!= nil {
            return nil, err
        }
		doc.FirstName = req.FirstName
		doc.LastName = req.LastName
        data.Warable = append(data.Warable, &doc)
	}

	if err = cursor.Err(); err!= nil {
        return nil, err
    }

	return &data, nil
}

func (r *wearableRepositoryImpl) GetWearableDataById(ctx context.Context, req *pb.GetWearableDataByIdReq) (*pb.GetWearableDataByIdRes, error) {
	var data pb.GetWearableDataByIdRes
	coll := r.coll.Collection("wareable")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"_id": req.Id}, {"deleted_at": 0}}}).Decode(&data)
	if err!= nil {
        return nil, err
    }

	return &data, nil
}

func (r *wearableRepositoryImpl) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataReq) (*pb.UpdateWearableDataRes, error) {
	coll := r.coll.Collection("wareable")
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "device_type", Value: req.DeviceType},
            {Key: "data_type", Value: req.DataType},
            {Key: "data_value", Value: req.DataValue},
            {Key: "updated_at", Value: time.Now()},
        }},
	}

	filter := bson.D{
        {Key: "_id", Value: req.Id},
		{Key: "deleted_at", Value: 0},
    }

	_, err := coll.UpdateOne(ctx, filter, update)
	if err!= nil {
        return &pb.UpdateWearableDataRes{
            Message: false,
        }, err
    }

	return &pb.UpdateWearableDataRes{
        Message: true,
    }, nil
}

func (r *wearableRepositoryImpl) DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataReq) (*pb.DeleteWearableDataRes, error) {
	coll := r.coll.Collection("wareable")
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deleted_at", Value: time.Now().Unix()},
        }},
	}

	filter := bson.D{
		{Key: "_id", Value: req.Id},
        {Key: "deleted_at", Value: 0},
    }
	_, err := coll.UpdateOne(ctx, filter, update)
	if err!= nil {
        return &pb.DeleteWearableDataRes{
            Message: false,
        }, err
    }
	
	return &pb.DeleteWearableDataRes{
        Message: true,
    }, nil
}