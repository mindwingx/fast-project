package main

import (
	"github.com/mindwingx/fast-project/pkg"
	"github.com/mindwingx/fast-project/process"
	"os"
)

func main() {
	path, _ := os.Getwd()

	items := process.NewItems()
	items.SetReader(*pkg.NewStd())
	items.SetPath(path)

	var (
		processes   = items.Process()
		setTitle    = processes.SetTitle()
		approvePath = processes.ApprovePath()
		setDirName  = processes.SetDirName()
		makeDir     = processes.MakeDir()
		createFile  = processes.CreateFile()
		serviceInit = processes.ServiceInit()
		gitInit     = processes.GitInit()
	)

	{
		approvePath.Next(&setTitle)
		setTitle.Next(&setDirName)
		setDirName.Next(&makeDir)
		makeDir.Next(&createFile)
		createFile.Next(&serviceInit)
		serviceInit.Next(&gitInit)
	}

	approvePath.Process(*items)
}
