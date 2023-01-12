package middlewares

import (
	"github.com/brpaz/echozap"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tinoquang/service-boilerplate/pkg/ctxlogger"
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

	// setup CORS, TODO: AllowOrigins should be configurable
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderCookie,
		},

		// Set to true if your using cookies across subdomains, default is false
		// AllowCredentials: true,
	}))

	// setup context logger
	e.Use(setupCtxLogger(l))
}

func setupCtxLogger(l *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger := l.With(zap.String("request_id", c.Response().Header().Get(echo.HeaderXRequestID)))

			currentCtx := c.Request().Context()
			c.SetRequest(c.Request().WithContext(ctxlogger.ToContext(currentCtx, logger)))

			return next(c)
		}
	}
}
