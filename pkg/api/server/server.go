package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/tinoquang/service-boilerplate/pkg/api/router"
	"github.com/tinoquang/service-boilerplate/pkg/config"
)

// Server embeds an Echo instance
type Server struct {
	http.Server

	defaultLogger *zap.Logger
}

// New creates a new server
func New(cfg *config.Config, l *zap.Logger) *Server {
	r := router.New(l)

	return &Server{
		Server: http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.Port),
			Handler: r,
		},
		defaultLogger: l,
	}
}

func (s *Server) Start(parentCtx context.Context) {
	go func() {
		<-parentCtx.Done()

		// close the server
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		s.defaultLogger.Info("Shutting down server")
		if err := s.Shutdown(ctx); err != nil {
			s.defaultLogger.Fatal("Server Shutdown", zap.Error(err))
		}
	}()

	// start the server
	s.defaultLogger.Info("server_start", zap.String("port", s.Addr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.defaultLogger.Fatal("listen: %s\n", zap.Error(err))
		return
	}
	s.defaultLogger.Info("Server closed")
}
