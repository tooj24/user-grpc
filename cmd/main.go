package main

import (
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	dbConfig "grpc-clean/internal/db"
	"grpc-clean/internal/models"
	interfaces "grpc-clean/pkg/v1"
	handler "grpc-clean/pkg/v1/handler/grpc"
	"grpc-clean/pkg/v1/repository"
	"grpc-clean/pkg/v1/usecase"
	"log"
	"net"
)

func main() {
	fx.New(
		fx.Provide(
			dbConfig.DatabaseSetup,
			grpc.NewServer,
			repository.NewUserRepositoryImpl,
			usecase.NewUserUseCaseImpl,
		),
		fx.Invoke(InitServerRPC),
		fx.Invoke(func(server *grpc.Server, db *gorm.DB) {
			migrations(db)
			reflection.Register(server)
			lis, err := net.Listen("tcp", ":5001")
			if err != nil {
				log.Fatalf("Error starting the server: %v", err)
			}

			server.Serve(lis)
		}),
	).Run()
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}

func InitServerRPC(server *grpc.Server, usecase interfaces.UserUseCase) {
	handler.NewServer(server, usecase)
}
