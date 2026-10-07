package ai

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildRepositoryContextIsBoundedAndReadOnly(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, ".git"), 0o755); err != nil { t.Fatal(err) }
	if err := os.WriteFile(filepath.Join(root, "main.go"), []byte("package main\n"), 0o644); err != nil { t.Fatal(err) }
	if err := os.WriteFile(filepath.Join(root, "binary.bin"), []byte{0, 1, 2}, 0o644); err != nil { t.Fatal(err) }
	if err := os.WriteFile(filepath.Join(root, "large.txt"), []byte(strings.Repeat("x", 20)), 0o644); err != nil { t.Fatal(err) }

	before, err := os.ReadFile(filepath.Join(root, "main.go"))
	if err != nil { t.Fatal(err) }

	ctx, err := BuildRepositoryContext(root, ContextOptions{MaxBytes: 100, MaxFileBytes: 10})
	if err != nil { t.Fatal(err) }
	if len(ctx.Files) != 1 || ctx.Files[0].Path != "main.go" {
		t.Fatalf("unexpected files: %+v", ctx.Files)
	}
	if ctx.Files[0].Data != string(before) {
		t.Fatalf("repository content changed during context build")
	}
	after, err := os.ReadFile(filepath.Join(root, "main.go"))
	if err != nil { t.Fatal(err) }
	if string(after) != string(before) {
		t.Fatal("context builder modified repository content")
	}
	if strings.Contains(ctx.Prompt(), "binary.bin") || strings.Contains(ctx.Prompt(), ".git") {
		t.Fatal("ignored content leaked into prompt")
	}
}

func TestRepositoryContextPromptIsDeterministic(t *testing.T) {
	root := t.TempDir()
	for name, data := range map[string]string{"b.txt": "B", "a.txt": "A"} {
		if err := os.WriteFile(filepath.Join(root, name), []byte(data), 0o644); err != nil { t.Fatal(err) }
	}
	first, err := BuildRepositoryContext(root, ContextOptions{})
	if err != nil { t.Fatal(err) }
	second, err := BuildRepositoryContext(root, ContextOptions{})
	if err != nil { t.Fatal(err) }
	if first.Prompt() != second.Prompt() {
		t.Fatal("repository context is not deterministic")
	}
}
