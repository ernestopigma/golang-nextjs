package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"golangnext/goapi"
	"net/http"
	"os"

	"github.com/sashabaranov/go-openai"
)

type BodyChatNew struct {
	ThreadID string `json:"threadId"`
	Text     string `json:"text"`
}

func ChatSend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		client := goapi.NewOpenAIClient()

		var body BodyChatNew
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = client.CreateMessage(context.TODO(), body.ThreadID, openai.MessageRequest{
			Role:    "user",
			Content: body.Text,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		run, err := client.CreateRun(context.TODO(), body.ThreadID, openai.RunRequest{
			AssistantID: os.Getenv("ASSISTANT_ID"),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "{runId: %s}", run.ID)
		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
