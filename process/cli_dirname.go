package process

import (
	"fmt"
	"github.com/mindwingx/fast-project/constants"
	"regexp"
)

type DirName struct {
	next IFlow
}

func newDirName() *DirName {
	return &DirName{}
}

func (d *DirName) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true
	dirName := item.slug
	dirNameRegex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]*[a-zA-Z0-9]?$`)

	if match := dirNameRegex.MatchString(dirName); match == false {
		fmt.Println("🔄  invalid dir name. please enter again")
		fmt.Println(constants.StopMsg)
		return
	}

	if l := len(dirName); l > 128 {
		fmt.Println("🔄  dir name too long(max 128 chars). please try again")

		fmt.Println(constants.StopMsg)
		return
	}

	item.path = fmt.Sprintf("%s/%s", item.path, dirName)
	item.fail = false

	fmt.Println("✅  dir name set!")

	if d.next != nil {
		d.next.Process(item)
	} else {
		item.Done()
	}
}

func (d *DirName) Next(next IFlow) {
	d.next = next
}
