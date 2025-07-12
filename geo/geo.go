package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := CheckCity(city)
		if !isCity {
			return nil, errors.New(fmt.Sprintf("city %s doesn't exist", city))
		}

		return &GeoData{
			City: city,
		}, nil
	}

	res, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

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
			"city": strings.ToLower(city),
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
		if res.StatusCode == 404 {
			fmt.Println(errors.New("request failed: city data not found"))
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

	var cityPopulationResponse CityPopulationResponse
	err = json.Unmarshal(body, &cityPopulationResponse)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return !cityPopulationResponse.Error
}
