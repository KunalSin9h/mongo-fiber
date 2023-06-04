package main

import (
	"context"
	"inventory/cmd/utils"
	"inventory/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddInventory API handler is used to add new inventory item
// to the mongodb database
func (app *Config) addInventory(c *fiber.Ctx) error {

	// Define request payload
	// which the client will send
	var reqPayload models.Item

	// Parse the request body
	if err := c.BodyParser(&reqPayload); err != nil {
		return utils.SendResponse(
			"failed to parse request body",
			nil,
			c,
			fiber.StatusBadRequest,
		)
	}

	// Insert the new inventory item to the mongodb database
	res, err := app.Inventory.InsertOne(context.Background(), reqPayload)

	if err != nil {
		return utils.SendResponse(
			"failed to add new inventory item",
			nil,
			c,
			fiber.StatusBadRequest,
		)
	}

	// Send the response to the client
	return utils.SendResponse("", map[string]any{
		"item_id": res.InsertedID,
	}, c, fiber.StatusOK)
}

// GetInventory API handler is used to get inventory items
// from the mongodb database
func (app *Config) getInventory(c *fiber.Ctx) error {

	// Get the query parameters
	// from the request
	name := c.Query("name")
	category := c.Query("category")
	subCategory := c.Query("sub_category")
	supplier := c.Query("supplier")

	var inventory []models.Item
	var err error

	// Check if none of the query parameters are provided
	// then get all the inventory items
	if name == "" && category == "" && subCategory == "" && supplier == "" {
		inventory, err = getInventoryFromMongoDB(app.Inventory)
	} else {
		inventory, err = getInventoryFromMongoDB(app.Inventory, bson.D{
			// $or operator is used to match any of the provided query parameters
			{Key: "$or", Value: bson.A{
				bson.D{{Key: "name", Value: name}},
				bson.D{{Key: "category", Value: category}},
				bson.D{{Key: "sub_category", Value: subCategory}},
				bson.D{{Key: "supplier", Value: supplier}},
			}},
		})
	}

	if err != nil {
		return utils.SendResponse(
			"failed to get inventory",
			nil,
			c,
			fiber.StatusInternalServerError,
		)
	}

	// Send the response to the client
	return utils.SendResponse("", map[string]any{
		"inventory": inventory,
	}, c, fiber.StatusOK)
}

// getInventoryFromMongoDB is a helper function
// used to get inventory items from the mongodb database
func getInventoryFromMongoDB(db *mongo.Collection, filter ...any) ([]models.Item, error) {

	filterLocal := bson.D{}

	// Check if filter is provided
	// if not then set it to empty bson.D
	// so that it can get all the inventory items
	if len(filter) != 0 {
		filterLocal = filter[0].(bson.D)
	}

	cursor, err := db.Find(context.Background(), filterLocal)

	if err != nil {
		return nil, err
	}

	var inventory []models.Item

	// collect all the inventory items
	for cursor.Next(context.Background()) {
		var item models.Item
		cursor.Decode(&item)
		inventory = append(inventory, item)
	}

	return inventory, nil
}
