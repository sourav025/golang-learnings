package main

import "fmt"

func main() {
	// Using var
	var name = "Sourav"
	var age int32 = 30
	var size float32 = 28.8
	const isCool = false
	fmt.Println(name, age, isCool, size)

	// Get typeof variable using fmt
	// https://pkg.go.dev/fmt
	fmt.Printf("%T - %T\n", isCool, size)

	// Shortcut method to create variable only allowed in function
	anyvar := "anyString"
	fmt.Println(anyvar)

}
