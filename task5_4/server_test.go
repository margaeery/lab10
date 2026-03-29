package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gorilla/websocket"
)

func TestWebSocketFlow(t *testing.T) {
	hub := &ChatHub{clients: make(map[*websocket.Conn]bool)}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleConnections(hub, w, r)
	}))
	defer server.Close()

	url := "ws" + strings.TrimPrefix(server.URL, "http") + "/chat"

	t.Run("ValidExchange", func(t *testing.T) {
		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer ws.Close()

		sent := Message{User: "User1", Content: "Test"}
		if err := ws.WriteJSON(sent); err != nil {
			t.Fatalf("%v", err)
		}

		var received Message
		if err := ws.ReadJSON(&received); err != nil {
			t.Fatalf("%v", err)
		}
		if received.User != sent.User || received.Content != sent.Content {
			t.Errorf("Mismatch: %v != %v", sent, received)
		}
	})

	t.Run("InvalidData", func(t *testing.T) {
		ws, _, _ := websocket.DefaultDialer.Dial(url, nil)
		defer ws.Close()

		if err := ws.WriteMessage(websocket.TextMessage, []byte("invalid")); err != nil {
			t.Fatalf("%v", err)
		}
	})
}