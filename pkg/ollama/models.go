package ollama

type Request struct {
	Model    string    `json:"model"`
	Stream   bool      `json:"stream"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type StreamChunk struct {
	Done    bool    `json:"done"`
	Message Message `json:"message"`
}

type WSRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}
