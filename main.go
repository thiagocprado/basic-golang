package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 3
const deley = 5

func main() {
	showIntroduction()

	for {
		showMenu()

		command := readCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			readLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Thiago"
	version := 1.0

	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão:", version)
	fmt.Println("")
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	fmt.Println("")
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("O comando escolhido foi:", readCommand)
	fmt.Println("")

	return readCommand
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	fmt.Println("")

	sites := readSitesFromFile()

	for i := 0; i < monitorings; i++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, "|", site)
			testSite(site)
		}

		time.Sleep(deley * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		saveLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		saveLog(site, false)
	}
}

func readSitesFromFile() []string {
	sites := []string{}
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func saveLog(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readLogs() {
	fmt.Println("Exibindo Logs...")
	fmt.Println("")

	file, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
