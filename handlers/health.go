package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// @Summary Show the status of server - Ping/Pong
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /ping [get]
func HandleHealthCheck(c *fiber.Ctx) error {
	log.Debug().Msg("Calling HandleHealthCheck()")
	return c.SendString("pong")
}
