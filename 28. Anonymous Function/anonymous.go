package main
import "fmt"

type BlackList func(string) bool 

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Hello", name)
	}
}

func main() {
	blacklist := func(name string) bool{
		return name == "anjay"
	}
	registerUser ("Daffa", blacklist)

	registerUser ("anjay", func(name string) bool {
		return name == "anjay"
	})
}