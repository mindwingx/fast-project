package process

import (
	"fmt"
	"github.com/mindwingx/fast-project/constants"
	"strings"
)

type ApprovePath struct {
	next IFlow
}

func newApprovePath() *ApprovePath {
	return &ApprovePath{}
}

func (a *ApprovePath) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true
	existLoop := false

	for i := 0; i < 3; i++ {
		input := item.reader.Ask("approve the current path(%s)? (y/n)", item.path)

		if input == "" {
			input = "y"
		}

		switch strings.ToLower(input) {
		case "y":
			existLoop = true
			fmt.Println("✅  path verified!")
			break
		case "n":
			fmt.Println(constants.StopMsg)
			return
		default:
			if i == 2 {
				fmt.Println(constants.StopMsg)
				return
			}

			fmt.Println("🔄  invalid input. please enter 'y' or 'n'")
		}

		item.path = "."

		if i == 2 {
			fmt.Println(constants.StopMsg)
			return
		}

		if existLoop == true {
			existLoop = false
			break
		}
	}

	item.fail = false

	if a.next != nil {
		a.next.Process(item)
	} else {
		item.Done()
	}
}

func (a *ApprovePath) Next(next IFlow) {
	a.next = next
}
