package main

import (
	//"fmt"
	//"log"
	"strings"
	"unicode"
	"bytes"

	"golang.org/x/net/html"
)

func Extract(input []byte) ([]string, []string, error) {
    var words, hrefs []string
    doc, err := html.Parse(bytes.NewReader(input))
    if err != nil {
        return nil, nil, err
    }

    var f func(*html.Node)
    f = func(n *html.Node) {
		// added skip to style and title tags
		if n.Type == html.ElementNode && n.Data == "style"{
			return
		}
		
		if n.Type == html.TextNode {
			checkWords := func(c rune) bool{
				return !unicode.IsLetter(c) && !unicode.IsNumber(c)
			}
			words = append(words, strings.FieldsFunc(n.Data, checkWords)...)
		} else if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					hrefs = append(hrefs, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
    f(doc)

    return words, hrefs, nil
}