package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	fmt.Println(PI)

	//! Typed Constants
	const A int = 1
	fmt.Println(A)

	//! Untyped Constants
	const B = 1
	// const B = 2 		//Gives error as value of cosnt can't be changes once declared
	fmt.Println(B)

	//! Multiple Constants Declaration
	const (
		P int = 1
		Q     = 3.14
		R     = "Hi"
	)
	fmt.Println(P)
	fmt.Println(Q)
	fmt.Println(R)
}
