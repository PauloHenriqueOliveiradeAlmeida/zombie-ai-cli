package ai_models

import (
	"context"
	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
}

func NewGemini(key string) (*Gemini, error) {
	genai, error := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})

	if error != nil {
		return nil, error
	}

	return &Gemini{
		client: genai,
	}, nil
}

func (this *Gemini) GetResponse(prompt string, maxTokens int) (string, error) {
	response, error := this.client.Models.GenerateContent(
		context.Background(),
		"gemini-2.0-flash",
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			MaxOutputTokens: int32(maxTokens),
		},
	)

	if error != nil {
		return "", error
	}

	return response.Text(), nil
}
