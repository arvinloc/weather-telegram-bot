package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather-tg-bot/internal/config"
	"weather-tg-bot/internal/models"
)

type WeatherClient struct {
	apiKey string
}

func NewWeatherClient(cfg *config.Config) *WeatherClient {
	return &WeatherClient{
		apiKey: cfg.APIKey,
	}
}

func (c *WeatherClient) GetWeatherByCity(cityName string) (*models.WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru", cityName, c.apiKey)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	dataByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("CANNOT READ %w", err)
	}

	var weatherData models.WeatherData

	if err := json.Unmarshal(dataByte, &weatherData); err != nil {
		return nil, fmt.Errorf("CANNOT UNMARSHAL JSON TO STRUCT: %w", err)
	}

	if len(weatherData.Weather) == 0 {
		return nil, fmt.Errorf("no weather description available")
	}
	return &weatherData, nil
}
