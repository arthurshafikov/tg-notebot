package app

import (
	"context"

	"github.com/thewolf27/wolf-notebot/internal/services"
	server "github.com/thewolf27/wolf-notebot/internal/transport/http"
	handler "github.com/thewolf27/wolf-notebot/internal/transport/http/v1"
)

func Run() {
	ctx := context.Background()

	handler := handler.NewHandler(ctx, services.NewServices(services.Deps{}))
	s := server.NewServer(handler)
	s.Serve(ctx, "8123")
}
