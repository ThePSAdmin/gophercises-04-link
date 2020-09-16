package main

import (
	"flag"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file := flag.String("file", "ex1.html", "The html file to parse")
	flag.Parse()

	r, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var f func(n *html.Node)
	f = func(n *html.Node) {
		log.Printf("Data is %s", n.Data)

		if v, ok := getAttrValue(n.Attr, "href"); ok {
			log.Printf("Found a href with value %s", *v)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}

func getAttrValue(as []html.Attribute, k string) (*string, bool) {
	for _, a := range as {
		if a.Key == k {
			return &a.Val, true
		}
	}
	return nil, false
}

// Link is a struct which represents a link (a tag) within a html document
type Link struct {
	url  string
	text string
}
