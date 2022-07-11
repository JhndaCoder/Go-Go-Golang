package main

import (
	"fmt"
)

func main() {

	//! Go Slices
	//? Slices are similar to arrays, but are more powerful and flexible.
	//? Like arrays, slices are also used to store multiple values of the same type in a single variable.
	//? However, unlike arrays, the length of a slice can grow and shrink as you see fit.
	//? In Go, there are several ways to create a slice:

	// mySlice := []int{1, 2, 3}
	//? The code above declares a slice of integers of length 3 and also the capacity of 3.

	//? In Go, there are two functions that can be used to return the length and capacity of a slice:

	//? len() function - returns the length of the slice (the number of elements in the slice)
	//? cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	//! Create a Slice From an Array

	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end]          // A slice made from the array

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	//! Create a Slice With The make() Function

	// slice_name := make([]type, length, capacity)

	mySlice3 := make([]int, 5, 10)
	fmt.Printf("mySlice3 = %v\n", mySlice3)
	fmt.Printf("length = %d\n", len(mySlice3))
	fmt.Printf("capacity = %d\n", cap(mySlice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	//! Access Elements of a Slice
	//?  You can access a specific slice element by referring to the index number.
	//? In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//! Change Elements of a Slice
	//? You can also change a specific slice element by referring to the index number.

	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//! Append Elements To a Slice
	//? You can append elements to the end of a slice using the append()function:
	// slice_name = append(slice_name, element1, element2, ...)

	mySlice5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("mySlice5 = %v\n", mySlice5)
	fmt.Printf("length = %d\n", len(mySlice5))
	fmt.Printf("capacity = %d\n", cap(mySlice5))

	mySlice5 = append(mySlice5, 20, 21)
	fmt.Printf("mySlice5 = %v\n", mySlice5)
	fmt.Printf("length = %d\n", len(mySlice5))
	fmt.Printf("capacity = %d\n", cap(mySlice5))

	//! Append One Slice To Another Slice
	//? To append all the elements of one slice to another slice, use the append()function:
	//  slice3 = append(slice1, slice2...)

	myslice6 := []int{1, 2, 3}
	myslice7 := []int{4, 5, 6}
	myslice8 := append(myslice6, myslice7...)

	fmt.Printf("myslice8=%v\n", myslice8)
	fmt.Printf("length=%d\n", len(myslice8))
	fmt.Printf("capacity=%d\n", cap(myslice8))

	//! Change The Length of a Slice
	//? Unlike arrays, it is possible to change the length of a slice.

	arr2 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice9 := arr2[1:5]                 // Slice array
	fmt.Printf("myslice9 = %v\n", myslice9)
	fmt.Printf("length = %d\n", len(myslice9))
	fmt.Printf("capacity = %d\n", cap(myslice9))

	myslice9 = arr2[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice9 = %v\n", myslice9)
	fmt.Printf("length = %d\n", len(myslice9))
	fmt.Printf("capacity = %d\n", cap(myslice9))

	myslice9 = append(myslice9, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice9 = %v\n", myslice9)
	fmt.Printf("length = %d\n", len(myslice9))
	fmt.Printf("capacity = %d\n", cap(myslice9))

	//! Memory Efficiency
	//? When using slices, Go loads all the underlying elements into the memory.
	//? If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
	//? The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.

	// copy(dest, src)
	// The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
