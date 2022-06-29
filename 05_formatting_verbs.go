package main

import (
	"fmt"
)

func main() {

	//! General Formatting Verbs
	var i = 15.5
	var txt = "Hello World"

	fmt.Printf("%v\n", i)   // Prints the value in the default format
	fmt.Printf("%#v\n", i)  // Prints the value in Go-syntax format
	fmt.Printf("%v%%\n", i) // %% Prints the % sign
	fmt.Printf("%T\n", i)   // Prints the type of the value
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	//! Integer Formatting Verbs
	var j = 15

	fmt.Printf("%b\n", j)   // Base 2
	fmt.Printf("%d\n", j)   // Base 10
	fmt.Printf("%+d\n", j)  // Base 10 and always show sign
	fmt.Printf("%o\n", j)   // Base 8
	fmt.Printf("%O\n", j)   // Base 8, with leading 0o
	fmt.Printf("%x\n", j)   // Base 16, lowercase
	fmt.Printf("%X\n", j)   // Base 16, uppercase
	fmt.Printf("%#x\n", j)  // Base 16, with leading 0x
	fmt.Printf("%4d\n", j)  // Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", j) // Pad with spaces (width 4, left justified)
	fmt.Printf("%04d\n", j) // Pad with zeroes (width 4)

	//! String Formatting Verbs

	fmt.Printf("%s\n", txt)   // Prints the value as plain string
	fmt.Printf("%q\n", txt)   // Prints the value as a double-quoted string
	fmt.Printf("%8s\n", txt)  // Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", txt) // Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", txt)   // Prints the value as hex dump of byte values
	fmt.Printf("% x\n", txt)  // Prints the value as hex dump with spaces

	//! Boolean Formatting Verbs
	var a = true
	var b = false

	fmt.Printf("%t\n", a) //Value of the boolean operator in true or false format (same as using %v)
	fmt.Printf("%t\n", b)
	// fmt.Printf(a) // gives Error

	//! Float Formatting Verbs

	var x = 3.141

	fmt.Printf("%e\n", x)    // Scientific notation with 'e' as exponent
	fmt.Printf("%f\n", x)    // Decimal point, no exponent
	fmt.Printf("%.2f\n", x)  // Default width, precision 2
	fmt.Printf("%6.2f\n", x) // Width 6, precision 2
	fmt.Printf("%g\n", x)    // Exponent as needed, only necessary digits
}
