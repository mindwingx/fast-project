package CoR

import (
	"fmt"
	"prj-init/constants"
	"prj-init/pkg"
	"regexp"
)

func SetProjectTitle(reader *pkg.Std) (title, slug string, ok bool) {
	prjTitle := ""
	existLoop := false

	for i := 0; i < 3; i++ {
		if i == 2 {
			existLoop = true
		}

		prjTitle = reader.Ask("enter project title:")
		if prjTitle == "" {
			if existLoop == true {
				return
			}

			fmt.Println("🔄no name entered. please try again")
			continue
		}

		titleRegx := regexp.MustCompile(`^[a-zA-Z0-9]([ ._'-]*[a-zA-Z0-9])*$`)

		if match := titleRegx.MatchString(prjTitle); match == false {
			fmt.Println("🔄invalid title. please enter again")

			if existLoop == true {
				fmt.Println(constants.StopMsg)
				return
			}

			continue
		}

		if l := len(prjTitle); l > 128 {
			fmt.Println("🔄title too long(max 128 chars). please try again")

			if existLoop == true {
				fmt.Println(constants.StopMsg)
				return
			}

			continue
		}

		title = prjTitle
		slug = pkg.Slug(prjTitle)

		if existLoop == true {
			existLoop = false
		}

		break
	}

	ok = true
	return
}
