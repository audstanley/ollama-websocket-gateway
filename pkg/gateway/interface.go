package gateway

import (
	"context"

	"github.com/audstanley/ollama-websocket-gateway/pkg/config"
	"github.com/gorilla/websocket"
)

type Gateway interface {
	Start(ctx context.Context) error
	Stop() error
}

type MessageHandler interface {
	HandleMessage(ctx context.Context, conn *websocket.Conn, msg []byte) error
}

type ServerInterface interface {
	Gateway
	MessageHandler
}

type Config interface {
	GetServerAddr() string
	GetOllamaURL() string
	GetLogLevel() string
}

func NewGatewayConfig(cfg *config.Config) Config {
	return &gatewayConfig{cfg: cfg}
}

type gatewayConfig struct {
	cfg *config.Config
}

func (g *gatewayConfig) GetServerAddr() string {
	return g.cfg.Server.Host + ":" + g.cfg.Server.Port
}

func (g *gatewayConfig) GetOllamaURL() string {
	return g.cfg.Ollama.URL
}

func (g *gatewayConfig) GetLogLevel() string {
	return g.cfg.Logging.Level
}
