package gateway

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/audstanley/ollama-websocket-gateway/pkg/ollama"
	"github.com/gorilla/websocket"
)

type Server struct {
	config       Config
	upgrader     websocket.Upgrader
	ollamaClient *ollama.Client
}

func NewServer(cfg Config) *Server {
	return &Server{
		config: cfg,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		ollamaClient: ollama.NewClient(cfg.GetOllamaURL()),
	}
}

func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.handleWS)

	addr := s.config.GetServerAddr()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
	}()

	log.Printf("WebSocket gateway running on %s", addr)
	return server.ListenAndServe()
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) HandleMessage(ctx context.Context, conn *websocket.Conn, msg []byte) error {
	var req ollama.WSRequest
	if err := json.Unmarshal(msg, &req); err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("invalid JSON"))
		return err
	}

	streamer := ollama.NewStreamer(s.ollamaClient, conn)
	return streamer.StreamFromOllama(ctx, req)
}
