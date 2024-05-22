package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func traverseHTML(n *html.Node, targetElement, targetClass string) {
	if n.Type == html.ElementNode && n.Data == targetElement {
		for _, a := range n.Attr {
			if a.Key == "class" && strings.Contains(a.Val, targetClass) {

				fmt.Println(extractText(n))
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTML(c, targetElement, targetClass)
	}
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return text
}

func main() {
	url := "https://www.chrismytton.com/plain-text-websites/"

	doc, err := fetchHTML(url)
	if err != nil {
		fmt.Printf("Error fetching HTML: %v\n", err)
		return
	}

	targetElement := "div"
	targetClass := "post"
	traverseHTML(doc, targetElement, targetClass)
}
