package main
import "fmt"

func faktorial(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func faktorialRecursive(value int) int{
	if value == 1{
		return 1
	} else {
		return value * faktorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(faktorial(3))
	fmt.Println(faktorialRecursive(5))
}