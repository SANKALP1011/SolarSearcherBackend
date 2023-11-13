package Model

type IssLocationModel struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MarsRoverModel struct {
	Mars_sol     int
	Camera_name  string
	Mars_img     string
	Sol_to_Earth string
}
