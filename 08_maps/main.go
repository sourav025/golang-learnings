package main

import "fmt"

func main() {
	// define map
	hash := make(map[string]string)
	hash["name"] = "hello"
	hash["last"] = "world"

	// Length of entries
	fmt.Println(len(hash))

	// Print the whole hashMap
	fmt.Println(hash)

	// get entry
	fmt.Println(hash["name"])

	// exists
	_, ok := hash["foo"]
	if ok {
		fmt.Println("Yes foo exists in hash")
	} else {
		fmt.Println("Does not exist")
	}

	for k, v := range hash {
		fmt.Println(k, v)
	}

}
