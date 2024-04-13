package anthropic

import (
	"context"
	"errors"
)

const (
	anthropicDefaultModel = "claude-3-opus-20240229"
)

type Client struct {
	ApiKey string
	Model  string
}

func NewClient(client Client) (*Client, error) {

	if client.ApiKey == "" {
		return nil, errors.New("API key is required")
	}

	if client.Model == "" {
		client.Model = anthropicDefaultModel
	}

	return &client, nil
}

func (c *Client) CreateMessage(role, content string) *Message {
	return &Message{
		Role:    role,
		Content: content,
	}
}

func (c *Client) DoCompletionRequest(ctx context.Context, request CompletionRequest) (*CompletionResponse, error) {
	return postMessages(c, &request)
}
