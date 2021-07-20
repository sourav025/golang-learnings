package main

import "fmt"

func greet(name string) string {
	return "Hello "+name+"!!!"
}

func getSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func main() {
	name := "Sourav"
	fmt.Println(greet(name))
	fmt.Println(getSum(10))
}
