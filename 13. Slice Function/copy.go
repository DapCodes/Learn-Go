package main
import "fmt"

func main() {
    days := [...]string {
		"Senin", 
		"Selasa",
		"Rabu",
		"Kamis", 
		"Jumat", 
		"Sabtu",
		"Minggu",
		"Hallo",
	}

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
