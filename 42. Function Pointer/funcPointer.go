package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func ubahNegara(address *Address) {
	address.Negara = "Indonesia"
}

func main() {
	alamat := Address{"Bandung", "Jawa Barat", ""}
	ubahNegara(&alamat)

	fmt.Println(alamat)
}
