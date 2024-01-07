package models

type WeatherData struct {
	Date         string `json:"date"`
	TemperatureC int    `json:"temperatureC"`
	Summary      string `json:"summary"`
	TemperatureF int    `json:"temperatureF"`
}
