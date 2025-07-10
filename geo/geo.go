package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		if res.StatusCode == 429 {
			return nil, errors.New("request failed: too many requests")
		}
		return nil, errors.New("request failed: status code is not 200")
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	var geoData GeoData
	err = json.Unmarshal(body, &geoData)
	if err != nil {
		return nil, err
	}
	
	return &geoData, nil
}
