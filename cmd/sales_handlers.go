package main

import (
	"context"
	"inventory/cmd/utils"
	"inventory/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Add Sales Record
func (app *Config) addSales(c *fiber.Ctx) error {
	// Get seller details from access_token
	sellerEmail := c.Locals("userEmail").(string)

	var seller models.User

	// find the seller details from the database
	// for adding to the sales record
	err := app.Users.FindOne(context.Background(), bson.D{
		{Key: "email", Value: sellerEmail},
	}).Decode(&seller)

	if err != nil {
		return utils.SendResponse("failed to get seller details", nil, c, fiber.StatusBadRequest)
	}

	// Parse the request body
	var salesRecord models.SalesRecord

	if err := c.BodyParser(&salesRecord); err != nil {
		return utils.SendResponse("failed to parse request body", nil, c, fiber.StatusBadRequest)
	}

	// Add seller details to the sales record
	salesRecord.SellerName = seller.Name
	salesRecord.SellerEmail = seller.Email
	salesRecord.SellerPhone = seller.PhoneNumber
	salesRecord.Date = time.Now().Format("02-01-2006")
	salesRecord.Time = time.Now().Format("15:04:05")

	var item models.Item

	// Find the item details from the database
	// fot checking if stock is available
	// and deducting the quantity from the stock
	err = app.Inventory.FindOne(context.Background(), bson.D{
		{Key: "name", Value: salesRecord.Product},
	}).Decode(&item)

	// Check if the item exists in the database
	if err != nil {
		return utils.SendResponse("failed to get item details", nil, c, fiber.StatusBadRequest)
	}

	// Check if the stock is available
	if item.Quantity < salesRecord.Quantity {
		return utils.SendResponse("insufficient stock", nil, c, fiber.StatusBadRequest)
	}

	// Deduct the quantity from the stock
	_, err = app.Inventory.UpdateOne(context.Background(), bson.D{
		{Key: "name", Value: salesRecord.Product},
	}, bson.D{
		{Key: "$set", Value: bson.D{{Key: "quantity", Value: item.Quantity - salesRecord.Quantity}}},
	})

	// Check if the update operation was successful
	if err != nil {
		return utils.SendResponse("failed to update item details", nil, c, fiber.StatusBadRequest)
	}

	// Insert the sales record to the database
	res, err := app.Sales.InsertOne(context.Background(), salesRecord)

	if err != nil {
		return utils.SendResponse("failed to insert sales record", nil, c, fiber.StatusInternalServerError)
	}

	return utils.SendResponse("", map[string]any{
		"sales_id": res.InsertedID,
	}, c, fiber.StatusOK)
}

// Get Sales Record
func (app *Config) getSales(c *fiber.Ctx) error {

	// sort the sales records
	var sales []models.SalesRecord

	// Find all the sales records from the database
	cursor, err := app.Sales.Find(context.Background(), bson.D{})

	if err != nil {
		return utils.SendResponse("failed to get sales records", nil, c, fiber.StatusInternalServerError)
	}

	for cursor.Next(context.Background()) {
		var salesRecord models.SalesRecord
		cursor.Decode(&salesRecord)
		sales = append(sales, salesRecord)
	}

	return utils.SendResponse("", map[string]any{
		"sales": sales,
	}, c, fiber.StatusOK)
}
