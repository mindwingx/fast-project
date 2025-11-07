package main

import (
	"fmt"
	"os"
	"prj-init/CoR"
	"prj-init/pkg"
)

func main() {
	var (
		reader  = pkg.NewStd()
		path, _ = os.Getwd()
	)

	title, slug, ok := CoR.SetProjectTitle(reader)
	if !ok {
		return
	}

	path, ok = CoR.ApprovePath(reader, path)
	if !ok {
		return
	}

	path, ok = CoR.SetDirName(path, slug)
	if !ok {
		return
	}

	ok = CoR.MakeDir(path)
	if !ok {
		return
	}

	ok = CoR.CreateFiles(path, title)
	if !ok {
		return
	}

	ok = CoR.InitService(path, slug)
	if !ok {
		return
	}

	ok = CoR.InitGit(path)
	if !ok {
		return
	}

	fmt.Println("✅ project initialized!")
}
