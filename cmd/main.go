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

	"google.golang.org/grpc"
)

func main() {
	logger := logs.NewLogger()
	cfg := config.Load()

	mdb, err := mongodb.ConnectMongoDB()
	if err != nil {
		logger.Error("Failed to connect to MongoDB", "error", err)
	}

	storage := storage.NewStorage(mdb)

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
	if err = s.Serve(listener); err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	reader, err := kafka.NewKafkaConsumInit([]string{"kafka:9092"}, "delete", "group")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	go func() {
		reader.ComsumeMessages(ComsumeMessage)
	}()

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}

func ComsumeMessage(message []byte) {
	fmt.Println("Consume message: ", string(message))
	db, err := mongodb.ConnectMongoDB()
	if err != nil {
		panic(err)
	}
	var req health.GenerateHealthRecommendationsReq
	err = json.Unmarshal(message, &req)
	if err != nil {
		fmt.Printf("error unmarshalling message: %v\n", err)
		return
	}
	storage := storage.NewStorage(db)

	_, err = storage.HealthRepository().GenerateHealthRecommendations(context.Background(), &req)
	if err != nil {
		fmt.Printf("error delete language: %v\n", err)
		return
	}
}
