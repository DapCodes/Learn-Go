package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1

	*address2 = Address{"Tarakan", "Kalimantan", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
