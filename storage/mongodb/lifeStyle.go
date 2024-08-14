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
    id := uuid.NewString()

	_, err := coll.InsertOne(ctx, bson.M{
		"_id":    id,
        "userId": req.UserId,
        "dataType": req.DataType,
		"dataValue": req.DataValue,
		"analysisDate": time.Now().Format("2006/01/02"),
		"createdAt": time.Now().Format("2006/01/02"),
        "updatedAt": time.Now().Format("2006/01/02"),
        "deletedAt": 0,
	})

	if err!= nil {
        return nil, err
    }

	return &pb.AddLifeStyleDataRes{
		Id: id,
	}, err
}

func (r *lifeStyleRepositoryImpl) GetLifeStyleData(ctx context.Context, req *pb.GetLifeStyleDataReq) (*pb.GetLifeStyleDataRes, error) {
    var user pb.GetLifeStyleDataRes

    coll := r.coll.Collection("lifestyle")

	cursor, err := coll.Find(ctx, bson.M{"$and": []bson.M{{"userId": req.UserId}, {"deletedAt": 0}}})
	if err!= nil {
        return nil, err
    }

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var doc pb.GetLifeStyle
        if err := cursor.Decode(&doc); err!= nil {
            return nil, err
        }
        user.LifeStyle = append(user.LifeStyle, &doc)
	}
	if err := cursor.Err(); err!= nil {
        return nil, err
    }

	return &user, nil
}

func (r *lifeStyleRepositoryImpl) GetLifeStyleDataById(ctx context.Context, req *pb.GetLifeStyleDataByIdReq) (*pb.GetLifeStyleDataByIdRes, error) {
    var user pb.GetLifeStyleByIdRes
    coll := r.coll.Collection("lifestyle")

	err := coll.FindOne(ctx, bson.M{"$and": []bson.M{{"_id": req.Id}, {"deletedAt": 0}}}).Decode(&user)
	if err!= nil {
        return nil, err
    }

	return &pb.GetLifeStyleDataByIdRes{LifeStyle: &user}, nil
}

func (r *lifeStyleRepositoryImpl) UpdateLifeStyleData(ctx context.Context, req *pb.UpdateLifeStyleDataReq) (*pb.UpdateLifeStyleDataRes, error) {
    coll := r.coll.Collection("lifestyle")
	update := bson.D{
        {Key: "$set", Value: bson.D{
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
        return &pb.UpdateLifeStyleDataRes{
            Message: false,
        }, err
    }

    return &pb.UpdateLifeStyleDataRes{
		Message: true,
	}, err
}

func (r *lifeStyleRepositoryImpl) DeleteLifeStyleData(ctx context.Context, req *pb.DeleteLifeStyleDataReq) (*pb.DeleteLifeStyleDataRes, error) {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deletedAt", Value: time.Now().Unix()}}}}

	filter := bson.D{{Key: "_id", Value: req.Id}, {Key: "deletedAt", Value: 0}}
    coll := r.coll.Collection("lifestyle")

    _, err := coll.UpdateOne(ctx, filter, update)

    if err!= nil {
        return &pb.DeleteLifeStyleDataRes{
            Message: false,
        }, err
    }

    return &pb.DeleteLifeStyleDataRes{
        Message: true,
    }, nil
}