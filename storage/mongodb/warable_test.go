package mongodb

import (
	"context"
	"fmt"
	pb "health-service/genproto/health"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWearableData(t *testing.T) {
  db, err := ConnectMongoDB()
  if err != nil {
    t.Fatal(err)
  }
  defer db.Client().Disconnect(context.Background())

  repo := NewWearableRepository(db)
  test := &pb.AddWearableDataReq{
    UserId:   "test_user_id",
	DataType: "test_data_type",
    DataValue: "test_data_value",
  }
  resp, err := repo.AddWearableData(context.Background(), test)
  if err != nil {
    t.Fatal(err)
  }
  assert.Equal(t, true, resp.Message)
}

func TestGetWearableData(t *testing.T) {
	mDB, err := ConnectMongoDB()
	assert.NoError(t, err)

	repo := NewWearableRepository(mDB)


	res, err := repo.GetWearableData(context.Background(), &pb.GetWearableDataReq{UserId: "test_user_id"})
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetWearableDataById(t *testing.T) {
	mDB, err := ConnectMongoDB()
	assert.NoError(t, err)

	repo := NewWearableRepository(mDB)

	res, err := repo.GetWearableDataById(context.Background(), &pb.GetWearableDataByIdReq{Id: "572a2033-a716-4ba2-aec4-955eee49e3cc"})
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUpdateWearableData(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewWearableRepository(mDB)

    res, err := repo.UpdateWearableData(context.Background(), &pb.UpdateWearableDataReq{Id: "572a2033-a716-4ba2-aec4-955eee49e3cc", DataType: "test_updated_data_type", DataValue: "test_updated_data_value"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.Equal(t, true, res.Message)
}

func TestDeleteWearableData(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewWearableRepository(mDB)

    res, err := repo.DeleteWearableData(context.Background(), &pb.DeleteWearableDataReq{Id: "572a2033-a716-4ba2-aec4-955eee49e3cc"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.Equal(t, true, res.Message)
}