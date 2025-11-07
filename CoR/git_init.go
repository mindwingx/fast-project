package CoR

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InitGit(path string) (ok bool) {
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		fmt.Println(".git directory already exists → Git repo already initialized!")
		fmt.Println("✅ Skipping git init")
		return true
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git init failed: %v", err)
		return
	}

	cmd = exec.Command("git", "add", ".")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git add to stage failed: %v", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", "\"init\"")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git add to stage failed: %v", err)
		return
	}

	cmd = exec.Command("git", "branch", "-m", "main")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git add to stage failed: %v", err)
		return
	}

	fmt.Println("✅ git initialized!")

	ok = true
	return
}
