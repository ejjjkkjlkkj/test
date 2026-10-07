package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	zeroai "zero-bootstrap/ai"
)

func main() {
	repo := flag.String("repo", ".", "repository path")
	model := flag.String("model", zeroai.DefaultModel, "Ollama model")
	host := flag.String("host", zeroai.DefaultHost, "Ollama URL")
	prompt := flag.String("prompt", "Analyse ce dépôt en lecture seule. Sépare faits, inférences, éléments non prouvés et risques.", "analysis prompt")
	flag.Parse()

	contextData, err := zeroai.BuildRepositoryContext(*repo, zeroai.ContextOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "repository context error:", err)
		os.Exit(1)
	}

	client := zeroai.NewClient(*host, *model)
	answer, err := client.Chat(context.Background(), []zeroai.Message{
		{Role: "system", Content: zeroai.SystemPrompt},
		{Role: "user", Content: "Analyse le contexte fourni en lecture seule. N'exécute aucune action et ne prétends pas avoir modifié ou testé quoi que ce soit.\n\n" + contextData.Prompt() + "\n\nDEMANDE:\n" + strings.TrimSpace(*prompt)},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "AI error:", err)
		os.Exit(1)
	}
	fmt.Println(answer)
}
