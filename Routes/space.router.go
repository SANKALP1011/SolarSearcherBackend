package Routes

import (
	"github.com/gofiber/fiber/v2"
	"solarseacher.com/solareracherbackend/Controller"
)

func SpaceRoutes(app *fiber.App) {
	app.Get("/v1/getIssLocation", Controller.GetISSLocation)
	app.Get("/v1/getMarsRoverPic", Controller.GetMarsRoverImage)
	app.Get("/v1/getApodApiImage", Controller.GetApodImages)
	app.Get("/v1/performWeatherAnalysis", Controller.PerformWeatherAnalysis)
}
