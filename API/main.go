package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Testing API calls")

	weatherResponse, weatherError := http.Get("https://wttr.in/Stockholm?T")

	if weatherError != nil {
		fmt.Println(weatherError.Error())
		os.Exit(1)
	}

	weatherResponseData, weatherError := io.ReadAll(weatherResponse.Body)
	if weatherError != nil {
		log.Fatal(weatherError)
	}

	fmt.Println(string(weatherResponseData))
}
