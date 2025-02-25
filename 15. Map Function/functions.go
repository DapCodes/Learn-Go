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

	// LEN
	len := len(siswa["name"])
	fmt.Println("ini len", len)

	// select by key
	pick := siswa["class"]
	fmt.Println("ini pick", pick)

	// change data. daffa -> dhea
	person["name"] = "Dhea" 
	fmt.Println("ini change", person["name"])

	// make. membuat map baru
	newMap := make(map[string]string)
	newMap["name"] = "Ani"
	newMap["adress"] = "bandung"
	fmt.Println("ini map baru", newMap)

	// delete, menghapus data dalam map
	delete(newMap, "name")
	fmt.Println(newMap)
}
