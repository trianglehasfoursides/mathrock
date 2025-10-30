package main

import "github.com/gofiber/fiber/v2"

func index(c *fiber.Ctx) error {
	return c.Send([]byte("Hidup Jokowi!!!"))
}
