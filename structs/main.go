package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Shreyansh", lastName: "jain"}

	pointerToPerson := &alex

	pointerToPerson.updateFirstName("Hello")

	fmt.Print(alex)
}

func (pointerToPerson *person) updateFirstName(newName string) {
	(*pointerToPerson).firstName = newName
}
