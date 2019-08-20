package main

/*
variables have the ability to read from and write to memory.
In Go, access to memory is type safe.
This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.
*/

// NOTES:
// The purpose of all programs and all parts of those programs is to transform data from one form to the other.
// Code primarily allocates, reads and writes to memory.
// Understanding type is crucial to writing good code and understanding code.
// If you don't understand the data, you don't understand the problem.
// You understand the problem better by understanding the data.
// When variables are being declared to their zero value, use the keyword var.
// When variables are being declared and initialized, use the short variable declaration operator.

import "fmt"

func main() {

	// Declare variables that are set to their zero value.
	var aa1 float32

	// Display the value of those variables.
	fmt.Printf("value of aa1 is %v \n ", aa1)

	// Declare variables and initialize.

	// Using the short variable declaration operator.
	myPi := 3.14

	// Display the value of those variables.
	fmt.Printf("value of myPi is %v \n", myPi)
	// Perform a type conversion.
	pi := float32(myPi)

	// Display the value of that variable.
	fmt.Printf("%T %v \n", pi, pi)

}
