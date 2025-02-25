package main
import "fmt"

func main() {
    nilai := 10

	switch nilai {
		case 100:
			fmt.Println("Grade A")
		case 85:
			fmt.Println("Grade b")
		case 75:
			fmt.Println("Grade c")
		default:
			fmt.Println("Nilai anda", nilai)
	}
}
