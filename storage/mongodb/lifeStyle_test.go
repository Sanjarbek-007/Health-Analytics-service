package mongodb

import (
	"context"
	"fmt"
	pb "health-service/genproto/health"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLifeStyleData(t *testing.T) {
  db, err := ConnectMongoDB()
  if err != nil {
    t.Fatal(err)
  }
  defer db.Client().Disconnect(context.Background())

  repo := NewLifeStyleRepository(db)
  test := &pb.AddLifeStyleDataReq{
    UserId:   "test_user_id",
	DataType: "test_data_type",
    DataValue: "test_data_value",
  }
  resp, err := repo.AddLifeStyleData(context.Background(), test)
  if err != nil {
    t.Fatal(err)
  }
  assert.Equal(t, true, resp.Id)
}

func TestGetLifeStyleData(t *testing.T) {
	mDB, err := ConnectMongoDB()
	assert.NoError(t, err)

	repo := NewLifeStyleRepository(mDB)


	res, err := repo.GetLifeStyleData(context.Background(), &pb.GetLifeStyleDataReq{UserId: "test_user_id"})
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetLifeStyleDataById(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewLifeStyleRepository(mDB)


    res, err := repo.GetLifeStyleDataById(context.Background(), &pb.GetLifeStyleDataByIdReq{Id: "bcafedd4-5780-4e29-8300-fc5cb58a60e1"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestUpdateLifeStyleData(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewLifeStyleRepository(mDB)


    res, err := repo.UpdateLifeStyleData(context.Background(), &pb.UpdateLifeStyleDataReq{Id: "bcafedd4-5780-4e29-8300-fc5cb58a60e1", DataType: "test_data_type", DataValue: "new_test_data_value"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.Equal(t, true, res.Message)
}

func TestDeleteLifeStyleData(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewLifeStyleRepository(mDB)


    res, err := repo.DeleteLifeStyleData(context.Background(), &pb.DeleteLifeStyleDataReq{Id: "bcafedd4-5780-4e29-8300-fc5cb58a60e1"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.Equal(t, true, res.Message)
}