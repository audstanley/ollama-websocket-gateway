package ollama

import (
	"encoding/json"
	"testing"
)

func TestWSRequest_Unmarshal(t *testing.T) {
	jsonData := `{"model": "llama2", "prompt": "Hello world"}`

	var req WSRequest
	err := json.Unmarshal([]byte(jsonData), &req)

	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if req.Model != "llama2" {
		t.Errorf("Expected model 'llama2', got '%s'", req.Model)
	}

	if req.Prompt != "Hello world" {
		t.Errorf("Expected prompt 'Hello world', got '%s'", req.Prompt)
	}
}

func TestRequest_Creation(t *testing.T) {
	req := Request{
		Model:  "llama2",
		Stream: true,
		Messages: []Message{
			{Role: "user", Content: "Hello"},
		},
	}

	if req.Model != "llama2" {
		t.Errorf("Expected model 'llama2', got '%s'", req.Model)
	}

	if !req.Stream {
		t.Error("Expected Stream to be true")
	}

	if len(req.Messages) != 1 {
		t.Fatalf("Expected 1 message, got %d", len(req.Messages))
	}

	msg := req.Messages[0]
	if msg.Role != "user" || msg.Content != "Hello" {
		t.Errorf("Expected message {user, Hello}, got {%s, %s}", msg.Role, msg.Content)
	}
}
