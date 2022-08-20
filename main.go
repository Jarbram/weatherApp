package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var API_KEY = "a133bc4edf349c737976b2541308aaef"

type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

type WeatherData struct {
	Main Main   `json:"main"`
	Name string `json:"name"`
}

func main() {
	var city string

	fmt.Println("Enter your city name: ")
	fmt.Scanf("%s", &city)
	fmt.Println()

	final_url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, API_KEY)

	response, err := http.Get(final_url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	textBytes := []byte(body)
	weather_data := WeatherData{}
	json.Unmarshal(textBytes, &weather_data)

	fmt.Println("City:", weather_data.Name)
	fmt.Println("Temperature:", weather_data.Main.Temp)
	fmt.Println("Temperature(Min):", weather_data.Main.TempMin)
	fmt.Println("Temperature(Max):", weather_data.Main.TempMax)
	fmt.Println("Pressure:", weather_data.Main.Pressure)
	fmt.Println("Humidity:", weather_data.Main.Humidity)

}
