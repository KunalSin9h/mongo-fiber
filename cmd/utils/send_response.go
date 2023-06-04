package utils

import (
	"github.com/gofiber/fiber/v2"
)

// SendResponse is a helper function to send a response
// to the client in a consistent manner
func SendResponse(err string, data map[string]any, c *fiber.Ctx, code int) error {
	var sendResponse struct {
		Error string         `json:"error,omitempty"`
		Data  map[string]any `json:"data,omitempty"`
	}
	sendResponse.Error = err
	sendResponse.Data = data

	return c.Status(code).JSON(sendResponse)
}
