package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}
type Animal struct {
	Name   string
	Weight string
}

func (person Person) getName() string {
	return person.Name
}
func (animal Animal) getName() string {
	return animal.Weight
}

func main() {
	person := Person{Name: "Daffa"}
	sayHello(person)

	animal := Animal{Weight: "20"}
	sayHello(animal)
}
