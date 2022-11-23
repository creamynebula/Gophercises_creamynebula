package main

import "html/template"

type Chapter struct {
	Title      string   `json:"title"`   // título do chapter
	Paragraphs []string `json:"story"`   // conteúdo (parágrafos)
	Options    []Option `json:"options"` // opções de próximo chapter
}

type Story map[string]Chapter

type Option struct {
	// opções de chapter são um nome do chapter (Chapter) e uma descrição (Text)
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	// handler é um tipo que implementa o método ServeHTTP
	// e Story é um map[string]Chapter
	s Story
	t *template.Template
}
