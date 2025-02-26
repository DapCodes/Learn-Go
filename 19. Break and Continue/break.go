package main
import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i == 72 {
			break
		}

		fmt.Println("Pengulangan ke -", i)

	}
}