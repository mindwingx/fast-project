package process

import (
	"fmt"
	"github.com/mindwingx/fast-project/pkg"
)

type (
	Items struct {
		reader  pkg.Std
		path    string
		title   string
		slug    string
		fail    bool
		done    string
		process Processes
	}

	Processes struct {
		approvePath ApprovePath
		setDirName  DirName
		createFile  ScaffoldFile
		gitInit     GitInit
		makeDir     MakeDir
		setTitle    ProjectTitle
		serviceInit ServiceInit
	}
)

func NewItems() *Items {
	items := Items{done: "✅  project initialized!"}
	items.process = Processes{
		approvePath: *newApprovePath(),
		setDirName:  *newDirName(),
		createFile:  *newScaffoldFile(),
		gitInit:     *newGitInit(),
		makeDir:     *newMakeDir(),
		setTitle:    *newProjectTitle(),
		serviceInit: *newServiceInit(),
	}

	return &items
}

func (i *Items) Reader() pkg.Std {
	return i.reader
}

func (i *Items) SetReader(reader pkg.Std) {
	i.reader = reader
}

func (i *Items) Path() string {
	return i.path
}

func (i *Items) SetPath(path string) {
	i.path = path
}

func (i *Items) Title() string {
	return i.title
}

func (i *Items) SetTitle(title string) {
	i.title = title
}

func (i *Items) Slug() string {
	return i.slug
}

func (i *Items) SetSlug(slug string) {
	i.slug = slug
}

func (i *Items) Done() {
	fmt.Println(i.done)
}

//

func (i *Items) Process() Processes {
	return i.process
}

func (p *Processes) ApprovePath() ApprovePath {
	return p.approvePath
}

func (p *Processes) SetDirName() DirName {
	return p.setDirName
}

func (p *Processes) CreateFile() ScaffoldFile {
	return p.createFile
}

func (p *Processes) GitInit() GitInit {
	return p.gitInit
}

func (p *Processes) MakeDir() MakeDir {
	return p.makeDir
}

func (p *Processes) SetTitle() ProjectTitle {
	return p.setTitle
}

func (p *Processes) ServiceInit() ServiceInit {
	return p.serviceInit
}
