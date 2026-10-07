package ai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChatUsesOllamaAPI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/chat" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("method = %q", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":{"role":"assistant","content":"OK"},"done":true}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-model")
	got, err := client.Ask(context.Background(), "ping")
	if err != nil {
		t.Fatal(err)
	}
	if got != "OK" {
		t.Fatalf("response = %q", got)
	}
}

func TestChatRejectsEmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":{"role":"assistant","content":""},"done":true}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-model")
	if _, err := client.Ask(context.Background(), "ping"); err == nil {
		t.Fatal("expected empty response error")
	}
}
