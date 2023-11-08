package Controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetWeatherAnalaysis(c *fiber.Ctx) error {
	data := "This is the end point for the analysis for the weather"
	response := fiber.Map{
		"Message": data,
		"code":    "200",
	}
	return c.JSON(response)
}

func GetSpaceData() {
	fmt.Println("this would perform somke space realted data")
}

func GetISSLocation() {
	fmt.Println("this would get the location of iss using our iss api")
}
