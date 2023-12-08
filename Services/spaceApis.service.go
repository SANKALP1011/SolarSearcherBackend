package Services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"solarseacher.com/solareracherbackend/Error"
	"solarseacher.com/solareracherbackend/Helper"
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

func GetAstronomyPicOfTheDay() (*http.Response, error) {
	resp, err := http.Get(Utils.ApodApi)
	if err != nil {
		return nil, Error.CustomErrorHandler("Issue while fetching the date from apod api", http.StatusInternalServerError)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, Error.CustomErrorHandler("Response status code returned by the api is not okay", resp.StatusCode)
	}

	return resp, nil
}

func GetWeatherData(city string) (*http.Response, error) {
	fmt.Print(os.Getenv("OPEN_WEATHER_API_KEY"))
	resp := Helper.GetGeoCoordinatesData(city)
	lat := resp[0].Geometry.Latitude
	long := resp[0].Geometry.Longitude
	url := fmt.Sprintf(Utils.WeatherApi, lat, long, os.Getenv("OPEN_WEATHER_API_KEY"))
	fmt.Print(url)
	response, err := http.Get(url)
	if err != nil {
		return nil, Error.CustomErrorHandler("unable to fetch the weather data from the api", http.StatusInternalServerError)
	}
	fmt.Print(response.StatusCode)
	if response.StatusCode != http.StatusOK {
		return nil, Error.CustomErrorHandler("Response status code returned by the api is not okay", response.StatusCode)
	}
	return response, nil
}
