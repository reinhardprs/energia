package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"energia/entities"
	"energia/repository/weather"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type WeatherService struct {
	weatherRepoInterface weather.WeatherRepoInterface
	apiKey               string
	client               HTTPClient
}

func NewWeatherService(wr weather.WeatherRepoInterface) *WeatherService {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		panic("API key untuk OpenWeather tidak ditemukan. Set OPENWEATHER_API_KEY di environment.")
	}
	return &WeatherService{
		weatherRepoInterface: wr,
		apiKey:               apiKey,
		client:               &http.Client{},
	}
}

func NewWeatherServiceWithClient(wr weather.WeatherRepoInterface, apiKey string, client HTTPClient) *WeatherService {
	if apiKey == "" {
		panic("API key untuk OpenWeather tidak boleh kosong.")
	}
	return &WeatherService{
		weatherRepoInterface: wr,
		apiKey:               apiKey,
		client:               client,
	}
}

func (w *WeatherService) FetchAndStoreWeather(city string) (entities.Weather, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s&units=metric", city, w.apiKey)

	resp, err := w.client.Get(url)
	if err != nil {
		return entities.Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entities.Weather{}, errors.New("gagal mengambil data dari OpenWeather API")
	}

	var apiResponse struct {
		Main struct {
			Temp     float32 `json:"temp"`
			Humidity float32 `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return entities.Weather{}, err
	}

	weatherData := entities.Weather{
		City:        city,
		Date:        time.Now(),
		Temperature: apiResponse.Main.Temp,
		Humidity:    apiResponse.Main.Humidity,
		Description: apiResponse.Weather[0].Description,
	}

	savedWeather, err := w.weatherRepoInterface.Create(weatherData)
	if err != nil {
		return entities.Weather{}, err
	}

	return savedWeather, nil
}

func (w *WeatherService) GetWeatherByCityAndDate(city string, date time.Time) (entities.Weather, error) {
	weathers, err := w.weatherRepoInterface.FindAll(city)
	if err != nil {
		return entities.Weather{}, err
	}

	for _, weather := range weathers {
		if weather.Date.Format("2006-01-02") == date.Format("2006-01-02") {
			return weather, nil
		}
	}

	weatherData, err := w.FetchAndStoreWeather(city)
	if err != nil {
		return entities.Weather{}, err
	}

	return weatherData, nil
}
