package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:location`
	Current struct {
		Temp_c    float64 `json:"temp_c"`
		Condition struct {
			Text string `json:text`
		} `json:"condition"`
	} `json:current`
}

func main() {
	fmt.Println("Ready to work")
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=c1a2724b286a497e846191744241910&q=Moscow&aqi=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("The weather api is not available")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	// fmt.Println(weather)
	location, current :=
		weather.Location,
		weather.Current

	fmt.Printf("%s, %s:, %.0fC, %s\n",
		location.Country,
		location.Name,
		current.Temp_c,
		current.Condition.Text)

}

// вытащить с json по ключам значени?
