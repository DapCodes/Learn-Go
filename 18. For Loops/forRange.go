package main
import "fmt"

func main() {

	names := []string {
		"Daffa",
		"Ramadhan",
		"Maulana",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// cara menggunakan for range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

}
