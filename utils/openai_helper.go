package utils

import (
	"going-places-api/services"
	"going-places-api/config"
)

func FetchCompletionFromOpenAI(prompt string) (string, error) {
	apiKey := config.OPENAI_API_KEY
	openAIClient := services.NewOpenAIClient(apiKey)

	messages := []services.OpenAIChatMessage{
    	{
    		Role:    "system",
    		Content: prompt,
    	},
    }

	return openAIClient.GetChatCompletion(messages)
}
