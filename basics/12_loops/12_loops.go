package main

import "fmt"

func main() {

	//! Go For Loops
	//? The for loop loops through a block of code a specified number of times.
	//?The for loop is the only loop available in Go.

	//? Loops are handy if you want to run the same code over and over again, each time with a different value.
	//? Each execution of a loop is called an iteration.
	//? The for loop can take up to three statements:

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	//! The continue Statement
	//? The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//! The break Statement
	//? The break statement is used to break/terminate the loop execution.

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//! Nested Loops
	//? It is possible to place a loop inside another loop.
	//? Here, the "inner loop" will be executed one time for each iteration of the "outer loop":

	adj := [2]string{"big", "tasty"}
	fruits1 := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits1); j++ {
			fmt.Println(adj[i], fruits1[j])
		}
	}

	//! The Range Keyword
	//? The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

	//? The range keyword is used like this:
	// for index, value := array|slice|map {
	//  code to be executed for each iteration
	// }

	//? This example uses range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):
	fruits2 := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits2 {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	//? Here, we want to omit the indexes (idx stores the index, val stores the value):
	fruits3 := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits3 {
		fmt.Printf("%v\n", val)
	}

	//? Here, we want to omit the values (idx stores the index, val stores the value):
	fruits4 := [3]string{"apple", "orange", "banana"}

	for idx, _ := range fruits4 {
		fmt.Printf("%v\n", idx)
	}
}
