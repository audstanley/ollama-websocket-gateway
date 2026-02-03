package ollama

import (
	"context"

	"github.com/gorilla/websocket"
)

type Streamer struct {
	client *Client
	conn   *websocket.Conn
}

func NewStreamer(client *Client, conn *websocket.Conn) *Streamer {
	return &Streamer{
		client: client,
		conn:   conn,
	}
}

func (s *Streamer) StreamFromOllama(ctx context.Context, req WSRequest) error {
	ollamaReq := Request{
		Model:  req.Model,
		Stream: true,
		Messages: []Message{
			{Role: "user", Content: req.Prompt},
		},
	}

	chunkChan, errChan := s.client.StreamChat(ctx, ollamaReq)

	for {
		select {
		case chunk, ok := <-chunkChan:
			if !ok {
				return nil
			}
			if chunk.Message.Content != "" {
				if err := s.conn.WriteMessage(websocket.TextMessage, []byte(chunk.Message.Content)); err != nil {
					return err
				}
			}
			if chunk.Done {
				return s.conn.WriteMessage(websocket.TextMessage, []byte("[DONE]"))
			}

		case err := <-errChan:
			if err != nil {
				return s.conn.WriteMessage(websocket.TextMessage, []byte("error: "+err.Error()))
			}
			return nil

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
