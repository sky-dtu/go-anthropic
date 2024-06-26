package anthropic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Constants related to the API configuration
const (
	anthropicBaseUrl      = "https://api.anthropic.com/v1" // Base URL of the Anthropic API
	anthropicApiKeyHeader = "x-api-key"                    // HTTP header for API key authentication

	anthropicApiVersionHeader = "anthropic-version" // HTTP header for specifying API version
	antrhopicApiVersion       = "2023-06-01"
)

// ChatCompletionRequest defines the structure for the API request payload
type ChatCompletionRequest struct {
	Messages  []ChatCompletionMessage `json:"messages"`   // List of messages to process
	Model     string                  `json:"model"`      // Model identifier
	MaxTokens int                     `json:"max_tokens"` // Maximum number of tokens to generate
	System    string                  `json:"system"`     // System identifier
}

// ChatCompletionMessage represents a single interaction unit within a request
type ChatCompletionMessage struct {
	Role    MessageRole `json:"role"`    // Role of the message sender
	Content string      `json:"content"` // Text content of the message
}

// ChatCompletionResponse defines the structure for the API response payload
type ChatCompletionResponse struct {
	Id           string                  `json:"id"`            // Unique identifier of the completion
	Content      []ChatCompletionContent `json:"content"`       // List of content generated by the model
	Model        string                  `json:"model"`         // Model used for generating the response
	StopReason   string                  `json:"stop_reason"`   // Reason for stopping the generation
	StopSequence string                  `json:"stop_sequence"` // Sequence that caused the generation to stop
	Usage        Usage                   `json:"usage"`         // Usage information of the tokens
}

// String returns a string representation of the ChatCompletionResponse
func (c *ChatCompletionResponse) String(del string) string {
	if len(c.Content) == 0 {
		return ""
	}

	if del == "" {
		del = "\n\n"
	}

	resp := ""
	for _, content := range c.Content {
		resp += content.Text + del
	}

	// remove the last delimiter
	if len(resp) >= len(del) {
		resp = resp[:len(resp)-len(del)]
	}

	return resp
}

// ChatCompletionContent represents the generated content
type ChatCompletionContent struct {
	Text  string      `json:"text"`  // Generated text
	Id    string      `json:"id"`    // Identifier for this piece of content
	Name  string      `json:"name"`  // Name associated with the content
	Input interface{} `json:"input"` // Original input related to the generated content
}

// Usage provides information about token usage
type Usage struct {
	InputTokens  int `json:"input_tokens"`  // Number of tokens used in the input
	OutputTokens int `json:"output_tokens"` // Number of tokens produced in the output
}

// postMessages sends a POST request to the /messages endpoint of the Anthropic API
// and handles the response to return a parsed CompletionResponse or an error.
func postMessages(client *Client, ctx context.Context, completion *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	// Construct the URL for the request
	url := anthropicBaseUrl + "/messages"

	// Set the default model if not provided
	if completion.Model == "" {
		completion.Model = client.Model
	}

	// Marshal the completion request data into JSON format
	reqBody, err := json.Marshal(completion)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal completion request")
	}

	// Create a new HTTP request with the JSON body
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HTTP request")
	}

	// Set required headers for the request
	req.Header.Add(anthropicApiKeyHeader, client.ApiKey)
	req.Header.Add(anthropicApiVersionHeader, antrhopicApiVersion)
	req.Header.Add("Content-Type", "application/json")

	// Initialize a new HTTP client
	hc := &http.Client{}

	// Execute the HTTP request
	resp, err := hc.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute HTTP request")
	}
	defer resp.Body.Close() // Ensure the response body is closed after reading

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// Handle non-OK status codes
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("API error [%d]: %v", resp.StatusCode, string(body))
	}

	// Unmarshal the JSON response into a CompletionResponse struct
	var completionResponse ChatCompletionResponse
	err = json.Unmarshal(body, &completionResponse)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal completion response")
	}

	return &completionResponse, nil // Return the parsed response
}
