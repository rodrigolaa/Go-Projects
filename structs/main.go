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

	//rodrigo := person{firstName: "Rodrigo", lastName: "Alencar"}

	jim := person{firstName: "Jim", lastName: "Nerd",
		contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 94000}}

	//jimPointer := &jim

	jim.updateName("Rodrigo")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {

	fmt.Printf("%+v", p)

}
