package process

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type GitInit struct {
	next IFlow
}

func newGitInit() *GitInit {
	return &GitInit{}
}

func (g *GitInit) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true

	gitDir := filepath.Join(item.path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		fmt.Println(".git directory already exists → Git repo already initialized!")
		fmt.Println("✅  Skipping git init")
		return
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = item.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git init failed: %v", err)
		return
	}

	cmd = exec.Command("git", "add", ".")
	cmd.Dir = item.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git add to stage failed: %v", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", "\"init\"")
	cmd.Dir = item.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git add to stage failed: %v", err)
		return
	}

	cmd = exec.Command("git", "branch", "-m", "main")
	cmd.Dir = item.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git add to stage failed: %v", err)
		return
	}

	fmt.Println("✅  git initialized!")

	item.fail = false

	if g.next != nil {
		g.next.Process(item)
	} else {
		item.Done()
	}
}

func (g *GitInit) Next(next IFlow) {
	g.next = next
}
