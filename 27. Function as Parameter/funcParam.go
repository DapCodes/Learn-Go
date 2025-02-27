package main
import "fmt"

type Filter func(string) string
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjay" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Daffa", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjay", filter)
}