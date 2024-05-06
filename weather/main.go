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
			Name                       string      `json:"name"`
			Temperature                interface{} `json:"temperature"`
			ForecastDetailed           interface{} `json:"detailedForecast"`
			ProbabilityOfPrecipitation struct {
				Value interface{} `json:"value"`
			} `json:"probabilityOfPrecipitation"`
		} `json:"periods"`
	} `json:"properties"`
}

func main() {

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var forecastData Forecast
	if err := json.Unmarshal(body, &forecastData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

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

	forecastBody, err := ioutil.ReadAll(forecastRes.Body)
	if err != nil {
		fmt.Println("Error reading forecast response body:", err)
		return
	}

	var detailedForecast DetailedForecast
	if err := json.Unmarshal(forecastBody, &detailedForecast); err != nil {
		fmt.Println("Error decoding forecast JSON:", err)
		return
	}

	for _, period := range detailedForecast.Properties.Periods {
		var temperature string
		switch t := period.Temperature.(type) {
		case string:
			temperature = t
		case float64:
			temperature = fmt.Sprintf("%.1f", t)
		}

		var chanceOfPrecipitation string
		if period.ProbabilityOfPrecipitation.Value != nil {
			switch p := period.ProbabilityOfPrecipitation.Value.(type) {
			case float64:
				chanceOfPrecipitation = fmt.Sprintf("%.0f%%", p)
			default:
				chanceOfPrecipitation = "N/A"
			}
		} else {
			chanceOfPrecipitation = "N/A"
		}

		fmt.Printf("\nPeriod: %s\nTemperature: %s °F\nChance of precipitation: %s\nDetailed Forecast: %s\n\n", period.Name, temperature, chanceOfPrecipitation, period.ForecastDetailed)
	}
}
