package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var Daffa Customer
	Daffa.Name = "Daffa Ramadhan"
	Daffa.Address = "Bandung"
	Daffa.Age = 17
	fmt.Println(Daffa)

	Rio := Customer{
		Name:    "Rio Oktora",
		Address: "Bandung",
		Age:     7,
	}
	fmt.Println(Rio)

	Dhea := Customer{"Dhea Febrianti", "Bandung", 17}
	fmt.Println(Dhea)

}
