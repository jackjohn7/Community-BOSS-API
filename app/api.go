package main

import (
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
  "github.com/gofiber/fiber/v2"
)

func main() {
  // UNIX Time is faster and smaller than most timestamps
  zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

  log.Print("Hello, world")

  app := fiber.New()

  // serve up greeting page
  app.Static("/", "./public/pages")

  app.Listen(":3000")
}
