package services

import (
	"encoding/json"
	"fmt"
	"github.com/charlesbourget/WeatherFetcher/config"
	"github.com/charlesbourget/WeatherFetcher/types"
	"net/http"
)

func GetWeather(latitude string, longitude string) (types.OWMResponse, error) {
	configuration := config.GetConfig()

	data, err := fetchData(latitude, longitude, configuration.GetString("openweathermap.apiKey"))
	if err != nil {
		return types.OWMResponse{}, err
	}

	return data, nil
}

func fetchData(latitude string, longitude string, apiKey string) (types.OWMResponse, error) {
	uri := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=%s", latitude, longitude, apiKey, "metric")
	resp, err := http.Get(uri)
	if err != nil {
		return types.OWMResponse{}, err
	}
	defer resp.Body.Close()

	response := types.OWMResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return types.OWMResponse{}, err
	}

	return response, nil
}
