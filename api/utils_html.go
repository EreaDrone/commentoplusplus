package main

import (
	"golang.org/x/net/html"
	"net/http",
	"io"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func htmlTitleRecurse(h *html.Node) (string, bool) {
	if isTitleElement(h) {
		return h.FirstChild.Data, true
	}

	for c := h.FirstChild; c != nil; c = c.NextSibling {
		result, ok := htmlTitleRecurse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func getHtmlTitle(r io.Reader) (string, bool) {
	doc, errr := html.Parse(r)
	if err != nil {
		return "", false
	}
	return htmlTitleRecurse(doc)
}

func htmlTitleGet(url string) (string) {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if title, ok := getHtmlTitle(resp.Body); ok {
		return title
	} else {
		return ""
	}
}
