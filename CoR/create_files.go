package CoR

import (
	"fmt"
	"os"
	"path/filepath"
)

type scaffoldFile struct {
	name    string
	content []byte
	perm    os.FileMode
}

func CreateFiles(path, title string) (ok bool) {
	var files []scaffoldFile

	files = append(files,
		scaffoldFile{
			name:    filepath.Join(path, ".gitignore"),
			content: []byte(".idea/\n.vscode/\n.DS_Store\n"),
			perm:    0755,
		},
		scaffoldFile{
			name:    filepath.Join(path, "main.go"),
			content: []byte("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"test message\")\n}"),
			perm:    0755,
		},
		scaffoldFile{
			name:    filepath.Join(path, "README.md"),
			content: []byte("## " + title + " Service\n- Run service\n```shell\ngo run .\n```\n"),
			perm:    0755,
		},
	)

	for _, f := range files {
		if err := os.WriteFile(f.name, f.content, f.perm); err != nil {
			fmt.Printf("failed to create file with content: %v\n", err)
			return
		}
	}

	fmt.Println("✅ files created!")

	ok = true
	return
}
