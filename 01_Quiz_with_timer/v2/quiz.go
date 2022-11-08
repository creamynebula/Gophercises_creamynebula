package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// go run . -csv=nomedoarquivo.csv
func main() {
	// nome da flag == csv, valor default == problems.csv, "msg qdo chamar flag -h ou --help"
	csvFilename := flag.String("csv", "problems.csv", "um arquivo csv no formato 'pergunta,resposta'")
	timeLimit := flag.Int("limit", 15, "tempo limite (em segundos) pra responder o quiz")

	// parse the command line into the defined flags (o que isso significa???)
	flag.Parse()

	// abrir o arquivo csv, e lidar com erros
	file, err := os.Open(*csvFilename)
	if err != nil {
		// exit(msg string) imprime uma msg e encerra o programa
		// Sprintf(formatacao) retorna uma string, segundo alguma formatacao
		exit(fmt.Sprintf("Deu ruim! Não conseguimos abrir %s\n", *csvFilename))
	}

	// pegar os dados do arquivo
	r := csv.NewReader(file)

	// isso aqui lê o csv e retorna um slice,
	// aonde cada elemento é um outro slice contendo o record
	// record: 2 elementos (pergunta e resposta) que tavam separados por vírgula no csv.
	lines, err := r.ReadAll()
	if err != nil {
		exit("Não deu pra parse o arquivo csv que tu mandou")
	}

	fmt.Printf("\nlines: %s\n", lines)

	// é um jogo de perguntas e respostas, então:
	problems := parseLines(lines)
	fmt.Printf("\nproblems: %s\n", problems)

	// queremos que o quiz termine quando se esgote um timer.
	// timer tem um channel C, que recebe o sinal dps da duration especificada.
	// timer é um *Timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	howManyCorrect := 0 // counter de qtas respostas foram acertadas

	// imprime o quiz, coleta as respostas, e coordena com o timer
	for i, p := range problems {
		fmt.Printf("\nProblema #%d:\n%s = ?\n", i+1, p.question)
		// channel pra sinalizar que uma resposta foi recebida
		// e transmitir essa resposta
		answerCh := make(chan string)
		// chamar uma goroutine pra pegar a resposta, pro input nao blockar
		// o programa. ou seja, nao blockar o timer!
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer // envia resposta
		}()

		select {
		case <-timer.C: // se o timer concluir
			encerraQuiz(howManyCorrect, len(problems))
			return // encerra o main, sai do programa!
		case answer := <-answerCh: // se chegou uma respostat -
			if answer == p.answer {
				fmt.Printf("Certa resposta!\n")
				howManyCorrect++
			}
		} // fim do select

	} // fim do for

	// se sair do for antes de encerrar o timer,
	// acabaram as perguntas, então encerra o quiz.
	encerraQuiz(howManyCorrect, len(problems))

} // fim main

// recebe o conteúdo do csv, retorna um slice de problems
func parseLines(lines [][]string) []problem {
	returnValue := make([]problem, len(lines))
	for i, line := range lines {
		returnValue[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return returnValue
}

func encerraQuiz(howManyCorrect int, howManyQuestions int) {
	fmt.Printf("\nVocê acertou %v de um total de %v perguntas.", howManyCorrect, howManyQuestions)
	fmt.Printf("\nEncerrando o quiz!\n")
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // código 1 == algo deu errado

}
