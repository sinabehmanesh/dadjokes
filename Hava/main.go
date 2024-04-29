package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	Types "main/types"
	"net/http"
	"os"

	"github.com/enescakir/emoji"
	"github.com/joho/godotenv"
)

func current_weather(city string) string {
	godotenv.Load(".env")
	api_url := os.Getenv("api_url")
	api_key := os.Getenv("api_key")

	url := api_url + "current.json?q=" + city

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("KEY", api_key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	rawresponse, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	final_response := string(rawresponse)
	return final_response
}

// Define view Weather
func view_weather(city Types.Weather_type) {
	fmt.Println(city.Location.Name, " - ", city.Location.Localtime)
	fmt.Println(emoji.Thermometer, " Tmpurature is", city.Current.TempC, ", it feels like:", city.Current.FeelslikeC)
	if city.Current.Cloud < 30 {
		fmt.Println(emoji.SunWithFace, "sky resolution is perfect! few clouds can be seen in the area!")
	} else if city.Current.Cloud > 30 && city.Current.Cloud < 60 {
		fmt.Println(emoji.SunBehindCloud, "We have a Cloud sky today, hope it rains i Guess!")
	} else if city.Current.Cloud > 60 {
		fmt.Println(emoji.Cloud, "Clouds are everywhere!")
	}

	//city WIND report
	if city.Current.WindKph < 15 {
		fmt.Println(emoji.LeafFlutteringInWind, " Calm and chill...")
	} else if city.Current.WindKph >= 15 && city.Current.WindKph < 25 {
		fmt.Println(emoji.Fog, " Its quiet windy out there! ", city.Current.WindDir)
	} else if city.Current.WindKph > 25 {
		fmt.Println(emoji.Tornado, " Do not Go out! ", city.Current.WindDir)
	}

	fmt.Println()
}

// Define the main function
func main() {

	//Yes try to hardcode it.
	first_city := "Munich"
	second_city := "Tehran"

	//Get current weather or the first city

	current_first_city := current_weather(first_city)
	current_second_city := current_weather(second_city)

	var fcity_weather Types.Weather_type
	var scity_weather Types.Weather_type

	json.Unmarshal([]byte(current_first_city), &fcity_weather)
	json.Unmarshal([]byte(current_second_city), &scity_weather)

	view_weather(fcity_weather)
	view_weather(scity_weather)
}
