package main
import "fmt"

func main() {

	person := map[string]string{
		"name" : "Daffa",
		"address" : "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// i have 2 side :b

	var siswa = map[string]string{}
	siswa["name"] = "Rio"
	siswa["class"] = "1 B"

	fmt.Println(siswa)
	fmt.Println(siswa["name"])
	fmt.Println(siswa["class"])
}
