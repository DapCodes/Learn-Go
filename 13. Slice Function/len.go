package main
import "fmt"

func main() {
    names := [...]string {
		"Daffa", 
		"Ramadhan",
		"Maulana",
		"Rio", 
		"Dhea", 
		"Febrianti",
	}

	slice1 := names[3:6]
	fmt.Println(len(slice1))

	slice2 := names[:4]
	fmt.Println(cap(slice2))

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)
	
}


