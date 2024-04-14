package anthropic

import (
	"context"
	"errors"
)

type MessageRole string

const (
	anthropicDefaultModel = "claude-3-opus-20240229"

	RoleUser      MessageRole = "user"
	RoleAssistant MessageRole = "assistant"
)

type Client struct {
	ApiKey string
	Model  string
}

/*
NewClient creates a new Client with the provided API key and model.
If the model is not provided, the default model is used.
*/
func NewClient(apiKey, model string) (*Client, error) {

	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	if model == "" {
		model = anthropicDefaultModel
	}

	return &Client{
		ApiKey: apiKey,
		Model:  model,
	}, nil
}

/*
CreateChatCompletion creates a new chat completion request.
with the provided context and request.
*/
func (c *Client) CreateChatCompletion(ctx context.Context, request ChatCompletionRequest) (*ChatCompletionResponse, error) {
	return postMessages(c, ctx, &request)
}
