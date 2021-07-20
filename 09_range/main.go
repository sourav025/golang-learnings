package main

import (
	"fmt"
	"strconv"
)

func main() {
	ids := []int{10, 20, 40, 50, 60, 100}
	// with index
	for index, id := range ids {
		fmt.Printf("%d - %d\n", index, id)
	}

	// without index: _ is used to avoid variable declaration
	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	// sum
	sm := 0
	for _, id := range ids {
		sm += id
	}
	fmt.Println("Final sum = " + strconv.Itoa(sm))
	fmt.Println("Final sum = " + fmt.Sprint(sm))

	// Loop through a map
	emails := map[string]string{"bob": "bol@email.com", "alice": "alice@email.com"}

	for name, email := range emails {
		fmt.Printf("Name=%s, Email=%s\n", name, email)
	}
	// Get keys only
	for key := range emails {
		println(key)
	}
}
