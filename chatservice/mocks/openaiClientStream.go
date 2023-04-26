package mocks

import (
	"bufio"
	"context"
	"io"

	"github.com/henriquesbezerra/fclx/chatservice/internal/infra/grpc/pb"
	"google.golang.org/grpc"
)

type OpenAIClientStream struct {
}

type AiContentStream struct {
	Content string `json:"content"`
}

type AiMessageStream struct {
	Message AiContentStream
	Delta   AiContentStream
}

type OpenAIRequestStream struct {
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

type OpenAIClientResponses struct {
	Choices []AiMessage `json:"choices"`
}

// Implementa a interface grpc.ServerStream
type OpenAIStream struct {
	grpc.ServerStream
	isFinished bool
	messages   []pb.ChatResponse
	Choices    []AiMessageStream `json:"choices"`
	reader     *bufio.Reader
	index      int
}

// Implementa o método Send() da interface grpc.ServerStream
func (stream *OpenAIStream) Send(chatResponse *pb.ChatResponse) error {
	stream.messages = append(stream.messages, *chatResponse)
	return nil
}

// Implementa o método RecvMsg() da interface grpc.ServerStream
func (stream *OpenAIStream) RecvMsg(m interface{}) error {
	return nil
}

// Implementa o método SendMsg() da interface grpc.ServerStream
func (stream *OpenAIStream) SendMsg(m interface{}) error {
	return nil
}

func (oacrs *OpenAIStream) Recv() (ais *OpenAIStream, err error) {

	if oacrs.isFinished {
		err = io.EOF
		return
	}

	var maxLoop uint
	var fullmsg string

waitForData:

	fullmsg += oacrs.messages[maxLoop].Content + " "

	maxLoop++

	if maxLoop < 3 {
		goto waitForData
	}

	// fmt.Println("=> ", oacrs.Choices)
	if len(oacrs.Choices) >= 3 {
		oacrs.isFinished = true
	}

	m := AiMessageStream{
		Message: AiContentStream{Content: fullmsg}, Delta: AiContentStream{Content: fullmsg},
	}
	oacrs.Choices = append(oacrs.Choices, m)

	return oacrs, nil
}

func (oai *OpenAIClientStream) CreateChatCompletionStream(ctx context.Context, request OpenAIRequestStream) (*OpenAIStream, error) {
	responseStream := &OpenAIStream{}

	// Envia a primeira mensagem
	if err := responseStream.Send(&pb.ChatResponse{
		ChatId:  "1",
		UserId:  "1",
		Content: "Content1",
	}); err != nil {
		responseStream.isFinished = true
		return nil, io.EOF
	}

	// Envia a segunda mensagem
	if err := responseStream.Send(&pb.ChatResponse{
		ChatId:  "1",
		UserId:  "1",
		Content: "Content2",
	}); err != nil {
		return nil, err
	}

	// Envia a terceira mensagem
	if err := responseStream.Send(&pb.ChatResponse{
		ChatId:  "1",
		UserId:  "1",
		Content: "Content3",
	}); err != nil {
		return nil, err
	}

	// Envia uma última mensagem indicando que a stream foi concluída
	if err := responseStream.Send(&pb.ChatResponse{
		ChatId:  "1",
		UserId:  "1",
		Content: "fim da stream",
	}); err != nil {
		return nil, err
	}

	return responseStream, nil
}
