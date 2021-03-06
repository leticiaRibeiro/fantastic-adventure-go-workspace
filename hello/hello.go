package main

import (
	"bufio"
	"fmt" //format
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	showsIntroduction()

	for { // while (true)
		showsMenu()

		command := readsCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Exit... bye :)")
			os.Exit(0)
		default:
			fmt.Println("Dunno wat u said")
			os.Exit(-1)
		}
	}
}

func showsIntroduction() {
	name := "Lele"
	var version float32 = 1.1
	fmt.Println("Hey", name)
	fmt.Println("This project is on version", version)
}

func readsCommand() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Option choose:", comandoLido)

	return comandoLido
}

func showsMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	// sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br", "https://idwall.co"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
	}
	time.Sleep(delay * time.Second)
}

func showNames() {
	nomes := []string{"Leticia", "Ribeiro", "Bezerra"}
	nomes = append(nomes, "Campos")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("The slice has", len(nomes)) //length
}

func testeDeSlices() {
	teste := []string{}
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	teste = append(teste, "eae")
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	teste = append(teste, "eae")
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	teste = append(teste, "eae")
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	nomes := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
	nomes = append(nomes, "Novo")
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
	nomes = append(nomes, "Novo2", "Novo3")
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
	nomes = append(nomes, "Novo4")
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was reloaded successfully. Status code:", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "is crashed!!!! Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	fmt.Println(sites)
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Exception:", err)
	}

	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format(time.Kitchen) + "---" + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func printLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Exception:", err)
	}

	fmt.Println(string(arquivo))
}
