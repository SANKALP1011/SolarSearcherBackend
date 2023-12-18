package Model

type IssLocationModel struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MarsRoverModel struct {
	Photos []struct {
		ID     int `json:"id"`
		Sol    int `json:"sol"`
		Camera struct {
			Name string `json:"name"`
		} `json:"camera"`
		ImgSrc    string `json:"img_src"`
		EarthDate string `json:"earth_date"`
	} `json:"photos"`
}

type ApodModel struct {
	Copyright     string `json:"copyright"`
	Date_of_click string `json:"date"`
	Description   string `json:"explanation"`
	HD_image      string `json:"url"`
}

type WeatherModel struct {
	Cordinates struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	} `json:"coord"`
	Main struct {
		Curr_temp     float64 `json:"temp"`
		Curr_Pressure float64 `json:"pressure"`
		Curr_Humidity float64 `json:"humidity"`
	} `json:"main"`
	Curr_Visibilty int32 `json:"visibility"`
}
