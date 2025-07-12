package geo_test

import (
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
	if *got != expected {
		t.Errorf("ожидалось %v, получили %v", expected, got)
	}
	
}
