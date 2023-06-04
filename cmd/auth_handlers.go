package main

import (
	"context"
	"inventory/cmd/utils"
	"inventory/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login API handler is used to login use with email and password
func (app *Config) login(c *fiber.Ctx) error {

	// Define request payload
	// which the client will send
	var reqPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse the request body
	if err := c.BodyParser(&reqPayload); err != nil {
		return utils.SendResponse(
			"Invalid request body",
			nil,
			c,
			fiber.StatusBadRequest,
		)
	}

	email := reqPayload.Email
	password := reqPayload.Password

	user := models.User{}

	// Find user by email from the mongodb database
	err := app.Users.FindOne(context.Background(), bson.D{
		{Key: "email", Value: email},
	}).Decode(&user)

	if err != nil {
		return utils.SendResponse("user does not exist", nil, c, fiber.StatusBadRequest)
	}

	// Check if the password matches
	passwordMatch, err := utils.PasswordMatch(password, user.Password)

	// If there is an error while checking the password
	if err != nil {
		return utils.SendResponse("error while checking password", nil, c, fiber.StatusInternalServerError)
	}

	// If the password does not match
	if !passwordMatch {
		return utils.SendResponse("invalid password", nil, c, fiber.StatusBadRequest)
	}

	// Create a JWT token which will be used for authentication
	// this will be our access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userEmail": user.Email,
		"role":      user.Role,
	})

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return utils.SendResponse("failed to create access token", nil, c, fiber.StatusInternalServerError)
	}

	// Send the token to the client
	return utils.SendResponse("", map[string]any{
		"access_token": tokenString,
	}, c, fiber.StatusOK)
}

// Register API handler is used to register a new user
// Only owner can create new users
func (app *Config) register(c *fiber.Ctx) error {
	role := c.Locals("role")

	// Check if the user is owner
	if role != "owner" {
		return utils.SendResponse("only owner can create new users", nil, c, fiber.StatusUnauthorized)
	}

	// Define request payload
	// which the client will send
	var reqPayload models.User

	// Parse the request body
	if err := c.BodyParser(&reqPayload); err != nil {
		return utils.SendResponse("failed to parser request body", nil, c, fiber.StatusBadRequest)
	}

	// hashing password before storing in database
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(reqPayload.Password), bcrypt.DefaultCost)

	reqPayload.Password = string(hashPassword)

	// create new user on the mongo database
	res, err := app.Users.InsertOne(context.Background(), reqPayload)

	if err != nil {
		return utils.SendResponse("failed to create new user", nil, c, fiber.StatusInternalServerError)
	}

	return utils.SendResponse("", map[string]any{
		"user_id": res.InsertedID,
	}, c, fiber.StatusOK)
}
