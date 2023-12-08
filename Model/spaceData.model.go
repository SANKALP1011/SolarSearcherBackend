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
	HD_image      string `json:"hdurl"`
}
