package gateway

import (
	"testing"

	"github.com/audstanley/ollama-websocket-gateway/pkg/config"
)

func TestGatewayConfig(t *testing.T) {
	cfg := &config.Config{
		Server: struct {
			Port string `mapstructure:"port"`
			Host string `mapstructure:"host"`
		}{
			Port: "11435",
			Host: "localhost",
		},
		Ollama: struct {
			URL string `mapstructure:"url"`
		}{
			URL: "http://localhost:11434",
		},
		Logging: struct {
			Level string `mapstructure:"level"`
		}{
			Level: "info",
		},
	}

	gwConfig := NewGatewayConfig(cfg)

	if gwConfig.GetServerAddr() != "localhost:11435" {
		t.Errorf("Expected server addr 'localhost:11435', got '%s'", gwConfig.GetServerAddr())
	}

	if gwConfig.GetOllamaURL() != "http://localhost:11434" {
		t.Errorf("Expected Ollama URL 'http://localhost:11434', got '%s'", gwConfig.GetOllamaURL())
	}

	if gwConfig.GetLogLevel() != "info" {
		t.Errorf("Expected log level 'info', got '%s'", gwConfig.GetLogLevel())
	}
}
