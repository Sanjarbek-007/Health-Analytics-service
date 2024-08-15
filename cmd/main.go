package main

import (
	"context"
	"encoding/json"
	"fmt"
	"health-service/config"
	"health-service/genproto/health"
	kafka "health-service/kafka/consumer"
	"health-service/logs"
	"health-service/service"
	"health-service/storage"
	"log"
	"net"

	mongodb "health-service/storage/mongodb"
	rediss "health-service/storage/redis"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var Mdb *mongo.Database
var Rdb *redis.Client
var err error

func main() {
	logger := logs.NewLogger()
	cfg := config.Load()

	Mdb, err = mongodb.ConnectMongoDB()
	if err != nil {
		logger.Error("Failed to connect to MongoDB", "error", err)
	}

	Rdb = rediss.ConnectRDB()

	storage := storage.NewStorage(Mdb, Rdb)

	listener, err := net.Listen("tcp", cfg.HEALTH_SERVICE)
	if err != nil {
		logger.Error("Failed to listen", "error", err)
	}
	defer listener.Close()

	serviceHealth := service.NewHealthService(logger, storage)
	serviceLifeStyle := service.NewLifeStyleService(logger, storage)
	serviceMedicalRecord := service.NewMedicalService(logger, storage)
	serviceWearable := service.NewWearableService(logger, storage)

	s := grpc.NewServer()

	health.RegisterHealthCheckServer(s, serviceHealth)
	health.RegisterLifeStyleServer(s, serviceLifeStyle)
	health.RegisterMedicalRecordServer(s, serviceMedicalRecord)
	health.RegisterWearableServer(s, serviceWearable)

	log.Printf("Server run: %v", cfg.HEALTH_SERVICE)

	reader, err := kafka.NewKafkaConsumInit([]string{"kafka:9092"}, "health", "group")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	go func() {
		reader.ComsumeMessages(ComsumeMessage)
	}()

	wear, err := kafka.NewKafkaConsumInit([]string{"kafka:9092"}, "werable", "group")
	if err != nil {
		log.Fatal(err)
	}
	defer wear.Close()

	go func() {
		wear.ComsumeMessages(WerableCreateMessage)
	}()

	fmt.Println("everything is ok")

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
	// if err = s.Serve(listener); err != nil {
	// 	logger.Error(err.Error())
	// 	panic(err)
	// }
}

func ComsumeMessage(message []byte) {
	fmt.Println("Consume message: ", string(message))

	var req health.GenerateHealthRecommendationsReq
	err = json.Unmarshal(message, &req)
	if err != nil {
		fmt.Printf("error unmarshalling message: %v\n", err)
		return
	}
	storage := storage.NewStorage(Mdb, Rdb)

	_, err = storage.HealthRepository().GenerateHealthRecommendations(context.Background(), &req)
	if err != nil {
		fmt.Printf("error delete language: %v\n", err)
		return
	}
}

func WerableCreateMessage(message []byte) {
	fmt.Println("Wear Message: ", string(message))

	var req health.AddWearableDataReq

	err = json.Unmarshal(message, &req)
	if err != nil {
		fmt.Printf("error unmarshalling message: %v\n", err)
		return
	}

	storage := storage.NewStorage(Mdb, Rdb)

	_, err = storage.WearableRepository().AddWearableData(context.Background(), &req)
	if err != nil {
		fmt.Printf("error delete language: %v\n", err)
		return
	}
}
