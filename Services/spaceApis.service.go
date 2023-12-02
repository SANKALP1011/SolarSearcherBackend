package Services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"solarseacher.com/solareracherbackend/Error"
	"solarseacher.com/solareracherbackend/Utils"
)

func GetIssLiveLocationService() (*http.Response, error) {
	resp, err := http.Get(Utils.ISSLiveLocationURL)
	if err != nil {
		return nil, Error.CustomErrorHandler("Issue while fetching the location", http.StatusInternalServerError)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Error.CustomErrorHandler("Response status returned is not okay", resp.StatusCode)
	}

	return resp, nil
}

func GetMarsRoverData() (*http.Response, error) {
	currentDate := time.Now()
	previousDate := currentDate.Add(-48 * time.Hour)
	// CHECK FOR THE DATE ISSUES , CURRENTLY GETTING NIL AS THE RESPONSE FROM THE MARS ROVER API.
	url := fmt.Sprintf(Utils.MarsRoverApi, previousDate.Format("2006-01-02"))
	resp, err := http.Get(url)
	log.Println(url)
	if err != nil {
		return nil, Error.CustomErrorHandler("Issue while getting the image from the mars rover api", http.StatusInternalServerError)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, Error.CustomErrorHandler("Response status code returned is not okay", resp.StatusCode)
	}
	log.Println(resp)
	return resp, nil
}
