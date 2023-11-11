package Services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solarseacher.com/solareracherbackend/Model"
	"solarseacher.com/solareracherbackend/Utils"
)

// GetIssLiveLocationService fetches the live location of the ISS

func GetIssLiveLocationService() (Model.IssLocationModel, error) {
	resp, err := http.Get(Utils.ISSLiveLocationURL)
	if err != nil {
		return Model.IssLocationModel{}, fmt.Errorf("failed to make request to ISS live location API: %v", err)
	}
	// This would wait for the entire function to execute first and once the execution is done then only the below line would be executed
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Model.IssLocationModel{}, fmt.Errorf("failed to get ISS live location, status code: %v", resp.StatusCode)
	}

	var location Model.IssLocationModel
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return Model.IssLocationModel{}, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return location, nil
}
