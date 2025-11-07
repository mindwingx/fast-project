package process

import (
	"fmt"
	"os"
	"path/filepath"
)

type MakeDir struct {
	next IFlow
}

func newMakeDir() *MakeDir {
	return &MakeDir{}
}

func (m *MakeDir) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true

	if err := os.MkdirAll(item.path, 0755); err != nil {
		fmt.Printf("failed to create directory: %v\n", err)
		return
	}

	if err := os.MkdirAll(filepath.Join(item.path, "internal", "handler"), 0755); err != nil {
		fmt.Printf("failed to create directory: %v\n", err)
		return
	}

	if err := os.MkdirAll(filepath.Join(item.path, "internal", "server", "http"), 0755); err != nil {
		fmt.Printf("failed to create directory: %v\n", err)
		return
	}

	if err := os.MkdirAll(filepath.Join(item.path, "pkg", "utils"), 0755); err != nil {
		fmt.Printf("failed to create directory: %v\n", err)
		return
	}

	fmt.Println("✅  dir created!")

	item.fail = false

	if m.next != nil {
		m.next.Process(item)
	} else {
		item.Done()
	}
}

func (m *MakeDir) Next(next IFlow) {
	m.next = next
}
