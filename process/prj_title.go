package process

import (
	"fmt"
	"github.com/mindwingx/fast-project/constants"
	"github.com/mindwingx/fast-project/pkg"
	"regexp"
)

type ProjectTitle struct {
	next IFlow
}

func newProjectTitle() *ProjectTitle {
	return &ProjectTitle{}
}

func (p *ProjectTitle) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true

	prjTitle := ""
	existLoop := false

	for i := 0; i < 3; i++ {
		if i == 2 {
			existLoop = true
		}

		prjTitle = item.reader.Ask("enter project title:")
		if prjTitle == "" {
			if existLoop == true {
				return
			}

			fmt.Println("🔄  no name entered. please try again")
			continue
		}

		titleRegx := regexp.MustCompile(`^[a-zA-Z0-9]([ ._'-]*[a-zA-Z0-9])*$`)

		if match := titleRegx.MatchString(prjTitle); match == false {
			fmt.Println("🔄  invalid title. please enter again")

			if existLoop == true {
				fmt.Println(constants.StopMsg)
				return
			}

			continue
		}

		if l := len(prjTitle); l > 128 {
			fmt.Println("🔄  title too long(max 128 chars). please try again")

			if existLoop == true {
				fmt.Println(constants.StopMsg)
				return
			}

			continue
		}

		item.title = prjTitle
		item.slug = pkg.Slug(prjTitle)

		if existLoop == true {
			existLoop = false
		}

		break
	}

	item.fail = false

	if p.next != nil {
		p.next.Process(item)
	} else {
		item.Done()
	}
}

func (p *ProjectTitle) Next(next IFlow) {
	p.next = next
}
