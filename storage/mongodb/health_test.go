package mongodb

import (
	"context"
	"fmt"
	pb "health-service/genproto/health"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHealthRecommendations(t *testing.T) {
  db, err := ConnectMongoDB()
  if err != nil {
    t.Fatal(err)
  }
  defer db.Client().Disconnect(context.Background())

  repo := NewHealthRepository(db)
  account := &pb.GenerateHealthRecommendationsReq{
    UserId:   "test_user_id",
    RecommendationType: "test_recommendation_type",
    Description: "test_description",
    Priority: 1,
  }
  resp, err := repo.GenerateHealthRecommendations(context.Background(), account)
  if err != nil {
    t.Fatal(err)
  }
  assert.Equal(t, true, resp.Message)
}

func TestGetRealtimeHealthMonitoring(t *testing.T) {
  mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewHealthRepository(mDB)


    res, err := repo.GetRealtimeHealthMonitoring(context.Background(), &pb.GetRealtimeHealthMonitoringReq{UserId: "test_user_id"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetDailyHealthSummary(t *testing.T) {
  mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewHealthRepository(mDB)


    res, err := repo.GetDailyHealthSummary(context.Background(), &pb.GetDailyHealthSummaryReq{UserId: "test_user_id", Date:"2024/08/11"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetWeeklyHealthSummary(t *testing.T) {
  mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewHealthRepository(mDB)


    res, err := repo.GetWeeklyHealthSummary(context.Background(), &pb.GetWeeklyHealthSummaryReq{UserId: "test_user_id", StartDate:"2024/08/11", EndDate:"2024/08/18"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}