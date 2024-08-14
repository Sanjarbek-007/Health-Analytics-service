package storage

import (
	"health-service/storage/mongodb"
	rediss "health-service/storage/redis"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	HealthRepository() mongodb.HealthRepository
	LifeStyleRepository() mongodb.LifeStyleRepository
	MedicalRecordRepository() mongodb.MedicalRecordRepository
	WearableRepository() mongodb.WearableRepository
	MonitoringRepository() rediss.HealthMonito
}

type storageImpl struct {
	db  *mongo.Database
	rdb *redis.Client
}

func NewStorage(db *mongo.Database, rdb *redis.Client) IStorage {
	return &storageImpl{db: db, rdb: rdb}
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

func (s *storageImpl) MonitoringRepository() rediss.HealthMonito {
	return rediss.NewHealthMonitorRepo(s.rdb)
}
