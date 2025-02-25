package main
import "fmt"

func main() {
    arr := [1] int {
		20,
	}

	if length := len(arr); length > 2 {
		fmt.Println("Array memiliki lebih dari 2 index.")
	} else {
		fmt.Println("Array memiliki kurang dari 2 index.")
	}
}


