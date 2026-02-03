package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaRequest struct {
	Model  string        `json:"model"`
	Stream bool          `json:"stream"`
	Messages []OllamaMsg `json:"messages"`
}

type OllamaMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaStreamChunk struct {
	Done    bool   `json:"done"`
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handleWS)

	log.Println("WebSocket gateway running on :11435")
	log.Fatal(http.ListenAndServe(":11435", nil))
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		var req WSRequest
		if err := json.Unmarshal(msg, &req); err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("invalid JSON"))
			continue
		}

		go streamFromOllama(conn, req)
	}
}

func streamFromOllama(conn *websocket.Conn, req WSRequest) {
	body, _ := json.Marshal(OllamaRequest{
		Model:  req.Model,
		Stream: true,
		Messages: []OllamaMsg{
			{Role: "user", Content: req.Prompt},
		},
	})

	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(body))
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("error contacting ollama"))
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	for {
		var chunk OllamaStreamChunk
		if err := dec.Decode(&chunk); err != nil {
			return
		}

		if chunk.Message.Content != "" {
			conn.WriteMessage(websocket.TextMessage, []byte(chunk.Message.Content))
		}

		if chunk.Done {
			conn.WriteMessage(websocket.TextMessage, []byte("[DONE]"))
			return
		}
	}
}
