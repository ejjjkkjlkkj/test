package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const DefaultHost = "http://127.0.0.1:11434"
const DefaultModel = "qwen3.8-coder:latest"

type Client struct {
	Host       string
	Model      string
	HTTPClient *http.Client
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type chatResponse struct {
	Message Message `json:"message"`
	Done    bool    `json:"done"`
	Error   string  `json:"error,omitempty"`
}

func NewClient(host, model string) *Client {
	if strings.TrimSpace(host) == "" {
		host = DefaultHost
	}
	if strings.TrimSpace(model) == "" {
		model = DefaultModel
	}
	return &Client{
		Host:  strings.TrimRight(host, "/"),
		Model: model,
		HTTPClient: &http.Client{Timeout: 10 * time.Minute},
	}
}

func (c *Client) Chat(ctx context.Context, messages []Message) (string, error) {
	if c == nil {
		return "", errors.New("nil Ollama client")
	}
	if strings.TrimSpace(c.Host) == "" || strings.TrimSpace(c.Model) == "" {
		return "", errors.New("Ollama host and model are required")
	}
	if len(messages) == 0 {
		return "", errors.New("at least one message is required")
	}
	body, err := json.Marshal(chatRequest{Model: c.Model, Messages: messages, Stream: false})
	if err != nil {
		return "", fmt.Errorf("encode Ollama request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Host+"/api/chat", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("create Ollama request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("connect to Ollama at %s: %w", c.Host, err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read Ollama response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("Ollama HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}
	var decoded chatResponse
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return "", fmt.Errorf("decode Ollama response: %w", err)
	}
	if decoded.Error != "" {
		return "", errors.New(decoded.Error)
	}
	if strings.TrimSpace(decoded.Message.Content) == "" {
		return "", errors.New("Ollama returned an empty response")
	}
	return decoded.Message.Content, nil
}

func (c *Client) Ask(ctx context.Context, prompt string) (string, error) {
	prompt = strings.TrimSpace(prompt)
	if prompt == "" {
		return "", errors.New("prompt is required")
	}
	return c.Chat(ctx, []Message{{Role: "user", Content: prompt}})
}
