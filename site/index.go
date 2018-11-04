package site

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/template"
)

// Index is a list of note titles and links
type Index struct {
	Title string
	Notes []Note
}

func processIndex(path string, notes []Note) {
	t, err := template.ParseFiles(indexTemplate)
	Check(err)

	indexData := Index{
		Title: siteTitle,
		Notes: notes,
	}

	indexFile, err := os.Create(filepath.Join(publicDirectory, htmlFileName))
	Check(err)

	err = t.Execute(indexFile, indexData)
	Check(err)
}
