package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"demo/weather/geo"
)

var (
	Not200Error      = errors.New("request failed: status code is not 200")
	WrongFormatError = errors.New("wrong weather format")
)

func GetWeather(geoData geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", WrongFormatError
	}

	baseUrl, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", Not200Error
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
