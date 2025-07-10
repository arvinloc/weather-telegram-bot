package models

type WeatherProvider interface {
	GetWeatherByCity(city string) (*WeatherData, error)
}

type WeatherData struct {
	Weather []WeatherDescription `json:"weather"`
	Main    WeatherMain          `json:"main"`
}

type WeatherDescription struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type WeatherMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
}
