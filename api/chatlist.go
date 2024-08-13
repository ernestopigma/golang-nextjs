package handler

import (
	"encoding/json"
	"net/http"
)

type Body struct {
	ThreadID string `json:"thread_id"`
	RunID    string `json:"run_id"`
}

func ChatList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

		// client.ListMessage(context.TODO(), os.Getenv(""), nil, nil, nil, nil)

		var body Body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body.ThreadID))
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
