package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tinoquang/service-boilerplate/pkg/ctxlogger"
	"go.uber.org/zap"
)

func (h *APIHandler) Ping(c echo.Context) error {
	reqCtx := c.Request().Context()

	if err := h.svc.Ping(reqCtx); err != nil {
		ctxlogger.Error(reqCtx, "ping error", zap.Error(err))
		return err
	}

	return c.String(http.StatusOK, "pong")
}
