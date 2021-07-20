package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		if i!=0 {
			fmt.Print(" ")
		}
		fmt.Print(sum(i))
	}
	fmt.Println()

	f := fibonacci()
	for i := 0; i < 10; i++ {
		if i!=0 {
			fmt.Print(" ")
		}
		fmt.Print(f())
	}
	fmt.Println()
}
