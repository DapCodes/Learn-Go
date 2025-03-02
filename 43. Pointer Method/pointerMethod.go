package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	daffa := Man{"Daffa"}
	daffa.Married()

	fmt.Println(daffa.Name)
}
