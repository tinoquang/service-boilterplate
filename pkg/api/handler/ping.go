package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *APIHandler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
