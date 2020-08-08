package main

import (
	"fmt" //format
)

func main() {
	name := "Lele"
	var version float32 = 1.1
	fmt.Println("Hey", name)
	fmt.Println("This project is on version", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)
	fmt.Println("Option choose:", command)

	if command == 1 {
		fmt.Println("Starting Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing logs...")
	} else if command == 0 {
		fmt.Println("Exit...")
	} else {
		fmt.Println("Dunno wat u said")
	}
}
