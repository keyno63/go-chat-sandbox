package main

import (
	"fmt"
	"net/http"
)

var chatRoom = NewChatRoom()

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", WebSocketHandler(chatRoom))

	go chatRoom.Run()

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
