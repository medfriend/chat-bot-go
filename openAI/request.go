package openAI

import (
	"chat-bot-go/config"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func MakeRequest(message string) {
	client := resty.New()

	var apiResponse ChatCompletionResponse
	key := config.GetInstance()
	auth := fmt.Sprintf("Bearer %s", key)

	request := ChatRequest{
		Model:     "gpt-3.5-turbo",
		Store:     true,
		MaxTokens: 150,
		Messages: []MessageRequest{
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	// Configura los headers y el cuerpo de la solicitud
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", auth).
		SetBody(request).
		SetResult(&apiResponse).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		fmt.Printf("Error haciendo la solicitud: %v\n", err)
		return
	}

	for _, choice := range apiResponse.Choices {
		fmt.Printf("Mensaje: %s\n", choice.Message.Content)
	}
}
