package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	log.Fatal(server.Listen(""))
}
