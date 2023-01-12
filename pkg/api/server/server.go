package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/tinoquang/service-boilerplate/pkg/api/middlewares"
	"github.com/tinoquang/service-boilerplate/pkg/config"
	"github.com/tinoquang/service-boilerplate/pkg/ctxlogger"
	"github.com/tinoquang/service-boilerplate/pkg/services"
)

type Server struct {
	cfg *config.Config
	e   *echo.Echo
}

// New creates a new server
func New(cfg *config.Config, l *zap.Logger, svc services.Services) *Server {
	// Create a new Echo instance
	e := echo.New()

	// Add custom validator
	e.Binder = newCustomBinder()

	// register middlewares
	middlewares.RegisterCommon(e, l)

	// register routes
	s := &Server{
		cfg: cfg,
		e:   e,
	}

	s.registerRoutes(cfg, svc)
	return s
}

func (s *Server) Start(parentCtx context.Context) {
	go func() {
		<-parentCtx.Done()

		// close the server
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ctxlogger.Info(parentCtx, "Shutting down server")
		if err := s.e.Shutdown(ctx); err != nil {
			ctxlogger.Error(parentCtx, "Server shutdown error", zap.Error(err))
		}
	}()

	// start the server
	if err := s.e.Start(fmt.Sprintf(":%s", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
		ctxlogger.Error(parentCtx, "Server start error", zap.Error(err))
		return
	}

	ctxlogger.Info(parentCtx, "Server stopped")
}
