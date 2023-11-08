package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"solarseacher.com/solareracherbackend/Routes"
)

func main() {
	app := fiber.New()

	port := 3000
	Routes.SpaceRoutes(app)
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	} else {
		fmt.Printf("Server is running on port %d\n", port)
	}
}
