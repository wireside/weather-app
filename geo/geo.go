package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CountriesPopulationResponse struct {
	error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	
	res, err := http.Get("http://ip-api.com/json/")
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

func CheckCity(city string) bool {
	postBody, _ := json.Marshal(
		map[string]string{
			"city": city,
		},
	)
	res, err := http.Post(
		"https://countriesnow.space/api/v0.1/countries/population/cities",
		"application/json",
		bytes.NewBuffer(postBody),
	)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		if res.StatusCode == 429 {
			fmt.Println(errors.New("request failed: too many requests"))
			return false
		}
		fmt.Println(errors.New("request failed: status code is not 200"))
		return false
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	
	var countriesPopulationResponse CountriesPopulationResponse
	err = json.Unmarshal(body, countriesPopulationResponse)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	
	return !countriesPopulationResponse.error
}
