package services

import (
	"context"

	"github.com/tinoquang/service-boilerplate/pkg/ctxlogger"
)

func (s *services) Ping(ctx context.Context) error {
	ctxlogger.Info(ctx, "ping")
	return nil
}
