package process

import (
	"fmt"
	"os"
	"path/filepath"
)

type ScaffoldFile struct {
	name    string
	content []byte
	perm    os.FileMode
	next    IFlow
}

func newScaffoldFile() *ScaffoldFile {
	return &ScaffoldFile{}
}

func (s *ScaffoldFile) Process(item Items) {
	if item.fail == true {
		return
	}

	item.fail = true

	var files []ScaffoldFile

	files = append(files,
		ScaffoldFile{
			name:    filepath.Join(item.path, ".gitignore"),
			content: []byte(".idea/\n.vscode/\n.DS_Store\n"),
			perm:    0644,
		},
		ScaffoldFile{
			name:    filepath.Join(item.path, "main.go"),
			content: []byte("package main\n\nimport (\n\t\"fmt\"\n\t\"net/http\"\n\trouter \"" + item.slug + "/internal/server/http\"\n)\n\nfunc main() {\n\trouter.Routes()\n\tfmt.Println(\"Server running on port :8080\")\n\t\n\terr := http.ListenAndServe(\"0.0.0.0:8080\", nil)\n\tif err != nil {\n\t\tfmt.Println(\"Error:\", err)\n\t}\n}\n"),
			perm:    0644,
		},
		ScaffoldFile{
			name:    filepath.Join(item.path, "README.md"),
			content: []byte("## " + item.title + " Service\n- Run service\n```shell\ngo run .\n```\n"),
			perm:    0644,
		},
		ScaffoldFile{
			name:    filepath.Join(item.path, "internal", "handler", "handshake.go"),
			content: []byte("package handler\n\nimport (\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"net/http\"\n\t\"sync/atomic\"\n)\n\nvar cnt uint64 = 0\n\nfunc Handshake(w http.ResponseWriter, r *http.Request) {\n\t atomic.AddUint64(&cnt, 1)\n\n\tfmt.Printf(\"\\rTotal Request: %d   \", cnt)\n\tresp := map[string]interface{}{\n        \"message\": \"handshake successful!\",\n    }\n\t\n\tw.Header().Set(\"Content-Type\", \"application/json\")\n    w.WriteHeader(http.StatusOK)\n\n    json.NewEncoder(w).Encode(resp)\n}\n"),
			perm:    0644,
		},
		ScaffoldFile{
			name:    filepath.Join(item.path, "internal", "server", "http", "router.go"),
			content: []byte("package http\n\nimport (\n\t\"net/http\"\n\t\"" + item.slug + "/internal/handler\"\n)\n\nfunc Routes() {\n\thttp.HandleFunc(\"/\", handler.Handshake)\n}\n"),
			perm:    0644,
		},
		ScaffoldFile{
			name:    filepath.Join(item.path, "pkg", "utils", "slug.go"),
			content: []byte("package utils\n\nimport \"strings\"\n\nfunc Slug(phrase string) (slug string) {\n\tslug = strings.ToLower(phrase)\n\tslug = strings.ReplaceAll(slug, \" \", \"-\")\n\treturn\n}\n"),
			perm:    0644,
		},
	)

	for _, f := range files {
		if err := os.WriteFile(f.name, f.content, f.perm); err != nil {
			fmt.Printf("failed to create file with content: %v\n", err)
			return
		}
	}

	fmt.Println("✅  files created!")

	item.fail = false

	if s.next != nil {
		s.next.Process(item)
	} else {
		item.Done()
	}
}

func (s *ScaffoldFile) Next(next IFlow) {
	s.next = next
}
