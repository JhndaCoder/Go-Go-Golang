package main

import (
	"fmt"
)

func main() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	//! Boolean Data Type
	//? A boolean data type is declared with the bool keyword and can only take the values true or false.
	//? The default value of a boolean data type is false.

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	//! Integer Data Type
	//? Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.
	//? The integer data type has two categories:
	//? Signed integers - can store both positive and negative values
	//? Unsigned integers - can only store non-negative values

	var x1 int = 500   // signed int
	var y1 int = -4500 // signed int
	fmt.Printf("Type: %T, value: %v", x1, x1)
	fmt.Printf("Type: %T, value: %v", y1, y1)

	var x2 uint = 500
	var y2 uint = 4500
	fmt.Printf("Type: %T, value: %v", x2, x2) // unsigned int (denoted by uint)
	fmt.Printf("Type: %T, value: %v", y2, y2) // unsigned int (denoted by uint)

	//! Float Data Type
	//? The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.
	//? The float data type has two keywords:

	//?	Type	  Size		 Range
	//? float32	  32 bits	-3.4e+38 to 3.4e+38.
	//? float64	  64 bits	-1.7e+308 to +1.7e+308.

	var x3 float32 = 123.78
	var y3 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x3, x3)
	fmt.Printf("Type: %T, value: %v", y3, y3)

	var x4 float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x4, x4)

	//! String Data Type
	//? The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %s\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
