package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func familyName1(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func familyName2(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction1(x int, y int) int {
	return x + y
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction3(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunction4(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func myFunction5(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func myFunction6(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {

	//! Go Functions
	//? A function is a block of statements that can be used repeatedly in a program.
	//? A function will not execute automatically when a page loads.
	//? A function will be executed by a call to the function.

	//! Create a Function
	//? To create (often referred to as declare) a function, do the following:
	//? 	1. Use the func keyword.
	//? 	2. Specify a name for the function, followed by parentheses ().
	//? 	3. Finally, add code that defines what the function should do, inside curly braces {}.

	//	sytnax
	//	func FunctionName() {
	//	 code to be executed
	//	}

	//! Call a Function
	//? Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.
	//? In the example below, we create a function named "myMessage()".
	//? The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function.
	//? The function outputs "I just got executed!".
	//? To call the function, just write its name followed by two parentheses ():

	myMessage() // call the function

	//? A function can be called multiple times.
	myMessage()
	myMessage()
	myMessage()

	//! Parameters and Arguments
	//? Information can be passed to functions as a parameter. Parameters act as variables inside the function.
	//? Parameters and their types are specified after the function name, inside the parentheses.
	//? You can add as many parameters as you want, just separate them with a comma:

	// syntax
	// func FunctionName(param1 type, param2 type, param3 type) {
	// code to be executed
	// }

	//! Function With Parameter Example
	//? The following example has a function with one parameter (fname) of type string.
	//? When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name:

	familyName1("Liam")
	familyName1("Jenny")
	familyName1("Anja")
	//? Note: When a parameter is passed to the function, it is called an argument.
	//? So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.

	//! Multiple Parameters
	//? Inside the function, you can add as many parameters as you want:

	familyName2("Liam", 3)
	familyName2("Jenny", 14)
	familyName2("Anja", 30)

	//! Return Values
	//? If you want the function to return a value, you need to define the data type of the return value
	//? (such as int, string, etc), and also use the return keyword inside the function:

	// syntax
	// func FunctionName(param1 type, param2 type) type {
	// code to be executed
	// return output
	// }

	//! Function Return Example
	//? Here, myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int):

	fmt.Println(myFunction1(1, 2))

	//! Named Return Values
	//? In Go, you can name the return values of a function.

	//? Here, we name the return value as result (of type int),
	//? and return the value with a naked return (means that we use the return statement without specifying the variable name):

	fmt.Println(myFunction2(1, 2))

	//? The example above can also be written like this. Here, the return statement specifies the variable name:
	fmt.Println(myFunction3(1, 2))

	//! Store the Return Value in a Variable
	//? Here, we store the return value in a variable called total:
	total := myFunction2(1, 2)
	fmt.Println(total)

	//! Multiple Return Values
	//? Go functions can also return multiple values.
	//? Here, myFunction() returns one integer (result) and one string (txt1):

	fmt.Println(myFunction4(5, "Hello"))

	//? Here, we store the two return values into two variables (a and b):
	a, b := myFunction5(5, "Hello")
	fmt.Println(a, b)

	//? If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
	//? Here, we want to omit the first returned value (result - which is stored in variable a):

	_, c := myFunction6(5, "Hello")
	fmt.Println(c)

	//? Here, we want to omit the second returned value (txt1 - which is stored in variable b):
	d, _ := myFunction6(5, "Hello")
	fmt.Println(d)

}
