package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	
	"demo/weather/geo"
)

func GetWeather(geoData geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	
	res, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		fmt.Println(errors.New("request failed: status code is not 200"))
		return ""
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	
	return string(body)
}
