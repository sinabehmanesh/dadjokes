package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	Types "main/types"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func current_weather(city string) string {
	fmt.Println("checking weather for: " + city)

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
func view_weather(fcity Types.Weather_type, scity Types.Weather_type) {
	fmt.Println("We have tmpurature of ", fcity.Current.TempC, " in "+fcity.Location.Name)
	fmt.Println("We have tmpurature of ", scity.Current.TempC, " in "+scity.Location.Name)
}

// Define the main function
func main() {

	first_city := "Paris"
	second_city := "Tehran"
	//Get current weather or the first city

	current_first_city := current_weather(first_city)
	current_second_city := current_weather(second_city)

	var fcity_weather Types.Weather_type
	var scity_weather Types.Weather_type

	json.Unmarshal([]byte(current_first_city), &fcity_weather)
	json.Unmarshal([]byte(current_second_city), &scity_weather)

	view_weather(fcity_weather, scity_weather)

}
