package main

import (
	"fmt" //format
	"net/http"
	"os"
	"reflect"
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
			showNames()
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
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br", "https://idwall.co"}
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

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was reloaded successfully. Status code:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "is crashed!!!! Status code:", resp.StatusCode)
	}
}
