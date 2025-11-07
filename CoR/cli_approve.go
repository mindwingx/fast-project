package CoR

import (
	"fmt"
	"prj-init/constants"
	"prj-init/pkg"
	"strings"
)

func ApprovePath(reader *pkg.Std, path string) (p string, ok bool) {
	existLoop := false

	for i := 0; i < 3; i++ {
		input := reader.Ask("approve the current path(%s)? (y/n)", path)

		if input == "" {
			input = "y"
		}

		switch strings.ToLower(input) {
		case "y":
			existLoop = true
			fmt.Println("✅ path verified!")
			break
		case "n":
			fmt.Println(constants.StopMsg)
			return
		default:
			if i == 2 {
				fmt.Println(constants.StopMsg)
				return
			}

			fmt.Println("🔄invalid input. please enter 'y' or 'n'")
		}

		p = "."

		if i == 2 {
			fmt.Println(constants.StopMsg)
			return
		}

		if existLoop == true {
			existLoop = false
			break
		}
	}

	ok = true
	return
}
