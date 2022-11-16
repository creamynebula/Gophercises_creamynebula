package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Story map[string]Chapter

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var defaultHandlerTemplate = `<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Choose your own adventure!</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="">
</head>

<body>

    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>

</body>

</html>`

func main() {
	// go run . -file=gopher.json
	fileName := flag.String("file", "gopher.json", "o arquivo JSON com a historinha do jogo, default == gopher.json")
	flag.Parse()
	fmt.Printf("Usando a história do arquivo %v\n\n", *fileName)

	// abrir arquivo. os.Open() recebe string!
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("\nerro abrindo o arquivo mano\n")
		panic(err)
	}

	story, err := JsonStory(file)
	// se não deu erro, story tem as infos decodificadas do arquivo
	// %+v é mais verboso que %v aparentemente, com mais detalhes
	if err == nil {
		fmt.Printf("decoded story1:\n\n%+v\n\n", story)
	}

}

func JsonStory(file io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&story); err != nil {
		fmt.Printf("\nerro no decoding\n")
		return nil, err
	}

	return story, nil
}

func NewHandler(s Story) http.Handler {

}
