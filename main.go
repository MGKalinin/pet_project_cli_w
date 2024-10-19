package main

import (
	"fmt"
	"io"
	"net/http"
)

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
	fmt.Println(string(body))
}

// вытащить с json по ключам значени?
