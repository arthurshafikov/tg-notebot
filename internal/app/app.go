package app

import (
	"context"
	"flag"
	"log"

	"github.com/arthurshafikov/tg-notebot/internal/config"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
	"github.com/arthurshafikov/tg-notebot/internal/repository/mongodb"
	"github.com/arthurshafikov/tg-notebot/internal/services"
	server "github.com/arthurshafikov/tg-notebot/internal/transport/http"
	handler "github.com/arthurshafikov/tg-notebot/internal/transport/http/v1"
)

var (
	envFilePath string
)

func init() {
	flag.StringVar(&envFilePath, "env", "./", "Path to .env file folder")
}

func Run() {
	ctx := context.Background()
	config := config.NewConfig(envFilePath)

	mongo, err := mongodb.NewMongoDB(ctx, mongodb.Config{
		Host:     config.DatabaseConfig.Host,
		Port:     config.DatabaseConfig.Port,
		Username: config.DatabaseConfig.Username,
		Password: config.DatabaseConfig.Password,
	})
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository.NewRepository(mongo)

	services := services.NewServices(services.Deps{
		Repository: repository,
	})

	handler := handler.NewHandler(ctx, services)
	s := server.NewServer(handler)
	s.Serve(ctx, config.AppConfig.HTTPPort)
}
