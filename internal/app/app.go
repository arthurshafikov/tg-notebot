package app

import (
	"context"

	"github.com/arthurshafikov/wolf-notebot/internal/repository"
	"github.com/arthurshafikov/wolf-notebot/internal/repository/mongodb"
	"github.com/arthurshafikov/wolf-notebot/internal/services"
	server "github.com/arthurshafikov/wolf-notebot/internal/transport/http"
	handler "github.com/arthurshafikov/wolf-notebot/internal/transport/http/v1"
)

func Run() {
	ctx := context.Background()

	mongo, err := mongodb.NewMongoDB(ctx, mongodb.Config{
		Host:     "localhost",
		Port:     "27017",
		Username: "root",
		Password: "supersecret", // todo .env
	})
	if err != nil {
		panic(err) // temp
	}

	repository := repository.NewRepository(mongo)

	services := services.NewServices(services.Deps{
		Repository: repository,
	})

	handler := handler.NewHandler(ctx, services)
	s := server.NewServer(handler)
	s.Serve(ctx, "8123")
}
