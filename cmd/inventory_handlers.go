package main

import (
	"context"
	"inventory/cmd/utils"
	"inventory/models"

	"github.com/gofiber/fiber/v2"
)

func (app *Config) addInventory(c *fiber.Ctx) error {

	var reqPayload models.Item

	if err := c.BodyParser(&reqPayload); err != nil {
		return utils.SendResponse(
			"failed to parse request body",
			nil,
			c,
			fiber.StatusBadRequest,
		)
	}

	res, err := app.Inventory.InsertOne(context.Background(), reqPayload)

	if err != nil {
		return utils.SendResponse(
			"failed to add new inventory item",
			nil,
			c,
			fiber.StatusBadRequest,
		)
	}

	return utils.SendResponse("", map[string]any{
		"item_id": res.InsertedID,
	}, c, fiber.StatusOK)
}

func (app *Config) getInventory(c *fiber.Ctx) error {
	return c.SendString("getInventory")
}
