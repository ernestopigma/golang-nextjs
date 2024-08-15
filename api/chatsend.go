package handler

import (
	"encoding/json"
	"fmt"
	"golangnext/goapi"
	"net/http"
	"os"
)

type BodyChatNew struct {
	ThreadID string `json:"threadId"`
	Text     string `json:"text"`
}

func ChatSend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		aiService := goapi.GetAIService()

		var body BodyChatNew
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newMessageResponse, err := aiService.CreateNewMessageChat(body.ThreadID, body.Text, os.Getenv("ASSISTANT_ID"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonString, err := json.Marshal(newMessageResponse)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		goapi.ResponseJson(w, jsonString, http.StatusOK)
		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
