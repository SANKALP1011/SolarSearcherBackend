package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"solarseacher.com/solareracherbackend/Routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	Routes.SpaceRoutes(app)
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	} else {
		fmt.Printf("Server is running on port %s\n", port)
	}
}
