package geo_test

import (
	"errors"
	"testing"
	
	"demo/weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: string(city),
	}
	
	got, err := geo.GetMyLocation(city)
	
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "WrongCity"
	
	_, err := geo.GetMyLocation(city)
	
	if !errors.Is(err, geo.NonExistedCityError) {
		t.Errorf("expected %v, got %v", geo.NonExistedCityError, err)
	}
}
