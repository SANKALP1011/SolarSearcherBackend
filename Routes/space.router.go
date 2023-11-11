package Routes

import (
	"github.com/gofiber/fiber/v2"
	"solarseacher.com/solareracherbackend/Controller"
)

func SpaceRoutes(app *fiber.App) {
	app.Get("/v1/getWeatherAnalysis", Controller.GetWeatherAnalaysis)
	app.Get("/v1/getIssLocation", Controller.GetISSLocation)
}
