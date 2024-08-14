package mongodb

import (
	"context"
	"fmt"
	pb "health-service/genproto/health"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMedicalReport(t *testing.T) {
  db, err := ConnectMongoDB()
  if err != nil {
    t.Fatal(err)
  }
  defer db.Client().Disconnect(context.Background())

  repo := NewMedicalRecordRepository(db)
  test := &pb.AddMedicalReportReq{
    UserId:   "test_user_id",
	RecordType: "test_record_type",
    Description: "test_description",
    DoctorId: "test_doctor_id",
    Attachments: []*pb.Attachment{{FileUrl: "test_attachment1"}, {FileUrl: "test_attachment2"}},
  }
  resp, err := repo.AddMedicalReport(context.Background(), test)
  if err != nil {
    t.Fatal(err)
  }
  assert.Equal(t, "Medical report created successfully", resp.Id)
}

func TestGetMedicalReport(t *testing.T) {
	mDB, err := ConnectMongoDB()
	assert.NoError(t, err)

	repo := NewMedicalRecordRepository(mDB)


	res, err := repo.GetMedicalReport(context.Background(), &pb.GetMedicalReportReq{UserId: "test_user_id"})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetMedicalReportById(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewMedicalRecordRepository(mDB)


    res, err := repo.GetMedicalReportById(context.Background(), &pb.GetMedicalReportByIdReq{Id: "271e385a-9364-43b7-aab8-c747c69bae5f"})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestUpdateMedicalReport(t *testing.T) {
	mDB, err := ConnectMongoDB()
    assert.NoError(t, err)

    repo := NewMedicalRecordRepository(mDB)


    res, err := repo.UpdateMedicalReport(context.Background(), &pb.UpdateMedicalReportReq{Id: "47199842-e970-4891-b2db-ee5cc90731eb", RecordType: "updated_record_type", Description: "updated_description", DoctorId: "updated_doctor_id", Attachments: []*pb.Attachment{{FileUrl: "updated_attachment1"}, {FileUrl: "updated_attachment2"}}})
    fmt.Println(res)
    assert.NoError(t, err)
    assert.Equal(t, true, res.Message)
}

func TestDeleteMedicalReport(t *testing.T) {
	mDB, err := ConnectMongoDB()
	assert.NoError(t, err)

	repo := NewMedicalRecordRepository(mDB)

	res, err := repo.DeleteMedicalReport(context.Background(), &pb.DeleteMedicalReportReq{Id: "47199842-e970-4891-b2db-ee5cc90731eb"})
	fmt.Println(res)
	assert.NoError(t, err)
	assert.Equal(t, true, res.Message)
}