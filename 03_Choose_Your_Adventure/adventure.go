package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

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

func init() {
	// renderiza um template default no começo do programa, acho
	tpl = template.Must(template.New("Nome do template").Parse(defaultHandlerTemplate))
}

var tpl *template.Template

// o layout da página do jogo
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
        {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>

</body>

</html>`

func main() {
	// go run . -file=gopher.json -port=3000
	port := flag.Int("port", 3000, "a porta pra iniciar a app web do jogo de \"escolha sua aventura\"")
	fileName := flag.String("file", "gopher.json", "o arquivo JSON com a historinha do jogo, default == gopher.json")
	flag.Parse()
	fmt.Printf("Usando a história do arquivo %v\n\n", *fileName)

	// abrir arquivo JSON.
	// os.Open() recebe string (nome do arquivo)!
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("\nerro abrindo o arquivo mano\n")
		panic(err)
	}

	// extrair story do arquivo, através de criar um novo decoder,
	// e chamar decoder.NewDecode(&var)
	story, err := JsonStory(file)

	// se não deu erro, story tem as infos decodificadas do arquivo
	// %+v é mais verboso que %v aparentemente, com mais detalhes
	if err == nil {
		fmt.Printf("decoded story1:\n\n%+v\n\n", story)
	}

	handler := NewHandler(story)
	fmt.Printf("Começando o server na porta: %v\n\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
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

// retornando uma interface que implementa o método ServeHTTP
// porque o type Handler é definido assim.
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	// handler é um tipo que implementa o método ServeHTTP
	// e Story é um map[string]Chapter
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//remover espaços, e, se path vazio, voltar pro começo (intro)
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	// "/a-certain-path" -> "a-certain-path"
	path = path[1:]

	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
