package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Awnser  int
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("não é permitido caractere diferente de número")
	}

	return i, nil
}

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(a)! ao quiz")
	fmt.Printf("Escreva o seu nome: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string!")
	}

	g.Name = name
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quiz-go.csv")
	if err != nil {
		panic("erro ao ler arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler arquivo!")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Awnser:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) ShowResults() {
	fmt.Println("\n--------- RESULTADO FINAL ---------")
	fmt.Printf("Jogador: %s", g.Name)
	fmt.Printf("Pontuação: %d/%d\n", g.Points, len(g.Questions))

	percentage := float64(g.Points) / float64(len(g.Questions)) * 100
	fmt.Printf("Percentual de acerto: %.1f%%\n", percentage)

	switch {
	case percentage >= 90:
		fmt.Println("🏆 Excelente! Você é um expert em Go!")
	case percentage >= 70:
		fmt.Println("👍 Muito bom! Você tem um bom conhecimento de Go!")
	case percentage >= 50:
		fmt.Println("📚 Razoável! Continue estudando Go!")
	default:
		fmt.Println("💪 Precisa estudar mais! Não desista!")
	}
}

func (g *GameState) Run() {
	reader := bufio.NewReader(os.Stdin)

	for index, question := range g.Questions {
		fmt.Printf("\n\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Print("\nSua resposta [1-4]: ")
		var answer int
		var err error

		for {
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])
			if err != nil || answer < 1 || answer > 4 {
				fmt.Println("Resposta inválida! Digite um número entre 1 e 4.")
				fmt.Print("Sua resposta [1-4]: ")
				continue
			}
			break
		}

		if answer == question.Awnser {
			fmt.Println("✓ Correto!")
			g.Points++
		} else {
			fmt.Println("✗ Incorreto!")
		}
	}
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()

	if len(game.Questions) == 0 {
		fmt.Println("Nenhuma pergunta foi carregada!")
		return
	}

	fmt.Printf("Carregadas %d perguntas. Vamos começar!\n\n", len(game.Questions))
	game.Run()

	game.ShowResults()
}
