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
	fmt.Scanf("%d", &command) //the method Scanf needs to have a formatter like "%d". On Scan method you don't need to say the format because it's inferred.
	fmt.Println("The memory address of command var is:", &command)
	fmt.Println("Option choose:", command)

}
