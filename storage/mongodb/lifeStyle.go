package mongodb

import (
	"context"
	pb "health-service/genproto/health"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LifeStyleRepository interface {
	AddLifeStyleData(ctx context.Context, Req *pb.AddLifeStyleDataReq) (*pb.AddLifeStyleDataRes, error)
	GetLifeStyleData(ctx context.Context, req *pb.GetLifeStyleDataReq) (*pb.GetLifeStyleDataRes, error)
	GetLifeStyleDataById(ctx context.Context, req *pb.GetLifeStyleDataByIdReq) (*pb.GetLifeStyleDataByIdRes, error)
	UpdateLifeStyleData(ctx context.Context, req *pb.UpdateLifeStyleDataReq) (*pb.UpdateLifeStyleDataRes, error)
	DeleteLifeStyleData(ctx context.Context, req *pb.DeleteLifeStyleDataReq) (*pb.DeleteLifeStyleDataRes, error)
}

type lifeStyleRepositoryImpl struct {
	coll *mongo.Database
}

func NewLifeStyleRepository(db *mongo.Database) LifeStyleRepository {
	return &lifeStyleRepositoryImpl{coll: db}
}

func (r *lifeStyleRepositoryImpl) AddLifeStyleData(ctx context.Context, req *pb.AddLifeStyleDataReq) (*pb.AddLifeStyleDataRes, error) {
    coll := r.coll.Collection("lifestyle")

	_, err := coll.InsertOne(ctx, bson.M{
		"_id":    uuid.NewString(),
        "user_id": req.UserId,
        "data_type": req.DataType,
		"data_value": req.DataValue,
		"analysis_date": time.Now().Format("2006/01/02"),
		"created_at": time.Now().Format("2006/01/02"),
        "updated_at": time.Now().Format("2006/01/02"),
        "deleted_at": 0,
	})

	if err!= nil {
        return &pb.AddLifeStyleDataRes{
            Message: false,
        }, err
    }

	return &pb.AddLifeStyleDataRes{
		Message: true,
	}, err
}

func (r *lifeStyleRepositoryImpl) GetLifeStyleData(ctx context.Context, req *pb.GetLifeStyleDataReq) (*pb.GetLifeStyleDataRes, error) {
    var user pb.GetLifeStyleDataRes

    coll := r.coll.Collection("lifestyle")

	cursor, err := coll.Find(ctx, bson.M{"$and": []bson.M{{"user_id": req.UserId}, {"deleted_at": 0}}})
	if err!= nil {
        return nil, err
    }

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var doc pb.GetLifeStyle
        if err := cursor.Decode(&doc); err!= nil {
            return nil, err
        }
        doc.FirstName = req.FirstName
        doc.LastName = req.LastName
        user.LifeStyle = append(user.LifeStyle, &doc)
	}
	if err := cursor.Err(); err!= nil {
        return nil, err
    }

	return &user, nil
}

func (r *lifeStyleRepositoryImpl) GetLifeStyleDataById(ctx context.Context, req *pb.GetLifeStyleDataByIdReq) (*pb.GetLifeStyleDataByIdRes, error) {
    var user pb.GetLifeStyleDataByIdRes
    coll := r.coll.Collection("lifestyle")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"_id": req.Id}, {"deleted_at": 0}}}).Decode(&user)
	if err!= nil {
        return nil, err
    }

	return &user, nil
}

func (r *lifeStyleRepositoryImpl) UpdateLifeStyleData(ctx context.Context, req *pb.UpdateLifeStyleDataReq) (*pb.UpdateLifeStyleDataRes, error) {
    coll := r.coll.Collection("lifestyle")
	update := bson.D{
        {Key: "$set", Value: bson.D{
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
        return &pb.UpdateLifeStyleDataRes{
            Message: false,
        }, err
    }

    return &pb.UpdateLifeStyleDataRes{
		Message: true,
	}, err
}

func (r *lifeStyleRepositoryImpl) DeleteLifeStyleData(ctx context.Context, req *pb.DeleteLifeStyleDataReq) (*pb.DeleteLifeStyleDataRes, error) {
	filter := bson.D{
        {Key: "_id", Value: req.Id},
        {Key: "deleted_at", Value: 0},
    }
    coll := r.coll.Collection("lifestyle")

    _, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now().Unix()}}}})

    if err!= nil {
        return &pb.DeleteLifeStyleDataRes{
            Message: false,
        }, err
    }

    return &pb.DeleteLifeStyleDataRes{
        Message: true,
    }, nil
}