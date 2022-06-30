package main

import (
	"fmt"
)

func main() {

	//! The Print() Function
	//? The Print() function prints its arguments with their default format.

	var i, j string = "Hello", "World"
	fmt.Print(i)
	fmt.Print(j) // Prints HelloWorld

	//?If we want to print the arguments in new lines, we need to use \n.
	fmt.Print(i, "\n") // Prints Hello
	fmt.Print(j, "\n") //        World

	//? It is also possible to only use one Print() for printing multiple variables.
	fmt.Print(i, "\n", j) // Prints Hello
	//        World

	//? If we want to add a space between string arguments, we need to use " ":
	fmt.Print(i, " ", j) // Prints Hello World

	//? Print() inserts a space between the arguments if neither are strings:
	var a, b = 1, 2
	fmt.Print(a, b) // Prints 1 2

	//! The Println() Function
	//? The Println() function is similar to Print()
	//? with the difference that a whitespace is added between the arguments, and a newline is added at the end

	fmt.Println(i, j) // Prints Hello World

	//! Printf() Function
	//?The Printf() function first formats its argument based on the given formatting verb and then prints them.
	//? Here we will use two formatting verbs:
	//? %v is used to print the value of the arguments
	//? %T is used to print the type of the arguments

	var x string = "Helllo"
	var y int = 10

	fmt.Printf("x has value: %v and type: %T", x, x)
	fmt.Printf("y has value: %v and type: %T", y, y)

	//Prints x has value: Hello and type: string
	//       y has value: 10 and type: int

}
