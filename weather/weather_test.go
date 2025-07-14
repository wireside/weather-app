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

func TestGetWeatherWrongFormat(t *testing.T) {
	format := 256
	geoData := geo.GeoData{
		City: "London",
	}
	
	_, err := weather.GetWeather(geoData, format)
	
	if !errors.Is(err, weather.WrongFormatError) {
		t.Errorf("expected %v, got %v", weather.WrongFormatError, err)
	}
}
