package middlewares

import (
	"net/http"

	"github.com/brpaz/echozap"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterCommon(e *echo.Echo, l *zap.Logger) {
	// setup logger first
	e.Use(echozap.ZapLogger(l))

	// common middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.RequestID())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CSRF())

	// setup CORS, TODO: AllowOrigins should be configurable
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
}
