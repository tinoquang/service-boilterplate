package services

import (
	"context"

	"github.com/tinoquang/service-boilerplate/pkg/config"
)

type Services interface {
	Ping(ctx context.Context) error
}

type services struct {
	cfg *config.Config
}

func New(cfg *config.Config) *services {
	return &services{
		cfg: cfg,
	}
}
