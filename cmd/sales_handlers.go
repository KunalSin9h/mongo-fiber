package main

import (
	"github.com/gofiber/fiber/v2"
)

func (app *Config) addSales(c *fiber.Ctx) error {
	return c.SendString("addSales")
}

func (app *Config) getSales(c *fiber.Ctx) error {
	return c.SendString("getSales")
}
