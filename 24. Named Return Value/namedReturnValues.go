package main
import "fmt"

func getCompleteName() (firstName, middleName, lastName string){
	firstName = "Daffa"
	middleName = "Ramadhan"
	lastName = "Maulana"

	return firstName, middleName, lastName
}

func main() {
	a, b, _ := getCompleteName()
	fmt.Println(a, b)
}