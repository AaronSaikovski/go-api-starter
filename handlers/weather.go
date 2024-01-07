package handlers

import (
	"errors"
	"math/rand"
	"time"

	"github.com/AaronSaikovski/go-api-starter/models"
	"github.com/gofiber/fiber/v2"
)

// Get Rnadom data
func geRandDate(NumDays int) time.Time {
	// Get the current date and time
	currentTime := time.Now()

	// Add days to the current date
	return currentTime.AddDate(0, 0, NumDays)
}

// Convert from Celsius to Fahrenheit
func convertToF(CelsiusValue int) int {
	return (CelsiusValue * 9 / 5) + 32
}

// Generate random Weatherforecast data
func generateRandomWeatherData(WeatherData *models.WeatherData) error {

	//check for valid weatherData object
	if WeatherData == nil {
		return errors.New("**Error: WeatherData data is empty or invalid..**")
	}

	// Weather summaries
	summaries := []string{
		"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching",
	}

	//seed the randomiser
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//get a random number between 1 & 5
	randDay := rand.Intn(5) + 1

	// calc a random temp
	randTemp := rand.Intn(55) + -22

	// Generate a random index within the length of the slice
	randomIndex := rand.Intn(len(summaries))

	// Get the random value from the slice using the random index
	randomSummary := summaries[randomIndex]

	// populate weatherdata struct
	WeatherData.Date = geRandDate(randDay)
	WeatherData.TemperatureC = randTemp
	WeatherData.Summary = randomSummary
	WeatherData.TemperatureF = convertToF(randTemp)

	return nil
}

// @Summary Sample weatherforecast
// @Description Sample random weatherforecast data
// @Tags weather
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /weatherforecast [get]
func HandleWeatherGet(c *fiber.Ctx) error {

	// get weather data and update pointer to struct
	WeatherData := new(models.WeatherData)
	err := generateRandomWeatherData(WeatherData)

	if err != nil {
		panic(err)
	}

	// return the weather data json
	return c.Status(200).JSON(WeatherData)

}
