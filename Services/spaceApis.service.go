package Services

import (
	"fmt"
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
	currentDate := time.Now().Format("2006-1-2")
	url := fmt.Sprintf(Utils.MarsRoverApi, currentDate)
	resp, err := http.Get(url)
	if err != nil {
		return nil, Error.CustomErrorHandler("Issue while getting the image from the mars rover api", http.StatusInternalServerError)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, Error.CustomErrorHandler("Response status code returned is not okay", resp.StatusCode)
	}
	return resp, nil
}
