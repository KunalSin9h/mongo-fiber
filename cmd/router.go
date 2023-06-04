package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Routes define all the API routes
func (app *Config) routes() *fiber.App {

	// Fiber is the Golang Web Framework
	// Just like Nodejs has Express.js
	router := fiber.New()
	router.Use(cors.New()) // Enable cors

	// API Routes

	// Auth Routes
	router.Post("/api/login", app.login)                                     // Login
	router.Post("/api/register", app.AuthenticationMiddleware, app.register) // Register

	// Inventory Routes
	router.Post("/api/inventory", app.AuthenticationMiddleware, app.addInventory) // Add Inventory Item
	router.Get("/api/inventory", app.AuthenticationMiddleware, app.getInventory)  // Get Inventory Item based on different parameters

	// Sales Routes
	router.Post("/api/sales", app.AuthenticationMiddleware, app.addSales) // Add Sales
	router.Get("/api/sales", app.AuthenticationMiddleware, app.getSales)  // Get Sales

	return router
}
