package site

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	configFileName = "config.yaml"
	htmlFileName   = "index.html"
	noteExtenstion = ".md"
)

var (
	// NotesDirectory ...
	NotesDirectory string

	siteTitle       string
	configFile      string
	publicDirectory string
	staticDirectory string
	indexTemplate   string
	noteTemplate    string
)

// SetConfig ...
func SetConfig(wd string) {
	viper.SetConfigType("yaml")
	configFile = filepath.Join(wd, configFileName)
	config, err := ioutil.ReadFile(configFile)
	Check(err)

	err = viper.ReadConfig(bytes.NewBuffer(config))
	Check(err)
	siteTitle = viper.GetString("title")

	NotesDirectory = filepath.Join(wd, "notes")
	publicDirectory = filepath.Join(wd, "public")
	staticDirectory = filepath.Join(wd, "static")
	indexTemplate = filepath.Join(wd, "theme", "index.html")
	noteTemplate = filepath.Join(wd, "theme", "note.html")
}

// Check ...
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
