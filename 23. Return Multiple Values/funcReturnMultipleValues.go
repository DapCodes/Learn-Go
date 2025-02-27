// funcReturnMultipleValues.go
package main
import "fmt"

func getFullName() (string, string) {
	return "Daffa", "Rio"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
