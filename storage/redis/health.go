package redis

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"health-service/genproto/health"

// 	"github.com/redis/go-redis/v9"
// )

// type RealTimeRepository interface {
// 	AddRealTimeData(ctx context.Context, in *health.AddRealTimeDataReq) (*health.AddRealTimeDataRes, error)
// 	GetRealTimeData(ctx context.Context, in *health.GetRealTimeDataReq) (*health.GetRealTimeDataRes, error)
// }

// type realTimeRepositoryImpl struct {
//     redis *redis.Client
// }

// type RealTimeData struct {
// 	UserId     string  `json:"user_id"`
// 	Data   string  `json:"data_type"`
// }

// func NewRealTimeRepository(client *redis.Client) RealTimeRepository {
//     return &realTimeRepositoryImpl{redis: client}
// }

// func (b *realTimeRepositoryImpl) AddRealTimeData(ctx context.Context, in *health.AddRealTimeDataReq) (*health.AddRealTimeDataRes, error){
// 	data, err := json.Marshal(in)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = b.redis.RPush(ctx, in.UserId, data).Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &health.AddRealTimeDataRes{
// 		Message: true,
// 	}, nil
// }


// func (b *realTimeRepositoryImpl) GetRealTimeData(ctx context.Context, in *health.GetRealTimeDataReq) (*health.GetRealTimeDataRes, error) {
//     // Fetch data from Redis (assuming you need the last entry)
//     data, err := b.redis.LIndex(ctx, in.UserId, -1).Result()
//     if err != nil {
//         return nil, fmt.Errorf("failed to fetch data from Redis: %w", err)
//     }

//     // Unmarshal the data into RealTimeData struct
//     var realTimeData health.GetRealTimeDataRes
//     err = json.Unmarshal([]byte(data), &realTimeData)
//     if err != nil {
//         return nil, fmt.Errorf("failed to unmarshal data: %w", err)
//     }

//     // Return the result
//     return &health.GetRealTimeDataRes{
//         UserId: realTimeData.UserId,
// 		RecommendationType: realTimeData.RecommendationType,
// 		DescriptionRecommendation: realTimeData.DescriptionRecommendation,
//         Priority: realTimeData.Priority,
// 		RecordType: realTimeData.RecordType,
//         DescriptionRecord: realTimeData.DescriptionRecord,
//         DataType: realTimeData.DataType,
// 		DataValue: realTimeData.DataValue,
// 		DeviceType: realTimeData.DeviceType,
// 		DeviceDataType: realTimeData.DeviceDataType,
//         DeviceDataValue: realTimeData.DeviceDataValue,
//     }, nil
// }
