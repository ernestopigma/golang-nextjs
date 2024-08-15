package handler

import (
	"encoding/json"
	"fmt"
	"golangnext/goapi"
	"net/http"
)

type Body struct {
	ThreadID string `json:"threadId"`
	RunID    string `json:"runId"`
}

func ChatList(w http.ResponseWriter, r *http.Request) {
	aiService := goapi.GetAIService()

	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chatListresponse, err := aiService.GetMessagesChat(body.ThreadID, body.RunID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonString, err := json.Marshal(chatListresponse)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	goapi.ResponseJson(w, jsonString, http.StatusOK)

}
