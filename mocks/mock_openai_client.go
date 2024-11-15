package mocks

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/mock"
)

// Mock untuk OpenAIClientInterface
type MockOpenAIClient struct {
	mock.Mock
}

func (m *MockOpenAIClient) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*openai.ChatCompletionResponse), args.Error(1)
}
