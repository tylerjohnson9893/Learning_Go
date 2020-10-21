package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "James",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "jim@email.com",
			zip:   12345,
		},
	}
	jim.updateName("Jim")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
