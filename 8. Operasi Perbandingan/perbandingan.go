package main
import "fmt"

func main() {
    var option1 = 5
	var option2 = 6

	var result1 bool = option1 == option2
	fmt.Println(option1, "==", option2 ,result1)

	var result2 bool = option1 != option2
	fmt.Println(option1, "!=", option2 ,result2)

	var result3 bool = option1 > option2
	fmt.Println(option1, ">", option2 ,result3)


	var result4 bool = option1 < option2
	fmt.Println(option1, "<", option2 ,result4)


	var result5 bool = option1 >= option2
	fmt.Println(option1, ">=", option2 ,result5)

	var result6 bool = option1 <= option2
	fmt.Println(option1, "<=", option2 ,result6)
}
