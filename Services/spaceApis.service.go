package Services

import (
	"encoding/json"

	"net/http"
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

// func GetMarsRoverData() (Model.CustomError, error) {
// 	resp, err := http.Get(Utils.MarsRoverApi)

// }
