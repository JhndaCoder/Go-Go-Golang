package main

import (
	"fmt"
)

func main() {

	//! Declare an Array
	//? In Go, there are two ways to declare an array:

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	//? This example declares two arrays (arr3 and arr4) with inferred lengths:
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	//? This example declares an array of strings:
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	//! Access Elements of an Array
	//? You can access a specific array element by referring to the index number.
	//? In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//! Change Elements of an Array
	//? You can also change the value of a specific array element by referring to the index number.

	prices[2] = 50
	fmt.Println(prices)

	//! Array Initialization
	//? If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
	//? Tip: The default value for int is 0, and the default value for string is "".

	arr5 := [5]int{}              //not initialized
	arr6 := [5]int{1, 2}          //partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	//! Initialize Only Specific Elements
	//? It is possible to initialize only specific elements in an array.

	arr8 := [5]int{1: 10, 2: 40}

	fmt.Println(arr8)

	//! Find the Length of an Array
	//? The len() function is used to find the length of an array:

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
