package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// vai receber um documento html
// e retornar um slice de links que foram parsed do doc
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}

	dfs(doc, "")
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	// se for um elemento html, especificamente o <a>
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var returnValue []*html.Node

	//se não for o <a>, fazer dept-first-search DFS
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnValue = append(returnValue, linkNodes(c)...)
	}

	return returnValue
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
