package suggestion_test

import (
	"context"
	"energia/entities"
	"energia/mocks"
	"energia/service/suggestion"
	"testing"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSuggestionService_GetSuggestion_Success(t *testing.T) {
	// Mock dependencies
	mockDeviceRepo := new(mocks.DeviceRepoInterface)
	mockWeatherRepo := new(mocks.WeatherRepoInterface)
	mockOpenAIClient := new(mocks.MockOpenAIClient)

	suggestionService := suggestion.NewSuggestionService(mockDeviceRepo, mockWeatherRepo, mockOpenAIClient)

	// Mock weather data for a city
	mockWeather := entities.Weather{
		Date:        time.Now(),
		City:        "Jakarta",
		Description: "Sunny",
		Temperature: 30.0,
		Humidity:    70.0,
	}

	// Mock device data
	mockDevices := []entities.Device{
		{ID: 1, UserID: 1, Name: "Smart Lamp", Power: 100},
		{ID: 2, UserID: 1, Name: "Air Conditioner", Power: 200},
	}

	// Define mock response for OpenAI API
	mockOpenAIResponse := &openai.ChatCompletionResponse{
		Choices: []openai.ChatCompletionChoice{
			{
				Message: openai.ChatCompletionMessage{
					Content: "For today, with sunny weather and high temperature, it's recommended to limit the use of the air conditioner and use the smart lamp efficiently.",
				},
			},
		},
	}

	// Define mocks for repositories
	mockDeviceRepo.On("FindAll", 1).Return(mockDevices, nil)
	mockWeatherRepo.On("FindAll", "Jakarta").Return([]entities.Weather{mockWeather}, nil)
	mockOpenAIClient.On("CreateChatCompletion", mock.Anything, mock.Anything).Return(mockOpenAIResponse, nil)

	// Call the service method
	result, err := suggestionService.GetSuggestion(context.Background(), 1, "Jakarta")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "For today, with sunny weather and high temperature, it's recommended to limit the use of the air conditioner and use the smart lamp efficiently.", result.Message)
	mockDeviceRepo.AssertExpectations(t)
	mockWeatherRepo.AssertExpectations(t)
	mockOpenAIClient.AssertExpectations(t)
}
