package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Negara = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
