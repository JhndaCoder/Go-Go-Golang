package main

import (
	"fmt"
)

func main() {

	//! Go Maps
	//? Maps are used to store data values in key:value pairs.
	//? Each element in a map is a key:value pair.
	//? A map is an unordered and changeable collection that does not allow duplicates.
	//? The length of a map is the number of its elements. You can find it using the len() function.
	//? The default value of a map is nil.
	//? Maps hold references to an underlying hash table.
	//? Go has multiple ways for creating maps.

	//! Creating Maps Using var and :=

	// Syntax
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

	var a1 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b1 := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a1\t%v\n", a1)
	fmt.Printf("b1\t%v\n", b1)

	//? Note: The order of the map elements defined in the code is different from the way that they are stored.
	//? The data are stored in a way to have efficient data retrieval from the map.

	//! Creating Maps Using Using make()Function:

	// Syntax
	// var a = make(map[KeyType]ValueType)
	// b := make(map[KeyType]ValueType)

	var a2 = make(map[string]string) // The map is empty now
	a2["brand"] = "Ford"
	a2["model"] = "Mustang"
	a2["year"] = "1964"
	// a2 is no longer empty
	b2 := make(map[string]int)
	b2["Oslo"] = 1
	b2["Bergen"] = 2
	b2["Trondheim"] = 3
	b2["Stavanger"] = 4

	fmt.Printf("a2\t%v\n", a2)
	fmt.Printf("b2\t%v\n", b2)

	//! Creating an Empty Map
	//? There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.
	// Syntax
	// var a map[KeyType]ValueType
	//? Note: The make()function is the right way to create an empty map.
	//? If you make an empty map in a different way and write to it, it will causes a runtime panic.

	//? This example shows the difference between declaring an empty map using with the make()function and without it.

	var a3 = make(map[string]string)
	var b3 map[string]string

	fmt.Println(a3 == nil)
	fmt.Println(b3 == nil)

	//! Allowed Key Types
	//? The map key can be of any data type for which the equality operator (==) is defined. These include:

	//? Booleans
	//? Numbers
	//? Strings
	//? Arrays
	//? Pointers
	//? Structs
	//? Interfaces (as long as the dynamic type supports equality)
	//? Invalid key types are:

	//? Slices
	//? Maps
	//? Functions

	//? These types are invalid because the equality operator (==) is not defined for them.

	//! Allowed Value Types
	//? The map values can be any type.

	//! Accessing Map Elements
	//? You can access map elements by:
	// value = map_name[key] var a = make(map[string]string)

	var a4 = make(map[string]string)
	a4["brand"] = "Ford"
	a4["model"] = "Mustang"
	a4["year"] = "1964"

	fmt.Printf(a4["brand"])

	//! Updating and Adding Map Elements
	//? Updating or adding an elements are done by:
	// map_name[key] = value

	var a5 = make(map[string]string)
	a5["brand"] = "Ford"
	a5["model"] = "Mustang"
	a5["year"] = "1964"

	fmt.Println(a5)

	a5["year"] = "1970" // Updating an element
	a5["color"] = "red" // Adding an element

	fmt.Println(a5)

	//! Remove Element from Map
	//? Removing elements is done using the delete() function.

	// syntax
	// delete(map_name, key)

	var a6 = make(map[string]string)
	a6["brand"] = "Ford"
	a6["model"] = "Mustang"
	a6["year"] = "1964"

	fmt.Println(a6)

	delete(a6, "year")

	fmt.Println(a6)

	//! Check For Specific Elements in a Map
	//? You can check if a certain key exists in a map using:

	// syntax
	// val, ok :=map_name[key]

	//? If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.

	var a7 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := a7["brand"] // Checking for existing key and its value
	val2, ok2 := a7["color"] // Checking for non-existing key and its value
	val3, ok3 := a7["day"]   // Checking for existing key and its value
	_, ok4 := a7["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	//! Maps Are References

	//? Maps are references to hash tables.
	//? If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)

	//! Iterating Over Maps
	//? You can use range to iterate over maps.

	a8 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a8 {
		fmt.Printf("%v : %v, ", k, v)
	}

	//! Iterate Over Maps in a Specific Order
	//? Maps are unordered data structures.
	//? If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.

	a9 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b4 []string // defining the order
	b4 = append(b4, "one", "two", "three", "four")

	for k, v := range a9 { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b4 { // loop with the defined order
	}

}
