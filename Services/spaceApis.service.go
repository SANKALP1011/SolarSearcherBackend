package Services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"solarseacher.com/solareracherbackend/Error"
	"solarseacher.com/solareracherbackend/Model"
	"solarseacher.com/solareracherbackend/Utils"
)

func GetIssLiveLocationService() (Model.IssLocationModel, error) {

	resp, err := http.Get(Utils.ISSLiveLocationURL)
	if err != nil {
		return Model.IssLocationModel{}, Error.CustomErrorHandler("Issue while fetching the location", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Model.IssLocationModel{}, Error.CustomErrorHandler("Response status returned is not okay", resp.StatusCode)
	}

	var location Model.IssLocationModel
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return Model.IssLocationModel{}, Error.CustomErrorHandler("Issue while decoding the json response", resp.StatusCode)
	}

	return location, nil
}

func GetMarsRoverData() (Model.MarsRoverModel, error) {
	currentDate := time.Now().Format("2006-01-02")
	url := fmt.Sprintf(Utils.MarsRoverApi, currentDate)
	resp, err := http.Get(url)
	if err != nil {
		return Model.MarsRoverModel{}, Error.CustomErrorHandler("Issue while fetching the image from the mars rover api", resp.StatusCode)
	}
	defer resp.Body.Close()
	var roverData Model.MarsRoverModel
	err = json.NewDecoder(resp.Body).Decode(&roverData)
	if err != nil {
		return Model.MarsRoverModel{}, Error.CustomErrorHandler("Issue while coverting the response to json", 400)
	}
	return roverData, nil
}
