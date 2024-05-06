package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Forecast struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type DetailedForecast struct {
	Properties struct {
		Periods []struct {
			Name        string      `json:"name"`
			Temperature interface{} `json:"temperature"`
		} `json:"periods"`
	} `json:"properties"`
}

func main() {
	// Fetch the forecast URL
	url := "https://api.weather.gov/points/37.6763,-77.4186"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Justins-Weather-App", "github.com/JustinShute")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal the JSON response to get the forecast URL
	var forecastData Forecast
	if err := json.Unmarshal(body, &forecastData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Fetch the detailed forecast using the forecast URL
	forecastURL := forecastData.Properties.Forecast
	forecastReq, err := http.NewRequest("GET", forecastURL, nil)
	if err != nil {
		fmt.Println("Error creating forecast request:", err)
		return
	}

	forecastRes, err := http.DefaultClient.Do(forecastReq)
	if err != nil {
		fmt.Println("Error sending forecast request:", err)
		return
	}
	defer forecastRes.Body.Close()

	// Read the detailed forecast response body
	forecastBody, err := ioutil.ReadAll(forecastRes.Body)
	if err != nil {
		fmt.Println("Error reading forecast response body:", err)
		return
	}

	// Unmarshal the JSON response to get detailed forecast data
	var detailedForecast DetailedForecast
	if err := json.Unmarshal(forecastBody, &detailedForecast); err != nil {
		fmt.Println("Error decoding forecast JSON:", err)
		return
	}

	// Print the high and low temperature data
	for _, period := range detailedForecast.Properties.Periods {
		switch t := period.Temperature.(type) {
		case string:
			fmt.Printf("Period: %s, Temperature: %s\n", period.Name, t)
		case float64:
			fmt.Printf("Period: %s, Temperature: %.1f\n", period.Name, t)
		}
	}
}
