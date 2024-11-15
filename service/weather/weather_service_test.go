package weather_test

import (
	"bytes"
	"energia/entities"
	"energia/mocks"
	"energia/service/weather"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock HTTP client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	if resp, ok := args.Get(0).(*http.Response); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestWeatherService_GetWeatherByCityAndDate_Success(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockHTTPClient := new(MockHTTPClient)
	dummyApiKey := "dummy_api_key"
	weatherService := weather.NewWeatherServiceWithClient(mockWeatherRepo, dummyApiKey, mockHTTPClient)

	city := "Jakarta"
	date := time.Now()
	existingWeather := entities.Weather{
		City:        city,
		Date:        date,
		Temperature: 28.5,
		Humidity:    80,
		Description: "light rain",
	}

	mockWeatherRepo.On("FindAll", city).Return([]entities.Weather{existingWeather}, nil)

	weatherData, err := weatherService.GetWeatherByCityAndDate(city, date)

	assert.NoError(t, err)
	assert.Equal(t, existingWeather, weatherData)
	mockWeatherRepo.AssertExpectations(t)
}

func TestWeatherService_GetWeatherByCityAndDate_FetchIfNotFound(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockHTTPClient := new(MockHTTPClient)
	dummyApiKey := "dummy_api_key"
	weatherService := weather.NewWeatherServiceWithClient(mockWeatherRepo, dummyApiKey, mockHTTPClient)

	city := "Bandung"
	date := time.Now()
	fetchedWeather := entities.Weather{
		City:        city,
		Date:        date,
		Temperature: 25.0,
		Humidity:    85,
		Description: "cloudy",
	}

	mockWeatherRepo.On("FindAll", city).Return([]entities.Weather{}, nil)

	apiResponse := `{"main":{"temp":25.0,"humidity":85},"weather":[{"description":"cloudy"}]}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(apiResponse)),
	}
	mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(resp, nil)
	mockWeatherRepo.On("Create", mock.AnythingOfType("entities.Weather")).Return(fetchedWeather, nil)

	weatherData, err := weatherService.GetWeatherByCityAndDate(city, date)

	assert.NoError(t, err)
	assert.Equal(t, fetchedWeather, weatherData)
	mockHTTPClient.AssertExpectations(t)
	mockWeatherRepo.AssertExpectations(t)
}

func TestWeatherService_FetchAndStoreWeather_HTTPError(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockHTTPClient := new(MockHTTPClient)
	dummyApiKey := "dummy_api_key"
	weatherService := weather.NewWeatherServiceWithClient(mockWeatherRepo, dummyApiKey, mockHTTPClient)

	city := "Surabaya"
	mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(nil, errors.New("network error"))

	weatherData, err := weatherService.FetchAndStoreWeather(city)

	assert.Error(t, err)
	assert.Equal(t, "network error", err.Error())
	assert.Equal(t, entities.Weather{}, weatherData)
	mockHTTPClient.AssertExpectations(t)
}

func TestWeatherService_FetchAndStoreWeather_StatusNotOK(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockHTTPClient := new(MockHTTPClient)
	dummyApiKey := "dummy_api_key"
	weatherService := weather.NewWeatherServiceWithClient(mockWeatherRepo, dummyApiKey, mockHTTPClient)

	city := "Bali"
	resp := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
	}
	mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(resp, nil)

	weatherData, err := weatherService.FetchAndStoreWeather(city)

	assert.Error(t, err)
	assert.Equal(t, "gagal mengambil data dari OpenWeather API", err.Error())
	assert.Equal(t, entities.Weather{}, weatherData)
	mockHTTPClient.AssertExpectations(t)
}

func TestWeatherService_FetchAndStoreWeather_InvalidJSONResponse(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockHTTPClient := new(MockHTTPClient)
	dummyApiKey := "dummy_api_key"
	weatherService := weather.NewWeatherServiceWithClient(mockWeatherRepo, dummyApiKey, mockHTTPClient)

	city := "Malang"
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`invalid json`)),
	}
	mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(resp, nil)

	weatherData, err := weatherService.FetchAndStoreWeather(city)

	assert.Error(t, err)
	assert.Equal(t, entities.Weather{}, weatherData)
	mockHTTPClient.AssertExpectations(t)
}

func TestWeatherService_FetchAndStoreWeather_CreateError(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockHTTPClient := new(MockHTTPClient)
	dummyApiKey := "dummy_api_key"
	weatherService := weather.NewWeatherServiceWithClient(mockWeatherRepo, dummyApiKey, mockHTTPClient)

	city := "Yogyakarta"
	apiResponse := `{"main":{"temp":27.0,"humidity":70},"weather":[{"description":"sunny"}]}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(apiResponse)),
	}
	mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(resp, nil)
	mockWeatherRepo.On("Create", mock.AnythingOfType("entities.Weather")).Return(entities.Weather{}, errors.New("database error"))

	weatherData, err := weatherService.FetchAndStoreWeather(city)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.Equal(t, entities.Weather{}, weatherData)
	mockHTTPClient.AssertExpectations(t)
	mockWeatherRepo.AssertExpectations(t)
}
