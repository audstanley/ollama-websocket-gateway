package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/audstanley/ollama-websocket-gateway/pkg/config"
	"github.com/audstanley/ollama-websocket-gateway/pkg/gateway"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	version = "dev"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "ollama-gateway",
		Short:   "WebSocket gateway for Ollama API",
		Version: version,
		RunE:    runServer,
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ollama-gateway/config.yaml)")
	rootCmd.Flags().StringP("port", "p", "11435", "server port")
	rootCmd.Flags().StringP("host", "H", "0.0.0.0", "server host")
	rootCmd.Flags().String("ollama-url", "http://localhost:11434", "Ollama API URL")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runServer(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if port, _ := cmd.Flags().GetString("port"); cmd.Flags().Changed("port") {
		cfg.Server.Port = port
	}
	if host, _ := cmd.Flags().GetString("host"); cmd.Flags().Changed("host") {
		cfg.Server.Host = host
	}
	if ollamaURL, _ := cmd.Flags().GetString("ollama-url"); cmd.Flags().Changed("ollama-url") {
		cfg.Ollama.URL = ollamaURL
	}

	gatewayConfig := gateway.NewGatewayConfig(cfg)
	server := gateway.NewServer(gatewayConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	fmt.Printf("Starting ollama-gateway on %s\n", gatewayConfig.GetServerAddr())
	return server.Start(ctx)
}
