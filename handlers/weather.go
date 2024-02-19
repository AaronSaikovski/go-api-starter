package handlers

// import (
// 	"math/rand"
// 	"net/http"
// 	"time"

// 	"github.com/AaronSaikovski/go-api-starter/models"
// 	"github.com/labstack/echo/v4"
// 	"github.com/rs/zerolog/log"
// )

// // Convert from Celsius to Fahrenheit
// func convertToF(CelsiusValue int) int {
// 	return (CelsiusValue * 9 / 5) + 32
// }

// // Get a weather summary based on tempC
// func getWeatherSummary(TempC int) string {

// 	switch {
// 	case TempC <= 10:
// 		return "Freezing"
// 	case TempC > 10 && TempC <= 20:
// 		return "Cool"
// 	case TempC > 20 && TempC <= 30:
// 		return "Moderate"
// 	case TempC > 30 && TempC <= 45:
// 		return "Moderate"
// 	case TempC > 45:
// 		return "Extreme"
// 	default:
// 		return "unknown"
// 	}
// }

// // Generate random Weatherforecast data
// func generateRandomWeatherData() []models.WeatherData {

// 	//seed the randomiser
// 	rand.New(rand.NewSource(time.Now().UnixNano()))

// 	//get todays daye
// 	currentTime := time.Now()

// 	//Create a struct of 5 elements
// 	var WeatherData []models.WeatherData

// 	//main loop for 5 days
// 	for i := 0; i < 5; i++ {

// 		//Set the date
// 		CurrentDay := currentTime.AddDate(0, 0, i)
// 		formattedDate := CurrentDay.Format("2006-01-02")

// 		// calc a random temp - TemperatureC
// 		RandomTemp := rand.Intn(55) + -22

// 		// Convert to - TemperatureF
// 		TemperatureF := convertToF(RandomTemp)

// 		//Get summary for temp
// 		WeatherSummary := getWeatherSummary(RandomTemp)

// 		//add items to the struct
// 		NewWeatherItem := models.WeatherData{
// 			Date:         formattedDate,
// 			TemperatureC: RandomTemp,
// 			TemperatureF: TemperatureF,
// 			Summary:      WeatherSummary,
// 		}
// 		WeatherData = append(WeatherData, NewWeatherItem)

// 	}

// 	return WeatherData
// }

// // @Summary Sample weatherforecast
// // @Description Sample random weatherforecast data
// // @Tags sampleweatherdata
// // @Accept */*
// // @Produce plain
// // @Success 200 "OK"
// // @Router /api/weatherforecast [get]
// func HandleWeatherGet(c echo.Context) error {

// 	log.Debug().Msg("Calling HandleWeatherGet()")

// 	// get weather data and return JSON Object
// 	WeatherData := generateRandomWeatherData()

// 	// return the weather data json
// 	return c.JSON(http.StatusOK, WeatherData)

// }
