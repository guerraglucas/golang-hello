package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const minutesToSleep = 5
const timesToRepeat = 5

func main() {
	versao := 1.1

	nome, idade := getNameAndAge()
	printNameAndAge(nome, idade)
	printVersion(versao)
	for {
		userSelection := getUserSelection()
		doWhatUserWants(userSelection)
	}
}

func getNameAndAge() (string, int) {
	var nome string = "Lucas"
	var idade int = 20
	return nome, idade
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
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		exitProgramWithSuccess()
	default:
		exitProgramWithError()
	}
}

func startMonitoring() {

	fmt.Println("Monitorando...")
	var listOfSites []string
	listOfSites = []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br", "https://www.google.com.br"}
	for i := 0; i < timesToRepeat; i++ {
		for _, site := range listOfSites {
			fmt.Println("Testando", site)
			monitorFeedback(site)
		}
		time.Sleep(minutesToSleep * time.Minute)
	}
	fmt.Println("")
}

func monitorFeedback(site string) {
	response, _ := http.Get(site)
	switch response.StatusCode {
	case 200:
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	case 404:
		fmt.Println("Site:", site, "não foi encontrado!")
	default:
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)

	}
}

func showLogs() {
	fmt.Println("Exibindo Logs...")
}

func exitProgramWithSuccess() {
	fmt.Println("Saindo do Programa...")
	os.Exit(0)
}

func exitProgramWithError() {
	fmt.Println("Saindo do Programa...")
	os.Exit(-1)
}
