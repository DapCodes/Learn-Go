package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "myname is", customer.Name)
}

func main() {
	daffa := Customer{Name: "Daffa"}
	daffa.sayHello("Rio")
}
