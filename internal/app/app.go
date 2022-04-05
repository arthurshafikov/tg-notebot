package app

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/repository"
	"github.com/thewolf27/wolf-notebot/internal/repository/mongo"
	"github.com/thewolf27/wolf-notebot/internal/services"
	server "github.com/thewolf27/wolf-notebot/internal/transport/http"
	handler "github.com/thewolf27/wolf-notebot/internal/transport/http/v1"
)

func Run() {
	ctx := context.Background()

	mongo, err := mongo.NewMongoDB(ctx, "localhost", "27017")
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
