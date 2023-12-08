package Helper

func WeatherAnalysisHelper(temp float64, visibility int32, pressure float64, humidity float64) string {
	if temp > 25 && visibility > 2500 && pressure > 30 && humidity < 80 {
		return "Good weather conditions! You can watch the astronomical stuff."
	}

	return "Weather condition is not so good for carrying out astronomical activities"
}
