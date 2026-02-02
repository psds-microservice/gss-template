package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

var copyOpts = cp.Options{
	Skip: func(info os.FileInfo, src, dest string) (bool, error) {
		return strings.HasSuffix(filepath.Base(src), ".git"), nil
	},
}

// MkdirAll creates directory and parents. Returns error on failure.
func MkdirAll(p string) error {
	return os.MkdirAll(p, 0755)
}

// Exists returns true if path exists.
func Exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// IsGitRepo returns true if path is a git repository root (has .git).
func IsGitRepo(p string) bool {
	gitDir := filepath.Join(p, ".git")
	info, err := os.Stat(gitDir)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Gitkeep creates dir if needed and writes .gitkeep inside it.
func Gitkeep(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, ".gitkeep"))
	if err != nil {
		return err
	}
	return f.Close()
}

// RunCommand runs command with args in dir. Returns error on failure.
func RunCommand(name string, args []string, dir string) error {
	path, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("command %s not found: %w", name, err)
	}
	cmd := exec.Cmd{
		Path: path,
		Args: append([]string{path}, args...),
		Dir:  dir,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run %s: %w", name, err)
	}
	return nil
}

// Copy copies src to dest (file or directory). Skips .git.
func Copy(src, dest string) error {
	return cp.Copy(src, dest, copyOpts)
}
