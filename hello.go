package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const minutesToSleep = 1
const timesToRepeat = 2

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

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	fmt.Println(sites)
	return sites
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
	listOfSites = readSitesFromFile()
	for i := 0; i < timesToRepeat; i++ {
		for _, site := range listOfSites {
			fmt.Println("Testando", site)
			monitorFeedback(site)
		}
		if i == timesToRepeat-1 {
			break
		}
		time.Sleep(minutesToSleep * time.Minute)
	}
	fmt.Println("")
}

func monitorFeedback(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	switch response.StatusCode {
	case 200:
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLogs(site, true)
	case 404:
		fmt.Println("Site:", site, "não foi encontrado!")
		registerLogs(site, false)
	default:
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		registerLogs(site, false)
	}
}

func showLogs() {
	fmt.Println("Exibindo Logs...")
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		fmt.Println(line)

	}
}

func exitProgramWithSuccess() {
	fmt.Println("Saindo do Programa...")
	os.Exit(0)
}

func exitProgramWithError() {
	fmt.Println("Saindo do Programa...")
	os.Exit(-1)
}

func registerLogs(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	if status {
		log.Println("Site:", site, "foi carregado com sucesso! Às", time.Now().Format("02/01/2006 15:04:05"))
	} else {
		log.Println("Site:", site, "está com problemas. Às", time.Now().Format("02/01/2006 15:04:05"))
	}

}
