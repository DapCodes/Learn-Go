package main
import "fmt"

func main() {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Daffa"
	newSlice[1] = "Ramadhan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Maulana")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Rio"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Menggunakan arr yang sama. kalo append membuat arr berbeda (baru)
}


