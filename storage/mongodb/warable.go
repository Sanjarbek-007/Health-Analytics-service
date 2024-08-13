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
		"userId": req.UserId,
		"deviceType": req.DeviceType,
		"dataType": req.DataType,
		"dataValue": req.DataValue,
		"recordedTimestamp": time.Now().Format("2006-01-02-03T15:04:05Z"),
		"createdAt": time.Now().Format("2006/01/02"),
        "updatedAt": time.Now().Format("2006/01/02"),
		"deletedAt": 0,
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

	cursor, err := coll.Find(ctx, bson.M{"$and": []bson.M{{"userId": req.UserId}, {"deletedAt": 0}}})
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

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"_id": req.Id}, {"deletedAt": 0}}}).Decode(&data)
	if err!= nil {
        return nil, err
    }

	return &data, nil
}

func (r *wearableRepositoryImpl) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataReq) (*pb.UpdateWearableDataRes, error) {
	coll := r.coll.Collection("wareable")
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deviceType", Value: req.DeviceType},
            {Key: "dataType", Value: req.DataType},
            {Key: "dataValue", Value: req.DataValue},
            {Key: "updatedAt", Value: time.Now()},
        }},
	}

	filter := bson.D{
        {Key: "_id", Value: req.Id},
		{Key: "deletedAt", Value: 0},
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
			{Key: "deletedAt", Value: time.Now().Unix()},
        }},
	}

	filter := bson.D{
		{Key: "_id", Value: req.Id},
        {Key: "deletedAt", Value: 0},
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