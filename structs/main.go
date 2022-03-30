package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Pardi",
		contactInfo: contactInfo{
			email:   "jim.pardi@gmail.com",
			zipCode: 94001,
		},
	}

	// NOTES:
	// example data: [0001][person{firstName:"Jim"}]
	// 0001 is the memory address
	// person{firstName:"Jim"} is the value
	// turn address into value with *address
	// turn value into address with &value

	// assign jimPointer to the memory address of jim
	jimPointer := &jim

	jimPointer.updateName("Jimmy")
	jim.print()

}

// Receiver functions

// *person - This is a type description, it means we are working with a pointer to a person
// *pointerToPerson = This is an operator, it means we want to manupulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
