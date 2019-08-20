package main

import "fmt"

// Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory.
// Arrays have some special features in Go related to how they are declared and viewed as types.

// NOTES
// Arrays are fixed length data structures that can't change.
// An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
// Arrays of different sizes are considered to be of different types.
// Memory is allocated as a contiguous block.
// Go gives you control over spacial locality.

func main() {

	// array of type string
	var cars [5]string

	cars[0] = "Lexus"
	cars[1] = "Acura"
	cars[2] = "Tesla"
	cars[3] = "Audi"
	cars[4] = "BMW"

	// iterate over the array of strings
	for i, car := range cars {
		fmt.Println(i, car)
	}

	// declaring and initializing an array of type int
	numbers := [4]int{10, 20, 30, 80}

	// Iterate over the lenght of the array of numbers
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// Declare an array of 5 integers that is initialized
	// to its zero value.
	var five [5]int

	// Declare an array of 5 integers that is initialized
	// with some values.
	four := [5]int{10, 20, 30, 40}

	// Assign one array to the other
	five = four

	fmt.Println(four)
	fmt.Println(five)

	// NOTE:
	// If array four was of length 4 then we would have got an error
	// since both arrays would have been of a different size
	// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment

	// Sample program to show how memory for an array is contiguous.

	// Declare an array of 5 strings initialized with values.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Iterate over the array displaying the value and
	// address of each element.
	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
	}

}
