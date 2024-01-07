package models

import "time"

//	{
//	  "date": "2024-01-07",
//	  "temperatureC": 0,
//	  "summary": "string",
//	  "temperatureF": 0
//	}
type WeatherData struct {
	Date         time.Time `json:"date"`
	TemperatureC int       `json:"temperatureC"`
	Summary      string    `json:"summary"`
	TemperatureF int       `json:"temperatureF"`
}
