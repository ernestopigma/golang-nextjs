package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	u := User{
		ID:       1,
		Username: "johndoe",
	}

	// Convert user struct to JSON string
	jsonString, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonString))
}
