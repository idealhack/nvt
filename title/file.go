package title

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	lineBreak           = "\n"
	filePermission      = 0644
	codeIndicator       = "    "
	titledLinkIndicator = "](http"
	httpLinkIndicator   = "http://"
	httpsLinkIndicator  = "https://"
)

// ProcessFile ...
func ProcessFile(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), lineBreak)
	for i, line := range lines {
		lines[i] = processLine(line)
	}
	result := strings.Join(lines, lineBreak)
	err = ioutil.WriteFile(filename, []byte(result), filePermission)
	if err != nil {
		log.Fatal(err)
	}
}

func processLine(line string) (result string) {
	result = line
	if strings.Contains(line, titledLinkIndicator) || strings.Contains(line, codeIndicator) {
		return
	}
	if strings.Contains(line, httpLinkIndicator) || strings.Contains(line, httpsLinkIndicator) {
		tokens := strings.Split(line, " ")
		for i, token := range tokens {
			tokens[i] = processToken(token)
		}
		result = strings.Join(tokens, " ")
	}
	return
}

func processToken(token string) (result string) {
	result = token
	if strings.Contains(token, httpLinkIndicator) || strings.Contains(token, httpsLinkIndicator) {
		var err error
		if result, err = addTitleToLink(token); err != nil {
			result = token
		}
	}
	return
}

func addTitleToLink(url string) (result string, err error) {
	result = url
	title, err := fetchTitleOfURL(url)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return
	}
	result = "[" + title + "](" + url + ")"
	log.Printf("DONE: %s", result)
	return
}

// fuctions below forked from
// https://siongui.github.io/2016/05/10/go-get-html-title-via-net-html/

func fetchTitleOfURL(url string) (title string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "html") || resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	if title, ok := getHTMLTitle(resp.Body); ok {
		return title, nil
	}
	log.Printf("Failed to get HTML title")

	return
}

func getHTMLTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Printf("Failed to parse html")
	}

	return traverse(doc)
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) && n.FirstChild != nil {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}
