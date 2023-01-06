package router

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
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
	r.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
