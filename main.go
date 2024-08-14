package main

import (
	handler "golangnext/api"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/chat/new", enableCORS(handler.ChatNew))
	http.HandleFunc("/chat/list", enableCORS(handler.ChatList))
	http.HandleFunc("/chat/send", enableCORS(handler.ChatSend))

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
