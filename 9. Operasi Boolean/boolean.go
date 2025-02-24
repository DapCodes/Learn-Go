package main
import "fmt"

func main() {
    var a = 5
	var b = 7
	var c = 7
	var d = 5

	var result1 = a == d && b == c
	fmt.Println(result1)

	var result2 = a == c || b == c
	fmt.Println(result2)

	var result3 = a == d && b == c
	fmt.Println(!result3)
}


