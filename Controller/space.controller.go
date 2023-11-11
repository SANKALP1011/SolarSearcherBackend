package Controller

import (
	"github.com/gofiber/fiber/v2"
	"solarseacher.com/solareracherbackend/Services"
)

func GetWeatherAnalaysis(c *fiber.Ctx) error {
	data := "This is the end point for the analysis for the weather"
	response := fiber.Map{
		"Message": data,
		"code":    "200",
	}
	return c.JSON(response)
}

func GetSpaceData(c *fiber.Ctx) error {
	data := " This is the endpoint to fetch the iss live location"
	response := fiber.Map{
		"Message": data,
		"code":    200,
	}
	return c.JSON(response)
}

func GetISSLocation(c *fiber.Ctx) error {
	data, err := Services.GetIssLiveLocationService()

	if err != nil {
		// Handle the error appropriately, for now, returning a JSON response
		response := fiber.Map{
			"Message": "Failed to get ISS location",
			"Error":   err.Error(),
			"code":    "500",
		}
		return c.JSON(response)
	}

	response := fiber.Map{
		"Message": data, // You might need to extract relevant information from data
		"code":    "200",
	}
	return c.JSON(response)
}
