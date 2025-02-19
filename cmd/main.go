package main

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		logger.Log.Debug("GEGEGE")
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%s", cfg.GetConfig().Server.Port))
}
