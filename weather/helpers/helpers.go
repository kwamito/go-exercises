package helpers

import "time"

type CurrentWeather struct {
	Temperature   float32
	Windspeed     float32
	Winddirection float32
	Weathercode   int
	time          time.Time
}

type WeatherResponse struct {
	Latitude              float32
	Longitude             float32
	Generationtime_ms     float32
	Utc_offset_seconds    int
	Timezone              string
	Timezone_abbreviation string
	Elevation             float32
	Current_weather       CurrentWeather
}
