package weather_test

import (
	"errors"
	"strings"
	"testing"
	
	"demo/weather/geo"
	"demo/weather/weather"
)

func TestGetWeather(t *testing.T) {
	format := 4
	geoData := geo.GeoData{
		City: "London",
	}
	
	got, err := weather.GetWeather(geoData, format)
	
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
	if !strings.Contains(got, geoData.City) || !strings.Contains(got, "°C") {
		t.Errorf("result does not contain city name or °C")
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 256},
	{name: "0 format", format: 0},
	{name: "Negative format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				geoData := geo.GeoData{
					City: "London",
				}
				
				_, err := weather.GetWeather(geoData, tc.format)
				
				if !errors.Is(err, weather.WrongFormatError) {
					t.Errorf("expected %v, got %v", weather.WrongFormatError, err)
				}
			},
		)
	}
}
