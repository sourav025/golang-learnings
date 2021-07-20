package main

import "fmt"

func fib(a *int, b *int) {
	tmp := *a + *b
	*a = *b
	*b = tmp
}

func printFibonacci(limit int) {
	p, q := 0, 1
	fmt.Print(p, q)
	for i := 0; i <= limit; i++ {
		fib(&p, &q)
		fmt.Printf(" %d", q)
	}
	fmt.Println()
}

func main() {
	// Pointer example
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)

	// Use * to read pointer value
	fmt.Println(*b)

	*b++
	fmt.Println(a)

	// Fibonacci with pointer
	printFibonacci(15)

}
