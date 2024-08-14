package handler

import (
	"context"
	"encoding/json"
	"golangnext/goapi"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type ResponseChatNew struct {
	ThreadID string `json:"threadId"`
}

func ChatNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		client := goapi.NewOpenAIClient()

		thread, eror := client.CreateThread(context.Background(), openai.ThreadRequest{})
		if eror != nil {
			http.Error(w, eror.Error(), http.StatusInternalServerError)
			return
		}

		respChatNew := ResponseChatNew{
			ThreadID: thread.ID,
		}

		jsonString, err := json.Marshal(respChatNew)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		goapi.ResponseJson(w, jsonString, http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
