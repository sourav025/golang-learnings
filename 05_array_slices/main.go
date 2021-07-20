package main

import "fmt"

func main() {
	// init
	var fruitArray [2]string

	// Assign values seperately
	fruitArray[0] = "apple"
	fruitArray[1] = "orange"

	fmt.Println(fruitArray)

	// Assign value while creating
	arr2 := [2]string{"hello", "world"}
	fmt.Println(arr2)

	arr3 := [3]string {"apple", "orange", "noway"}
	fmt.Println(arr3)

	fmt.Println(len(arr3), arr3[1:3])
}
