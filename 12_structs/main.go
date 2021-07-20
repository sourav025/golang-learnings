package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (person Person) greet() {
	fmt.Println("Greetings " + person.firstName + person.lastName)
}

func (person *Person) add1Yr() {
	// adding age
	person.age++
}

func (person *Person) getMarried(finalName string) {
	if person.gender == "f" {
		person.lastName = finalName
	} else {
		fmt.Println("No need to change last cause he is male!!!")
	}
}

func main() {
	// Init person
	person1 := Person{firstName: "Sourav", lastName: "Sarker", city: "SG", gender: "m", age: 31}
	// alternative
	// person1 := Person{"Sourav", "Sarker", "SG", "m", 31}
	fmt.Println(person1)

	fmt.Println(person1.firstName + " " + person1.lastName)
	person1.greet()
	person1.add1Yr()
	fmt.Println(person1)
	person1.getMarried("Noway")
	person2 := Person{"Wife", "BeforeMarried", "SG", "f", 25}
	fmt.Println(person2)
	person2.getMarried(person1.lastName)
	fmt.Println(person2)

}
