package main

import (
	"golang.org/x/net/websocket"
	"net/http"
)

func WebSocketHandler(chatRoom *ChatRoom) http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		client := &Client{ws: ws, send: make(chan []byte)}
		chatRoom.register <- client
		defer func() { chatRoom.unregister <- client }()
		go client.writePump()
		client.readPump()
	})
}

func (c *Client) readPump() {
	for {
		var message []byte
		err := websocket.Message.Receive(c.ws, &message)
		if err != nil {
			break
		}
		chatRoom.broadcast <- message
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			err := websocket.Message.Send(c.ws, string(message))
			if err != nil {
				return
			}
		}
	}
}
