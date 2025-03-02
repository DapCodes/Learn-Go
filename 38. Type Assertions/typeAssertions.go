package main

import (
	"fmt"
)

func random() interface{} {
	return 1
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

	// jika hati hati dalam melakukan konversi. jika salah maka akan menimbulkan panic
}
