package main
import "fmt"

func main() {
    name := "buddy"

	if name == "Daffa" {
		fmt.Println("Hallo Daffa")
	} else if name == "Rio" {
		fmt.Println("Hallo Rio")
	} else {
		fmt.Println("Hallo", name)
	}
}


