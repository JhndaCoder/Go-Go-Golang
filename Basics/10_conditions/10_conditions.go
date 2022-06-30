package main

import (
	"fmt"
)

func main() {

	//! GO conditions
	//? A condition can be either true or false.

	//? Go supports the usual comparison operators from mathematics:

	//? Less than <
	//? Less than or equal <=
	//? Greater than >
	//? Greater than or equal >=
	//? Equal to ==
	//? Not equal to !=
	//? Additionally, Go supports the usual logical operators:

	//? Logical AND &&
	//? Logical OR ||
	//? Logical NOT !

	//! The if Statement
	//? Use the if statement to specify a block of Go code to be executed if a condition is true.

	//In the example below, we test two values to find out if 20 is greater than 18. If the condition is true, print some text:
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	// We can also test variables:
	x := 20
	y := 18
	if x > y {
		fmt.Println("x is greater than y")
	}

	//! The else Statement
	//? Use the else statement to specify a block of code to be executed if the condition is false.

	//In this example, time (20) is greater than 18, so the if condition is false.
	// Because of this, we move on to the else condition and print to the screen "Good evening".
	//If the time was less than 18, the program would print "Good day":
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	// In this example, the temperature is 14 so the condition for if is false so the code line inside the else statement is executed:

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there") //Having the else brackets in a different line will raise an error:
	}

	//! The else if Statement
	//? Use the else if statement to specify a new condition if the first condition is false.

	time1 := 22
	if time1 < 10 {
		fmt.Println("Good morning.")
	} else if time1 < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	//? Another example for the use of else if.
	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}

	//? If conditions1 and condition2 are both correct, only the code for condition1 are executed:
	n := 30
	if n >= 10 {
		fmt.Println("n is larger than or equal to 10.")
	} else if n > 20 {
		fmt.Println("n is larger than 20.")
	} else {
		fmt.Println("n is less than 10.")
	}

	//! The Nested if Statement
	//?  You can have if statements inside if statements, this is called a nested if.

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

}
