package main

import (
	"fmt" //format
	"os"
)

func main() {

	showsIntroduction()
	for { // while (true)
		showsMenu()

		command := readsCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
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
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"
	sites[3] = "https://idwall.co"
	fmt.Println(sites)

	// resp, _ := http.Get(sites)

	// if resp.StatusCode == 200 {
	// 	fmt.Println("The website:", site, "was reloaded successfully")
	// } else {
	// 	fmt.Println("The website:", site, "is crashed. Status code:", resp.StatusCode)
	// }
}

func showLogs() {
	fmt.Println("Showing logs...")
}
