package CoR

import (
	"fmt"
	"prj-init/constants"
	"regexp"
)

func SetDirName(path, slug string) (dirname string, ok bool) {
	dirName := slug

	dirNameRegex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]*[a-zA-Z0-9]?$`)

	if match := dirNameRegex.MatchString(dirName); match == false {
		fmt.Println("🔄invalid dir name. please enter again")
		fmt.Println(constants.StopMsg)
		return
	}

	if l := len(dirName); l > 128 {
		fmt.Println("🔄dir name too long(max 128 chars). please try again")

		fmt.Println(constants.StopMsg)
		return
	}

	dirname = fmt.Sprintf("%s/%s", path, dirName)
	ok = true
	fmt.Println("✅ dir name set!")
	return
}
