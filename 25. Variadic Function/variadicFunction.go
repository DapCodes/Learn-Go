package main
import "fmt"

func sumAll(numbers ...int) int{
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(1,1,1,1,1)
	fmt.Println(total)

	numbers := []int {10, 20, 30}
	fmt.Println(sumAll(numbers...))

}