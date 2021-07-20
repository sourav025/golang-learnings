package main

import "fmt"

func main() {
	x, y := 10, 5
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// string match
	color := "red"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is " + color)
	}

	// switch-case
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not RED or Blue")

	}
}
