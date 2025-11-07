package process

import (
	"fmt"
	"os"
	"os/exec"
)

type ServiceInit struct {
	next IFlow
}

func newServiceInit() *ServiceInit {
	return &ServiceInit{}
}

func (s *ServiceInit) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true

	cmd := exec.Command("go", "mod", "init", item.slug)
	cmd.Dir = item.path // <-- critical: run inside the new dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("go mod init failed (is Go installed?): %v", err)
		return
	}

	fmt.Println("✅   go mod initialized!")

	item.fail = false

	if s.next != nil {
		s.next.Process(item)
	} else {
		item.Done()
	}
}

func (s *ServiceInit) Next(next IFlow) {
	s.next = next
}
