package mocks

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClientInterface interface {
	CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionResponse, error)
}
