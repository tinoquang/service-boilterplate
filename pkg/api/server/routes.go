package server

import (
	"github.com/tinoquang/service-boilerplate/pkg/api/handler"
	"github.com/tinoquang/service-boilerplate/pkg/config"
	"github.com/tinoquang/service-boilerplate/pkg/services"
)

func (s *Server) registerRoutes(cfg *config.Config, svc services.Services) {
	h := handler.New(cfg, svc)

	s.e.GET("/ping", h.Ping)

	// Add more routes here
}
