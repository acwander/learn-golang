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
		firstName: "John",
		lastName:  "Pardi",
		contactInfo: contactInfo{
			email:   "jim.pardi@gmail.com",
			zipCode: 94001,
		},
	}

	fmt.Printf("%+v", jim)
}
