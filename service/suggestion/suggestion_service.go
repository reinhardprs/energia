package suggestion

import (
	"context"
	"fmt"
	"time"

	"energia/entities"
	"energia/repository/device"
	"energia/repository/weather"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClientInterface interface {
	CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionResponse, error)
}

type SuggestionService struct {
	deviceRepo   device.DeviceRepoInterface
	weatherRepo  weather.WeatherRepoInterface
	openaiClient OpenAIClientInterface
}

func NewSuggestionService(dr device.DeviceRepoInterface, wr weather.WeatherRepoInterface, client OpenAIClientInterface) *SuggestionService {
	return &SuggestionService{
		deviceRepo:   dr,
		weatherRepo:  wr,
		openaiClient: client,
	}
}

func (s *SuggestionService) GetSuggestion(ctx context.Context, userID int, city string) (entities.Suggestion, error) {
	today := time.Now()
	weatherData, err := s.weatherRepo.FindAll(city)
	if err != nil || len(weatherData) == 0 {
		return entities.Suggestion{}, fmt.Errorf("tidak ada data cuaca yang ditemukan untuk kota %s", city)
	}

	var todayWeather entities.Weather
	for _, weather := range weatherData {
		if weather.Date.Format("2006-01-02") == today.Format("2006-01-02") {
			todayWeather = weather
			break
		}
	}

	if (todayWeather == entities.Weather{}) {
		return entities.Suggestion{}, fmt.Errorf("tidak ditemukan data cuaca untuk tanggal hari ini di %s", city)
	}

	devices, err := s.deviceRepo.FindAll(userID)
	if err != nil {
		return entities.Suggestion{}, fmt.Errorf("gagal mengambil data perangkat: %v", err)
	}

	deviceInfo := ""
	var totalPower float32 = 0
	for _, device := range devices {
		totalPower += device.Power
		deviceInfo += fmt.Sprintf("- %s (%.2f watt)\n", device.Name, device.Power)
	}

	prompt := fmt.Sprintf(
		"Hari ini cuaca di kota %s adalah %s dengan suhu %.1f°C dan kelembaban %.1f%%. \n"+
			"User memiliki perangkat sebagai berikut:\n%s"+
			"Total konsumsi daya perangkat: %.2f watt.\n"+
			"Berdasarkan informasi ini, berikan saran penggunaan perangkat yang efisien untuk hari ini.",
		city, todayWeather.Description, todayWeather.Temperature, todayWeather.Humidity, deviceInfo, totalPower,
	)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are an assistant that provides suggestions for efficient device usage based on current weather conditions.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	resp, err := s.openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	})
	if err != nil {
		return entities.Suggestion{}, fmt.Errorf("gagal mendapatkan respons dari OpenAI: %v", err)
	}

	answer := resp.Choices[0].Message.Content

	return entities.Suggestion{
		Message: answer,
	}, nil
}
