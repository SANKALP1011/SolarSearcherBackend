package Controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
