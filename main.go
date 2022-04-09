package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strzybug/cache"
	"strzybug/weather"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

const Addr = ":8080"

func dateFormatter(layout string) func(t time.Time) string {
	return func(t time.Time) string {
		return t.Format(layout)
	}
}

func main() {
	owmApiKey := os.Getenv("OpenWeatherMap_Api_Key")

	weather := cache.Cache{
		Filename: "weather-cache.json",
		Request: weather.Request{
			Latitude:     "52.40",
			Longitude:    "16.93",
			LanguageCode: "pl",
			ApiKey:       owmApiKey,
		},
	}

	if err := weather.Init(); err != nil {
		log.Fatalln(err)
	}

	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := t.ExecuteTemplate(w, "page", weather.Access().Daily); err != nil {
			log.Fatalln("http.HandleFunc(/): ", err)
		}
	})

	fmt.Printf("Listening on http://localhost%s\n", Addr)
	http.ListenAndServe(Addr, nil)
}
