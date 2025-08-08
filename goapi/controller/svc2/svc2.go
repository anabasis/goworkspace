package svc2

import "github.com/gofiber/fiber/v2"

func Req1(c *fiber.Ctx) error {
	// Handle request for Service 1, Request 1
	return c.JSON(&fiber.Map{"message": "Service 2 Request 1"})
}

func Req2(c *fiber.Ctx) error {
	// Handle request for Service 1, Request 1
	return c.JSON(&fiber.Map{"message": "Service 2 Request 2"})
}
