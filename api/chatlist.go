package handler

import (
	"context"
	"encoding/json"
	"golangnext/goapi"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type Body struct {
	ThreadID string `json:"threadId"`
	RunID    string `json:"runId"`
}

type Response struct {
	Messages []openai.Message `json:"messages"`
	Status   openai.RunStatus `json:"status"`
}

func ChatList(w http.ResponseWriter, r *http.Request) {

	client := goapi.NewOpenAIClient()

	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{}

	messages, err := client.ListMessage(context.TODO(), body.ThreadID, nil, nil, nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.RunID != "" {
		resp, err := client.RetrieveRun(context.TODO(), body.ThreadID, body.RunID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Status = resp.Status
	}
	response.Messages = messages.Messages

	jsonString, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	goapi.ResponseJson(w, jsonString, http.StatusOK)

}
