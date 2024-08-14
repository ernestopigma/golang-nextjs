package handler

import (
	"encoding/json"
	"fmt"
	"golangnext/goapi/models"
	"net/http"
	"os"

	"github.com/brianvoe/gofakeit/v7"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	u := models.User{
		ID:       1,
		Username: gofakeit.Name(),
		LastName: "test",
		Age:      25,
		Env:      os.Getenv("TEST_ENV"),
	}

	// Convert user struct to JSON string
	jsonString, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonString))
}
