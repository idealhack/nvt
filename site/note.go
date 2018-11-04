package site

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/russross/blackfriday"
)

// Note is a text file has title and link
type Note struct {
	Title string
	Link  string
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

		noteTitle := strings.TrimSuffix(f.Name(), noteExtenstion)

		note := Note{Title: noteTitle, Link: replaceSpaceWithDash(noteTitle)}
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
	htmlContent := parseMarkdown(markdownBytes)

	htmlPath := filepath.Join(publicDirectory, replaceSpaceWithDash(noteTitle))
	err = os.MkdirAll(htmlPath, os.ModePerm)
	Check(err)

	htmlFile := filepath.Join(htmlPath, htmlFileName)
	err = ioutil.WriteFile(htmlFile, htmlContent, 0644)
	Check(err)
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
	return fmt.Sprintf("[%s](../%s/)", title, replaceSpaceWithDash(title))
}

// replaceSpaceWithDash turns `internal link` to `internal-link`
func replaceSpaceWithDash(s string) string {
	return strings.Replace(s, " ", "-", -1)
}
