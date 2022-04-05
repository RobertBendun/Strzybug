package main

import (
	"fmt"
	"log"
	"os"
	"strzybug/cache"
	"strzybug/weather"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	owmApiKey := os.Getenv("OpenWeatherMap_Api_Key")

	weather, err := cache.New("weather-cache.json", weather.Request{
		Latitude:     "52.40",
		Longitude:    "16.93",
		LanguageCode: "pl",
		ApiKey:       owmApiKey,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(weather.Access())
}
