package services

import (
    "bytes"
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
)

type OpenAIClient struct {
	APIKey string
	client *http.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		APIKey: apiKey,
		client: &http.Client{},
	}
}

type OpenAIChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatRequest struct {
	Model    string             `json:"model"`
	Messages []OpenAIChatMessage `json:"messages"`
	MaxTokens int `json:"max_tokens"`
    Temperature float32 `json:"temperature"`
}

type OpenAIResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Created int `json:"created"`
	Model string `json:"model"`
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func (client *OpenAIClient) GetChatCompletion(messages []OpenAIChatMessage) (string, error) {
	reqBody, _ := json.Marshal(OpenAIChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
		MaxTokens: 100,
		Temperature: 1,
	})

    fmt.Println("reqBody: ")
    fmt.Println(bytes.NewBuffer(reqBody))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.client.Do(req)
	if err != nil {
	    fmt.Println("Error: ")
	    fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("OpenAI API returned non-200 status code")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	// Extracting the assistant's reply from the response
	if len(response["choices"].([]interface{})) > 0 {
		return response["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string), nil
	}

	return "", errors.New("No completions returned by OpenAI API")
}

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}
