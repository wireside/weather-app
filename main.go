package main

import (
	"flag"
	"fmt"

	"demo/weather/geo"
	"demo/weather/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 4, "Формат вывода погоды")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(weatherData)
}
