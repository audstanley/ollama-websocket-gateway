package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func resetViper() {
	viper.Reset()
}

func TestConfig_LoadDefaults(t *testing.T) {
	defer resetViper()
	backup := os.Getenv("GATEWAY_SERVER_PORT")
	defer os.Setenv("GATEWAY_SERVER_PORT", backup)

	os.Unsetenv("GATEWAY_SERVER_PORT")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Server.Port != "11435" {
		t.Errorf("Expected default port '11435', got '%s'", cfg.Server.Port)
	}

	if cfg.Server.Host != "0.0.0.0" {
		t.Errorf("Expected default host '0.0.0.0', got '%s'", cfg.Server.Host)
	}

	if cfg.Ollama.URL != "http://localhost:11434" {
		t.Errorf("Expected default Ollama URL 'http://localhost:11434', got '%s'", cfg.Ollama.URL)
	}
}

func TestConfig_EnvOverride(t *testing.T) {
	defer resetViper()
	backup := os.Getenv("GATEWAY_SERVER_PORT")
	defer os.Setenv("GATEWAY_SERVER_PORT", backup)

	os.Setenv("GATEWAY_SERVER_PORT", "8080")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected port '8080', got '%s'", cfg.Server.Port)
	}
}
