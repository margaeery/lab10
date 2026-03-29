package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

type ChatHub struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

func (h *ChatHub) broadcast(msg []byte) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	for client := range h.clients {
		client.WriteMessage(websocket.TextMessage, msg)
	}
}

func handleConnections(hub *ChatHub, w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	
	hub.mutex.Lock()
	hub.clients[conn] = true
	hub.mutex.Unlock()

	defer func() {
		hub.mutex.Lock()
		delete(hub.clients, conn)
		hub.mutex.Unlock()
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var m Message
		if err := json.Unmarshal(msg, &m); err == nil {
			fmt.Printf("[%s]: %s\n", m.User, m.Content)
			hub.broadcast(msg)
		}
	}
}

func main() {
	hub := &ChatHub{clients: make(map[*websocket.Conn]bool)}
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(hub, w, r)
	})
	fmt.Println("Чат-сервер запущен на :8080")
	http.ListenAndServe(":8080", nil)
}