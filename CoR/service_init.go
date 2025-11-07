package CoR

import (
	"fmt"
	"os"
	"os/exec"
)

func InitService(path, moduleName string) (ok bool) {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = path // <-- critical: run inside the new dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("go mod init failed (is Go installed?): %v", err)
		return
	}

	fmt.Println("✅ go mod initialized!")

	ok = true
	return
}
