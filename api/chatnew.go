package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sashabaranov/go-openai"
)

func ChatNew(w http.ResponseWriter, r *http.Request) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	thread, eror := client.CreateThread(context.Background(), openai.ThreadRequest{})
	if eror != nil {
		http.Error(w, eror.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, thread.ID)

}
