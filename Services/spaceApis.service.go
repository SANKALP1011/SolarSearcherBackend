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
	currentDate := time.Now()
	previousDate := currentDate.Add(-48 * time.Hour)

	url := fmt.Sprintf(Utils.MarsRoverApi, previousDate.Format("2006-01-02"))
	resp, err := http.Get(url)

	if err != nil {
		return nil, Error.CustomErrorHandler("Issue while getting the image from the mars rover api", http.StatusInternalServerError)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, Error.CustomErrorHandler("Response status code returned is not okay", resp.StatusCode)
	}
	return resp, nil
}
