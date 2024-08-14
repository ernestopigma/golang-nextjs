package goapi

import (
	"os"

	"github.com/sashabaranov/go-openai"
)

func NewOpenAIClient() *openai.Client {
	return openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}
