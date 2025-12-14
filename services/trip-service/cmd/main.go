package main

import (
	"context"
	"go-ride/services/trip-service/internal/domain"
	"go-ride/services/trip-service/internal/infrastructure/repository"
	"go-ride/services/trip-service/internal/service"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: "54",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}
	log.Println(t)

	for {
		time.Sleep(time.Second)
	}
}
