package main

import (
	"encoding/json"
	"errors"
	"example/kwame/weather/helpers"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := `https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true`

	fmt.Println("Fetching Weather Report ....")
	content, _ := getWeatherReport(url)

	var weatherReport helpers.WeatherResponse

	json.Unmarshal([]byte(content), &weatherReport)
	fmt.Println("Timezone:", weatherReport.Timezone)
	fmt.Println("Longitude:", weatherReport.Longitude)
	fmt.Println("Latitude:", weatherReport.Latitude)
	fmt.Println("Elevation:", weatherReport.Elevation)
	fmt.Println("Temperature:", weatherReport.Current_weather.Temperature)
	// getWeatherReport(url)

}

func getWeatherReport(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", errors.New("Error")
	}
	content, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(content))
	return string(content), nil
}
