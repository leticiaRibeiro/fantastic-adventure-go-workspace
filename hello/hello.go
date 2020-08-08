package main

import (
	"fmt"     //format
	"reflect" //to know the types of variables
)

func main() {
	var name = "Lele"         //not necessary to say it's a string (because it's inferred)
	var age = 25              //so does here
	var version float32 = 1.1 //so does here
	fmt.Println("Hey", name, "your age is", age)
	fmt.Println("This project is on version", version)

	fmt.Println("The type of the variable name is", reflect.TypeOf(name))
	fmt.Println("The type of the variable age is", reflect.TypeOf(age))
	fmt.Println("The type of the variable version is", reflect.TypeOf(version))
}
