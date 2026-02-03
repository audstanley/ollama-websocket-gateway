package errors

import "errors"

var (
	ErrInvalidJSON      = errors.New("invalid JSON")
	ErrOllamaConnection = errors.New("failed to connect to Ollama")
	ErrWebSocketUpgrade = errors.New("failed to upgrade WebSocket connection")
	ErrConfigLoad       = errors.New("failed to load configuration")
)
