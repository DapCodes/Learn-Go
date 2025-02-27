package main
import "fmt"

func getGoodBye(name string) string {
	goodbye := "Selamat Tinggal " + name
	return goodbye
}

func main() {
	 bye := getGoodBye
	 fmt.Println(bye("Daffa"))
}