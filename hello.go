package main

import (
	"fmt"
	"os"
)

func main() {
	nome := "Lucas"
	idade := 20
	versao := 1.1
	printNameAndAge(nome, idade)
	printVersion(versao)

	userSelection := getUserSelection()

	doWhatUserWants(userSelection)
}

func printNameAndAge(nome string, idade int) {
	fmt.Println("Olá Sr.", nome, "sua idade é", idade)
}

func printVersion(versao float64) {
	fmt.Println("Este programa está na versão", versao)
}

func getUserSelection() int {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	var comando int
	var ptr *int = &comando
	fmt.Scan(ptr)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func doWhatUserWants(userSelection int) {
	switch userSelection {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}
