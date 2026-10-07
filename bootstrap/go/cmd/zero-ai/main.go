package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	zeroai "zero-bootstrap/ai"
)

func main() {
	host := flag.String("host", envOr("ZERO_AI_HOST", zeroai.DefaultHost), "Ollama URL")
	model := flag.String("model", envOr("ZERO_AI_MODEL", zeroai.DefaultModel), "Ollama model")
	prompt := flag.String("prompt", "", "single prompt; when omitted, read prompts from stdin")
	flag.Parse()

	client := zeroai.NewClient(*host, *model)
	system := zeroai.SystemPrompt

	if strings.TrimSpace(*prompt) != "" {
		runPrompt(client, system, *prompt)
		return
	}

	fmt.Fprintf(os.Stderr, "ZERO AI | Ollama=%s | model=%s
", client.Host, client.Model)
	fmt.Fprintln(os.Stderr, "Enter a prompt. Ctrl+Z then Enter exits on Windows; Ctrl+D exits on Unix.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prompt := strings.TrimSpace(scanner.Text())
		if prompt == "" {
			continue
		}
		runPrompt(client, system, prompt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
		os.Exit(1)
	}
}

func runPrompt(client *zeroai.Client, system, prompt string) {
	ctx := context.Background()
	answer, err := client.Chat(ctx, []zeroai.Message{
		{Role: "system", Content: system},
		{Role: "user", Content: prompt},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "AI error:", err)
		os.Exit(1)
	}
	fmt.Println(answer)
}

func envOr(name, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}

var _ io.Reader
