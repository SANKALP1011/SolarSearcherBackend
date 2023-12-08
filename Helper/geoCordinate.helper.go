package Helper

import (
	"os"

	"github.com/rubenv/opencagedata"
	"solarseacher.com/solareracherbackend/Error"
)

func GetGeoCoordinatesData(city string) []opencagedata.GeocodeResultItem {
	geocoder := opencagedata.NewGeocoder(os.Getenv("GEO_CODING"))
	res, err := geocoder.Geocode(city, nil)
	if err != nil {
		Error.CustomErrorHandler("Unable to geocode the location", 400)
	}

	return res.Results
}
