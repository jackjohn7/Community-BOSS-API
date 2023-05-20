package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
  // UNIX Time is faster and smaller than most timestamps
  zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

  log.Print("Hello, world")

  app := fiber.New()

  // use bundled rate limiter (at least for now)
  app.Use(limiter.New(limiter.Config {
    Max: 20,
    Expiration: 30 * time.Second,
  }))

  // TODO: Implement custom logger (saves logs to database)
  // app.Use(logger)

  // serve up greeting page
  app.Static("/", "./public/pages")

  app.Listen(":3000")
}
