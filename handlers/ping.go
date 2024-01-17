package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Show the status of server - Ping/Pong
// @Description get the status of server.
// @Tags ping
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/ping [get]
func HandlePingCheck(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
