package Controller

import (
	"github.com/gofiber/fiber/v2"
	"solarseacher.com/solareracherbackend/Services"
)

func GetISSLocation(c *fiber.Ctx) error {
	data, err := Services.GetIssLiveLocationService()

	if err != nil {
		response := fiber.Map{
			"Message": "Failed to get ISS location",
			"Error":   err.Error(),
			"code":    "500",
		}
		return c.JSON(response)
	}

	response := fiber.Map{
		"Message": data,
		"code":    "200",
	}
	return c.JSON(response)
}
