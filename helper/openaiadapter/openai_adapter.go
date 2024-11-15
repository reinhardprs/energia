package openaiadapter

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

// Adaptor untuk membuat *openai.Client memenuhi OpenAIClientInterface
type OpenAIClientAdapter struct {
	client *openai.Client
}

func NewOpenAIClientAdapter(client *openai.Client) *OpenAIClientAdapter {
	return &OpenAIClientAdapter{
		client: client,
	}
}

func (o *OpenAIClientAdapter) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionResponse, error) {

	response, err := o.client.CreateChatCompletion(ctx, req)

	return &response, err

}
