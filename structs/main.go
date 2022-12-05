package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	rodrigo := person{firstName: "Rodrigo", lastName: "Alencar"}

	fmt.Println(rodrigo)
	fmt.Printf("%+v", rodrigo)
}
