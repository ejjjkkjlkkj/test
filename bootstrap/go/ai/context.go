package ai

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	DefaultContextMaxBytes = 256 * 1024
	DefaultFileMaxBytes    = 64 * 1024
)

type ContextOptions struct {
	MaxBytes     int64
	MaxFileBytes int64
}

type ContextFile struct {
	Path string
	Size int64
	Data string
}

type RepositoryContext struct {
	Root  string
	Files []ContextFile
	Bytes int64
}

func BuildRepositoryContext(root string, opts ContextOptions) (RepositoryContext, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return RepositoryContext{}, errors.New("repository root is required")
	}
	info, err := os.Stat(root)
	if err != nil {
		return RepositoryContext{}, fmt.Errorf("stat repository root: %w", err)
	}
	if !info.IsDir() {
		return RepositoryContext{}, errors.New("repository root is not a directory")
	}
	if opts.MaxBytes <= 0 {
		opts.MaxBytes = DefaultContextMaxBytes
	}
	if opts.MaxFileBytes <= 0 {
		opts.MaxFileBytes = DefaultFileMaxBytes
	}

	var paths []string
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return relErr
		}
		if entry.IsDir() {
			if rel != "." && isIgnoredDir(rel) {
				return filepath.SkipDir
			}
			return nil
		}
		if !entry.Type().IsRegular() || isIgnoredFile(rel) {
			return nil
		}
		paths = append(paths, rel)
		return nil
	})
	if err != nil {
		return RepositoryContext{}, fmt.Errorf("walk repository: %w", err)
	}
	sort.Strings(paths)

	out := RepositoryContext{Root: root}
	for _, rel := range paths {
		path := filepath.Join(root, rel)
		info, err := os.Stat(path)
		if err != nil {
			return RepositoryContext{}, fmt.Errorf("stat %s: %w", rel, err)
		}
		if info.Size() > opts.MaxFileBytes || out.Bytes+info.Size() > opts.MaxBytes {
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return RepositoryContext{}, fmt.Errorf("read %s: %w", rel, err)
		}
		if !isText(data) {
			continue
		}
		out.Files = append(out.Files, ContextFile{
			Path: filepath.ToSlash(rel),
			Size: info.Size(),
			Data: string(data),
		})
		out.Bytes += info.Size()
	}
	return out, nil
}

func (c RepositoryContext) Prompt() string {
	var b strings.Builder
	fmt.Fprintf(&b, "READ-ONLY REPOSITORY CONTEXT\nROOT: %s\nFILES: %d\nBYTES: %d\n", c.Root, len(c.Files), c.Bytes)
	for _, file := range c.Files {
		fmt.Fprintf(&b, "\n===== %s (%d bytes) =====\n%s\n", file.Path, file.Size, file.Data)
	}
	return b.String()
}

func isIgnoredDir(rel string) bool {
	first := strings.Split(filepath.ToSlash(rel), "/")[0]
	switch first {
	case ".git", ".hg", ".svn", "node_modules", "vendor":
		return true
	default:
		return false
	}
}

func isIgnoredFile(rel string) bool {
	name := strings.ToLower(filepath.Base(rel))
	switch {
	case name == ".ds_store", strings.HasSuffix(name, ".exe"), strings.HasSuffix(name, ".dll"), strings.HasSuffix(name, ".so"):
		return true
	default:
		return false
	}
}

func isText(data []byte) bool {
	if len(data) == 0 {
		return true
	}
	for _, b := range data {
		if b == 0 {
			return false
		}
	}
	return true
}
