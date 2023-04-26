package mocks

import (
	"context"
)

type OpenAIClient struct {
}

type AiContent struct {
	Content string `json:"content"`
}

type AiMessage struct {
	Message AiContent
}

type OpenAIClientResponse struct {
	Choices []AiMessage `json:"choices"`
}

type OpenAIRequest struct {
	Model            string         `json:"model"`
	Messages         []AiMessage    `json:"messages"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

func (oai *OpenAIClient) CreateChatCompletion(ctx context.Context, request OpenAIRequest) (*OpenAIClientResponse, error) {

	return &OpenAIClientResponse{
		Choices: []AiMessage{
			{Message: AiContent{Content: "msg1"}},
		},
	}, nil
}
