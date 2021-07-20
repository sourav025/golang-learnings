package main

import "fmt"

/*
Loop example
*/

func getUsingFor() int {
	sm := 0
	// for i := 0; i < 10; i++ {
	// 	sm += i
	// }
	i := 0
	for i < 10 {
		sm += i
		i++
	}
	return sm
}

// print fizbuzz
func fizzbuzz(n int) {
	for i := 0; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d FIZZBUZZ\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d FIZZ\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d BUZZ\n", i)
		}
	}
}

func main() {
	fmt.Printf("Sum is = %d\n", getUsingFor())
	fizzbuzz(15)
}
