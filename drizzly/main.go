package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) (rawr error) {
			switch os.Getenv("") {
			case "debug":
			}

			return
		},
	})

	log.Fatal(server.Listen(""))
}
