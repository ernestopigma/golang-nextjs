package main

import (
	handler "golangnext/api"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/users", handler.Users)
	http.HandleFunc("/chat-new", handler.ChatNew)
	http.HandleFunc("/chat-list", handler.ChatList)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
