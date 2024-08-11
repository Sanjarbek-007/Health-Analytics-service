package storage

import (
	"health-service/storage/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	HealthRepository() mongodb.HealthRepository
	LifeStyleRepository() mongodb.LifeStyleRepository
	MedicalRecordRepository() mongodb.MedicalRecordRepository
	WearableRepository() mongodb.WearableRepository
}

type storageImpl struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) IStorage {
	return &storageImpl{db: db}
}

func (s *storageImpl) HealthRepository() mongodb.HealthRepository {
	return mongodb.NewHealthRepository(s.db)
}

func (s *storageImpl) LifeStyleRepository() mongodb.LifeStyleRepository {
	return mongodb.NewLifeStyleRepository(s.db)
}

func (s *storageImpl) MedicalRecordRepository() mongodb.MedicalRecordRepository {
	return mongodb.NewMedicalRecordRepository(s.db)
}

func (s *storageImpl) WearableRepository() mongodb.WearableRepository {
	return mongodb.NewWearableRepository(s.db)
}
