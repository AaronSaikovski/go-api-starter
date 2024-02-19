package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Show the health status of the API.
// @Description get the status of server.
// @Tags healthcheck
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/health [get]
func HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
