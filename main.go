package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// WeatherAPI JSON Structure
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	// Getting city name as first argument from the command-line
	cityName := os.Args[0]

	// Getting the API key from env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")

	API_PATH := "http://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=" + cityName + "&days=1&aqi=no&alerts=no"

	res, err := http.Get(API_PATH)
	if err != nil {
		panic(err)
	}

	// Returned when func main returns
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}
