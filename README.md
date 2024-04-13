# Anthropic Go

This Go package provides a straightforward and idiomatic way to interact with the Anthropic API. It supports creating and sending message requests to generate responses using the Claude-3 model. This document will guide you through the installation process and basic usage of the package.

## Installation

To use this package, first ensure you have Go installed on your machine. You can download and install Go from [golang.org](https://golang.org/dl/).

Once Go is installed, you can add this package to your project by running:

```bash
go get -u github.com/sky-dtu/anthropic
```

## Configuration

To use the package, you need to provide your Anthropic API key. You can obtain an API key by signing up on the [Anthropic website](https://docs.anthropic.com/claude/docs/getting-access-to-claude#step-3-generate-an-api-key). Once you have an API key, you can set it in your environment variables:


## Usage

### Creating a Client

First, import the package and create a new client instance. If you don't specify a model, the default `claude-3-opus-20240229` model will be used.

```go
package main

import (
    "github.com/sky-dtu/anthropic"
    "log"
)

func main() {
    client, err := anthropic.NewClient(anthropic.Client{
        ApiKey: "your_api_key_here",
    })
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
}
```

### Creating a Message

Create a message to send to the API:

```go
message := client.CreateMessage("user", "Hello, how are you?")
```

### Sending a Completion Request

To send a completion request and receive a response, use the DoCompletionRequest function. Ensure you pass a context to manage request cancellation and timeouts.

```go
import "context"

func main() {
    // Assume client has been initialized as shown above

    // Create a completion request
    request := anthropic.CompletionRequest{
        Messages:  []anthropic.Message{*message},
        Model:     client.Model,
        MaxTokens: 150,
    }

    // Send the completion request
    response, err := client.DoCompletionRequest(context.Background(), request)
    if err != nil {
        log.Fatalf("Failed to send completion request: %v", err)
    }

    log.Printf("Response: %+v", response)
}
```

### Handling Errors

The package utilizes the errors package to provide detailed error messages. Always check for errors after making API calls.

### Full Example

Here is a full example that demonstrates how to create a client, send a completion request, and handle the response:

```go
package main

import (
    "context"
    "github.com/sky-dtu/anthropic"
    "log"
)

func main() {
    // Create a new client
    client, err := anthropic.NewClient(anthropic.Client{
        ApiKey: "your_api_key_here",
    })
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Create a message
    message := client.CreateMessage("user", "Hello, how are you?")
    request := anthropic.CompletionRequest{
        Messages:  []anthropic.Message{*message},
        Model:     client.Model,
        MaxTokens: 150,
    }

    // Send the completion request
    response, err := client.DoCompletionRequest(context.Background(), request)
    if err != nil {
        log.Fatalf("Failed to send completion request: %v", err)
    }

    log.Printf("Response: %+v", response)
}
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests, report issues, or suggest improvements via the GitHub repository.

## License

This package is released under the MIT License. See the [LICENSE](LICENSE) file for more information.
