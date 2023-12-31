package Controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"solarseacher.com/solareracherbackend/Helper"
	"solarseacher.com/solareracherbackend/Model"
	"solarseacher.com/solareracherbackend/Services"
)

func GetISSLocation(c *fiber.Ctx) error {
	data, err := Services.GetIssLiveLocationService()
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	/*

		  Reading the response before the data.body() is closed . If not read m, we would get this error

		  {
		    "Error": "Error decoding JSON: http: read on closed response body"
		}

	*/
	dataBytes, err := ioutil.ReadAll(data.Body)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	defer data.Body.Close()

	if data.StatusCode != http.StatusOK {
		response := fiber.Map{
			"Error": fmt.Sprintf("Unexpected response status: %d", data.StatusCode),
		}
		return c.JSON(response)
	}

	if len(dataBytes) == 0 {
		response := fiber.Map{
			"Error": "Empty response body",
		}
		return c.JSON(response)
	}
	// FIX THE ISSUE OF RESPONSE HERE , CURRENTLY RETURNED NIL.
	var location Model.IssLocationModel
	if err := json.NewDecoder(bytes.NewBuffer(dataBytes)).Decode(&location); err != nil {
		response := fiber.Map{
			"Error": fmt.Sprintf("Error decoding JSON: %s", err.Error()),
		}
		return c.JSON(response)
	}

	response := fiber.Map{
		"Message": location,
		"code":    "200",
	}

	return c.JSON(response)

}
func GetMarsRoverImage(c *fiber.Ctx) error {
	data, err := Services.GetMarsRoverData()
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}

	databytes, err := ioutil.ReadAll(data.Body)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	var marsRoverData Model.MarsRoverModel
	err = json.Unmarshal(databytes, &marsRoverData)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}

	response := fiber.Map{
		"Data": marsRoverData,
		"code": "200",
	}

	defer data.Body.Close()
	return c.JSON(response)
}

func GetApodImages(c *fiber.Ctx) error {
	data, err := Services.GetAstronomyPicOfTheDay()
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	databytes, err := ioutil.ReadAll(data.Body)
	if err != nil {
		response := fiber.Map{
			"Error": err,
		}
		return c.JSON(response)
	}
	var apodData Model.ApodModel
	err = json.Unmarshal(databytes, &apodData)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	response := fiber.Map{

		"Data":    apodData,
		"success": "200",
	}
	defer data.Body.Close()
	return c.JSON(response)
}

func PerformWeatherAnalysis(c *fiber.Ctx) error {
	city := c.Query("city")
	data, err := Services.GetWeatherData(city)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	databytes, err := ioutil.ReadAll(data.Body)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	var weatherModel Model.WeatherModel
	err = json.Unmarshal(databytes, &weatherModel)
	if err != nil {
		response := fiber.Map{
			"Error": err.Error(),
		}
		return c.JSON(response)
	}
	if weatherModel.Curr_Visibilty > 2500 && weatherModel.Main.Curr_temp > 20 && weatherModel.Main.Curr_Humidity < 70 {

	}
	analysis := Helper.WeatherAnalysisHelper(weatherModel.Main.Curr_temp, weatherModel.Curr_Visibilty, weatherModel.Main.Curr_Pressure, weatherModel.Main.Curr_Humidity)
	defer data.Body.Close()
	response := fiber.Map{
		"Data":     weatherModel,
		"Analysis": analysis,
		"Code":     200,
	}
	return c.JSON(response)
}
