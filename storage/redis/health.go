package rediss

import (
	"context"
	"encoding/json"
	"fmt"

	pb "health-service/genproto/health"

	"github.com/redis/go-redis/v9"
)

type HealthMonito interface {
	CreateHealthMonitor(ctx context.Context, req *pb.CreateHealthMonitorReq) error
	GetHealthMonitor(ctx context.Context, userID *pb.UserId) (*pb.GetHealthMonitorsRes, error)
}

type healthMonitorImpl struct {
	RedisClient *redis.Client
}

func NewHealthMonitorRepo(redisClient *redis.Client) HealthMonito {
	return &healthMonitorImpl{
		RedisClient: redisClient,
	}
}

func (r *healthMonitorImpl) CreateHealthMonitor(ctx context.Context, req *pb.CreateHealthMonitorReq) error {
	value, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal CreateHealthMonitorReq: %w", err)
	}
	err = r.RedisClient.Set(ctx, req.GetUserId(), value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to store data in Redis: %w", err)
	}

	return nil
}

func (r *healthMonitorImpl) GetHealthMonitor(ctx context.Context, userID *pb.UserId) (*pb.GetHealthMonitorsRes, error) {
	val, err := r.RedisClient.Get(ctx, userID.GetUserId()).Result()
	fmt.Println(userID.UserId)
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no data found for user_id: %s", userID.GetUserId())
		}
		return nil, fmt.Errorf("failed to get data from Redis: %w", err)
	}

	var res pb.GetHealthMonitorsRes
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &res, nil
}
