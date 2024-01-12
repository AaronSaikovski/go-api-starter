package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Show the status of server - Ping/Pong
// @Description get the status of server.
// @Tags ping
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/ping [get]
func HandlePingCheck(c *fiber.Ctx) error {
	return c.SendString("pong")
}
