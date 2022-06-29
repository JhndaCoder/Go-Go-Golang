package main

import (
	"fmt"
)

func main() {

	//! Variable Declaration With Initial Value
	var student1 string = "John"
	var student2 string = "Jane"
	x := 2 // := can be used inside a function only

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	//! Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool

	fmt.Println(a) // prints empty string as nothing was assigned
	fmt.Println(b) // prints 0 as nothing was assigned
	fmt.Println(c) // prints false as nothing was assigned

	//! Value Assignment After Declaration
	var student3 string
	student3 = "Uddu"
	fmt.Println(student3)

	//! Go Multiple Variable Declaration
	var p, q, r, s int = 1, 3, 5, 7

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

	//! If the type keyword is not specified, you can declare different types of variables in the same line:

	var a1, a2 = 6, "Hello"
	a3, a4 := 7, "World"

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	//! Go Variable Declaration in a Block
	var (
		b1 int
		b2 int    = 1
		b3 string = "Hello"
	)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
}
