// site-generator takes nvALT notes files and generates a static wiki-like website.
package main

import (
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	check(err)

	setConfig(wd)
	processNotes(notesDirectory)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
