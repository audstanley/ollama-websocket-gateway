# Ollama WebSocket Gateway

A WebSocket gateway that provides real-time streaming access to the Ollama API. Designed for terminal UI integration and modular usage.

## Features

- 🚀 Real-time WebSocket streaming from Ollama API
- 📦 Modular, importable library design
- ⚙️ Cobra CLI with comprehensive flags
- 🐍 Viper configuration management
- 🧪 Full test coverage
- 🔧 Clean separation of concerns
- 📝 Context-aware error handling

## Quick Start

### Installation

```bash
go install github.com/audstanley/ollama-websocket-gateway/cmd/gateway@latest
```

### Running the Server

```bash
# Run with default settings
ollama-gateway

# Run with custom port
ollama-gateway --port 8080 --host localhost

# Run with custom Ollama URL
ollama-gateway --ollama-url http://localhost:11434
```

## Configuration

The gateway supports configuration via:

1. **Command-line flags** (highest priority)
2. **Environment variables** (prefix: `GATEWAY_`)
3. **Config files** (YAML)

### Environment Variables

```bash
GATEWAY_SERVER_PORT=8080
GATEWAY_SERVER_HOST=localhost
GATEWAY_OLLAMA_URL=http://localhost:11434
GATEWAY_LOGGING_LEVEL=debug
```

### Config File

Create `config.yaml`:

```yaml
server:
  port: "11436"
  host: "0.0.0.0"
ollama:
  url: "http://localhost:11434"
logging:
  level: "info"
```

## Library Usage

Import as a library in your Go projects:

```go
package main

import (
    "context"
    "log"
    
    "github.com/audstanley/ollama-websocket-gateway/pkg/config"
    "github.com/audstanley/ollama-websocket-gateway/pkg/gateway"
)

func main() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }
    
    // Create gateway
    gwConfig := gateway.NewGatewayConfig(cfg)
    server := gateway.NewServer(gwConfig)
    
    // Start server
    ctx := context.Background()
    if err := server.Start(ctx); err != nil {
        log.Fatal(err)
    }
}
```

## WebSocket API

Connect to `ws://localhost:11436/ws` and send JSON:

```json
{
    "model": "llama2",
    "prompt": "Hello, how are you?"
}
```

The server will stream responses back as text messages, ending with `[DONE]`.

## Development

### Building

```bash
make build
```

### Testing

```bash
make test
make test-cover
```

### Linting

```bash
make lint
make fmt
```

### Development Mode

```bash
make dev
```

## Project Structure

```
cmd/gateway/          # CLI entry point
pkg/gateway/          # WebSocket server and interfaces
pkg/ollama/          # Ollama client and streaming
pkg/config/          # Configuration management
internal/errors/     # Internal error types
```

## Dependencies

- [gorilla/websocket](https://github.com/gorilla/websocket) - WebSocket implementation
- [spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [spf13/viper](https://github.com/spf13/viper) - Configuration management

## License

MIT License - see [LICENSE](LICENSE) file for details.