package goapi

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

type AIService interface {
	// TODO, not use open ai types on response
	GetMessagesChat(threadID string, runID string) (Response, error)
	CreateNewChat() (newChatResponse, error)
	CreateNewMessageChat(threadID string, content string, assistantID string) (newMessageResponse, error)
}

type OpenAIService struct {
	client *openai.Client
}

type Response struct {
	Messages []openai.Message `json:"messages"`
	Status   openai.RunStatus `json:"status"`
}

type newMessageResponse struct {
	RunID string `json:"runId"`
}

type newChatResponse struct {
	ThreadID string `json:"threadId"`
}

func NewOpenAIService(apiKey string) *OpenAIService {
	client := openai.NewClient(apiKey)
	return &OpenAIService{client: client}
}

func (openAI *OpenAIService) GetMessagesChat(threadID string, runID string) (Response, error) {
	response := Response{}
	messages, err := openAI.client.ListMessage(context.TODO(), threadID, nil, nil, nil, nil)
	if err != nil {
		return response, err
	}

	if runID != "" {
		resp, err := openAI.client.RetrieveRun(context.TODO(), threadID, runID)
		if err != nil {
			return response, err
		}
		response.Status = resp.Status
	}
	response.Messages = messages.Messages

	return response, nil
}

func (openAI *OpenAIService) CreateNewChat() (newChatResponse, error) {
	newChatResponse := newChatResponse{}
	thread, eror := openAI.client.CreateThread(context.Background(), openai.ThreadRequest{})
	if eror != nil {
		return newChatResponse, eror
	}

	newChatResponse.ThreadID = thread.ID

	return newChatResponse, nil
}

func (openAI *OpenAIService) CreateNewMessageChat(threadID string, content string, assistantID string) (newMessageResponse, error) {
	newMessageResponse := newMessageResponse{}

	_, err := openAI.client.CreateMessage(context.TODO(), threadID, openai.MessageRequest{
		Role:    "user",
		Content: content,
	})
	if err != nil {
		return newMessageResponse, err
	}

	run, err := openAI.client.CreateRun(context.TODO(), threadID, openai.RunRequest{
		AssistantID: assistantID,
	})
	if err != nil {
		return newMessageResponse, err
	}

	newMessageResponse.RunID = run.ID

	return newMessageResponse, nil
}

func GetAIService() AIService {
	return NewOpenAIService(os.Getenv("OPENAI_API_KEY"))

}
