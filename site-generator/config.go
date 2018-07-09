package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	configFileName = "config.yaml"
	htmlFileName   = "index.html"
	noteExtenstion = ".md"
)

var (
	siteTitle       string
	configFile      string
	notesDirectory  string
	publicDirectory string
	staticDirectory string
	indexTemplate   string
)

func setConfig(wd string) {
	viper.SetConfigType("yaml")
	configFile = filepath.Join(wd, configFileName)
	config, err := ioutil.ReadFile(configFile)
	check(err)

	viper.ReadConfig(bytes.NewBuffer(config))
	siteTitle = viper.GetString("title")

	notesDirectory = filepath.Join(wd, "wiki")
	publicDirectory = filepath.Join(wd, "public")
	staticDirectory = filepath.Join(wd, "static")
	indexTemplate = filepath.Join(wd, "theme", "index.html")
}
