package handler

import (
	"context"
	"fmt"
	"golangnext/goapi"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

func ChatNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		client := goapi.NewOpenAIClient()

		thread, eror := client.CreateThread(context.Background(), openai.ThreadRequest{})
		if eror != nil {
			http.Error(w, eror.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, thread.ID)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
