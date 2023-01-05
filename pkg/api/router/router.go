package router

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	r.registerRoutes()

	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}"}` + "\n",
	}))

	return r
}

func (r *Router) registerRoutes() {
	r.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
