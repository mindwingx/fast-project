package CoR

import (
	"fmt"
	"os"
)

func MakeDir(path string) (ok bool) {
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("failed to create directory: %v\n", err)
		return
	}

	fmt.Println("✅ dir created!")

	ok = true
	return
}
