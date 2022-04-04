package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	handler "github.com/thewolf27/wolf-notebot/internal/transport/http/v1"
)

type Server struct {
	httpSrv *http.Server
	handler *handler.Handler
	Engine  *gin.Engine
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{
		handler: handler,
		Engine:  gin.Default(),
	}
}

func (s *Server) Serve(ctx context.Context, port string) {
	s.handler.Init(s.Engine)

	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.Engine,
	}

	go s.shutdownOnContextDone(ctx)

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Could not start listener ", err)
	}
}

func (s *Server) shutdownOnContextDone(ctx context.Context) {
	<-ctx.Done()

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown: ", err)
	}
	cancel()
}
