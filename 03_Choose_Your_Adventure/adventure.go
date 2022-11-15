package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	var story Story
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&story); err != nil {
		fmt.Printf("\nerro no decoding\n")
		panic(err)
	}
	// se não deu erro, story tem as infos decodificadas do arquivo
	// %+v é mais verboso que %v aparentemente, com mais detalhes
	fmt.Printf("decoded story1:\n\n%+v\n\n", story)
}
