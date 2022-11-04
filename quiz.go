package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "um arquivo csv no formato 'pergunta,resposta'")
	flag.Parse()
	// go run . -csv=nomedoarquivo.csv

	file, err := os.Open(*csvFilename)
	if err != nil {
		//exit(msg string) imprime uma msg e encerra o programa
		//Sprintf(formatacao) retorna uma string, segundo alguma formatacao
		exit(fmt.Sprintf("Deu ruim! Não conseguimos abrir %s\n", *csvFilename))
	}

	r := csv.NewReader(file)

	// acho que isso aqui lê o csv e retorna um slice,
	// aonde cada elemento é um outro slice,
	// contendo os 2 elementos separados por vírgula no csv.
	lines, err := r.ReadAll()
	if err != nil {
		exit("Não deu pra parse o arquivo csv que tu mandou")
	}

	fmt.Println(lines)
}

func parseLines(lines [][]string) []problem {

}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // código 1 == algo deu errado

}
