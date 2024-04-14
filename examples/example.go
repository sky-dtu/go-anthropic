package examples

import (
	"context"
	"log"

	"github.com/sky-dtu/go-anthropic"
)

// Example demonstrates how to use the anthropic package to create a new client and send a chat completion request.
func Example() {

	client, err := anthropic.NewClient("your-api-key", "claude-3-opus-20240229")
	if err != nil {
		panic(err)
	}

	request := anthropic.ChatCompletionRequest{
		Messages: []anthropic.ChatCompletionMessage{
			{
				Role:    anthropic.RoleUser,
				Content: "Hello!",
			},
			{
				Role:    anthropic.RoleAssistant,
				Content: "Hi there! How can I help you today?",
			},
		},
		Model:     "claude-3-opus-20240229",
		MaxTokens: 100,
	}

	response, err := client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		panic(err)
	}

	log.Println(response)
}
