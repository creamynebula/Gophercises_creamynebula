package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// vai receber um documento html
// e retornar um slice de links que foram parsed do doc
func Parse(r io.Reader) ([]Link, error) {
	documento, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\nvalor de documento: %v\n", *documento)

	// node é um struct que tem vários fields tipo Data e Attr
	// https://pkg.go.dev/golang.org/x/net/html?utm_source=godoc#Node
	nodes := linkNodes(documento)
	fmt.Printf("\nDá uma olhada nos nodes:\n")
	for _, node := range nodes {
		fmt.Println(node)
	}

	// link é Href e Text
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	imprimeArvoreHtml(documento, "", true)
	return links, nil
}

// retorna um slice com pointers pros elementos <a> de um documento html,
// representado como uma árvore
func linkNodes(n *html.Node) []*html.Node {
	// se 'n' for um elemento html, especificamente o <a>
	if n.Type == html.ElementNode && n.Data == "a" {
		// retorna um slice de pointers pra Nodes contendo 'n'
		return []*html.Node{n}
	}

	var returnValue []*html.Node

	//se 'n' não for o <a>, fazer dept-first-search procurando <a>
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnValue = append(returnValue, linkNodes(c)...)
	}

	return returnValue
}

func imprimeArvoreHtml(n *html.Node, padding string, topoDaArvore bool) {
	if topoDaArvore {
		fmt.Printf("\nFunção 'imprimeArvoreHtml' chamada.\n")
	}

	msg := n.Data

	// se 'n' for uma tag html, vamos imprimi-la
	// com uma quantidade de espaçamento ("padding") apropriada
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)

	// depth-first search
	// 'c' de Child
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// aumentamos o padding porque, quanto mais fundo na html tree, mais espaçamento
		imprimeArvoreHtml(c, padding+"  ", false)
	}
}

func buildLink(n *html.Node) Link {
	var returnValue Link
	// cada nodo tem um array de atributos que chama Attr
	// e os atributos tem Key, Val, ...
	for _, atributo := range n.Attr {
		if atributo.Key == "href" {
			returnValue.Href = atributo.Val
			break // sai do loop, pq já deu bom
		}
	}
	returnValue.Text = text(n)
	return returnValue
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var returnValue string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnValue += text(c) + " "
	}

	return strings.Join(strings.Fields(returnValue), " ")

}
