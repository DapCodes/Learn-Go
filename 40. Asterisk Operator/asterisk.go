package main

import "fmt"

type GetName struct {
	firstName, lastName string
}

func main() {
	name := GetName{"Daffa", "Ramadhan"}
	name2 := &name

	name2.lastName = "Maulana"

	fmt.Println(name)
	fmt.Println(name2)

	*name2 = GetName{"Rio", "Oktora"}
	fmt.Println(name)
	fmt.Println(name2)

}
