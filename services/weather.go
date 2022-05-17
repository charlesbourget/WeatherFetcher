package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charlesbourget/WeatherFetcher/config"
	"github.com/charlesbourget/WeatherFetcher/types"
)

func GetCurrent(latitude string, longitude string) (types.Current, error) {
	configuration := config.GetConfig()

	data, err := fetchCurrentData(latitude, longitude, configuration.GetString("openweathermap.apiKey"))
	if err != nil {
		return types.Current{}, err
	}

	return data, nil
}

func GetForecast(latitude string, longitude string, count string) (types.Forecast, error) {
	configuration := config.GetConfig()

	data, err := fetchForecastData(latitude, longitude, configuration.GetString("openweathermap.apiKey"), count)
	if err != nil {
		return types.Forecast{}, err
	}

	return data, nil
}

func GetWeather(latitude string, longitude string, count string) (types.Weather, error) {
	configuration := config.GetConfig()

	response := types.Weather{}
	current, err := fetchCurrentData(latitude, longitude, configuration.GetString("openweathermap.apiKey"))
	if err != nil {
		return types.Weather{}, err
	}
	forecast, err := fetchForecastData(latitude, longitude, configuration.GetString("openweathermap.apiKey"), count)
	if err != nil {
		return types.Weather{}, err
	}

	response.Current = current
	response.Forecast = forecast

	return response, nil
}

func fetchCurrentData(latitude string, longitude string, apiKey string) (types.Current, error) {
	uri := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=%s", latitude, longitude, apiKey, "metric")
	resp, err := http.Get(uri)
	if err != nil {
		return types.Current{}, err
	}
	defer resp.Body.Close()

	response := types.Current{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return types.Current{}, err
	}

	return response, nil
}

func fetchForecastData(latitude string, longitude string, apiKey string, count string) (types.Forecast, error) {
	uri := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&appid=%s&units=%s&cnt=%s", latitude, longitude, apiKey, "metric", count)
	resp, err := http.Get(uri)
	if err != nil {
		return types.Forecast{}, err
	}
	defer resp.Body.Close()

	response := types.Forecast{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return types.Forecast{}, err
	}

	return response, nil
}
