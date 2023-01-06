package router

import (
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"

	"github.com/tinoquang/service-boilerplate/pkg/api/handler"
)

type Router struct {
	*echo.Echo
}

func New(l *zap.Logger) *Router {
	// Create a new Echo instance
	e := echo.New()

	// Add custom validator
	e.Binder = newCustomBinder()

	r := &Router{
		Echo: e,
	}

	// register middlewares

	// register routes
	r.registerRoutes(l)
	return r
}

func (r *Router) registerRoutes(l *zap.Logger) {
	h := handler.New(l)

	r.GET("/ping", h.Ping)
}
