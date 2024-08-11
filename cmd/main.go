package main

import (
	"health-service/config"
	"health-service/genproto/health"
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
        logger.Error("Failed to connect to MongoDB", err)
    }

	storage := storage.NewStorage(mdb)

	listener, err := net.Listen("tcp", cfg.PROGRESS_SERVICE)
	if err!= nil {
        logger.Error("Failed to listen", err)
    }
	defer listener.Close()

	serviceHealth := service.NewHealthService(logger, storage)
	serviceLifeStyle := service.NewLifeStyleService(logger, storage)
	serviceMedicalRecord := service.NewMedicalService(logger, storage)



	s := grpc.NewServer()

	health.RegisterHealthCheckServer(s, serviceHealth)
	health.RegisterLifeStyleServer(s, service)
	health.RegisterMedicalRecordServer(s, service)
	health.RegisterWearableServer(s, service)

	log.Printf("Server run: %v", cfg.PROGRESS_SERVICE)
	if err = s.Serve(listener); err != nil{
		logger.Error(err.Error())
		panic(err)
	}


}