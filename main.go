package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	
	flag.Parse()
	
	fmt.Println(*city)
	fmt.Println(*format)
}
