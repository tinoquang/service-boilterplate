package handler

import (
	"github.com/tinoquang/service-boilerplate/pkg/config"
	"github.com/tinoquang/service-boilerplate/pkg/services"
)

type APIHandler struct {
	cfg *config.Config

	svc services.Services
}

func New(cfg *config.Config, svc services.Services) *APIHandler {
	return &APIHandler{
		cfg: cfg,
		svc: svc,
	}
}
