package site

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/russross/blackfriday"
)

// Note is a text file has title and path
type Note struct {
	Title   string
	Path    string
	Content string
}

// ProcessNotes ...
func ProcessNotes(path string) {
	files, err := ioutil.ReadDir(path)
	Check(err)

	sortByModTime(files)

	notes := []Note{}
	for _, f := range files {
		if !isMarkdownFile(f.Name()) {
			continue
		}

		processNote(path, f.Name())

		noteTitle := generateTitleFromPath(strings.TrimSuffix(f.Name(), noteExtenstion))

		note := Note{Title: noteTitle, Path: generateURLPathFromTitle(noteTitle)}
		notes = append(notes, note)
	}

	processIndex(path, notes)
}

func sortByModTime(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Unix() > files[j].ModTime().Unix()
	})
}

func isMarkdownFile(filename string) bool {
	return strings.HasSuffix(filename, noteExtenstion)
}

func processNote(path, filename string) {
	markdownBytes, err := ioutil.ReadFile(filepath.Join(path, filename))
	Check(err)

	noteTitle := strings.TrimSuffix(filename, noteExtenstion)
	noteContent := parseMarkdown(markdownBytes)

	note := Note{
		Title:   fmt.Sprintf("%s • %s", noteTitle, siteTitle),
		Content: fmt.Sprintf("%s", noteContent),
	}
	htmlContent := renderNote(note)

	htmlPath := filepath.Join(publicDirectory, generateURLPathFromTitle(noteTitle))
	err = os.MkdirAll(htmlPath, os.ModePerm)
	Check(err)

	htmlFile := filepath.Join(htmlPath, htmlFileName)
	err = ioutil.WriteFile(htmlFile, htmlContent, 0644)
	Check(err)
}

func renderNote(note Note) []byte {
	t, err := template.ParseFiles(noteTemplate)
	Check(err)

	var noteContent bytes.Buffer
	err = t.Execute(&noteContent, note)
	Check(err)

	return noteContent.Bytes()
}

// parseMarkdown parse Markdown to HTML
func parseMarkdown(markdown []byte) []byte {
	markdownString := string(markdown)
	reWikiLink := regexp.MustCompile("\\[\\[(.*)]]")
	markdownString = reWikiLink.ReplaceAllStringFunc(markdownString, convertWikiLink)
	return blackfriday.Run([]byte(markdownString))
}

// convertWikiLink converts `[[Link]]` to `[Link](../Link/)`
func convertWikiLink(link string) string {
	title := link[2 : len(link)-2]
	return fmt.Sprintf("[%s](../%s/)", title, generateURLPathFromTitle(title))
}

// generateTitleFromPath returns the correct title of a note from its path
func generateTitleFromPath(s string) string {
	s = strings.Replace(s, ":", "/", -1)
	return s
}

// generateURLPathFromTitle returns the URL path of a note from its title
func generateURLPathFromTitle(s string) string {
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Replace(s, ":", "-", -1)
	s = strings.Replace(s, "/", "-", -1)
	return s
}
