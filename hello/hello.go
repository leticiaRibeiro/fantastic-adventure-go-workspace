package main

import (
	"fmt" //format
	"net/http"
	"os"
)

func main() {

	showsIntroduction()
	showsMenu()

	command := readsCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		fmt.Println("Exit...")
		os.Exit(0)
	default:
		fmt.Println("Dunno wat u said")
		os.Exit(-1)
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
	site := "https://www.alura.com.br"
	resp, err := http.Get(site)
}

func showLogs() {
	fmt.Println("Showing logs...")
}
