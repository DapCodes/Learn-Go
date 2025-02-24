package main
import "fmt"

func main() {
    days := [...]string {
		"Senin", 
		"Selasa",
		"Rabu",
		"Kamis", 
		"Jumat", 
		"Sabtu",
		"Minggu",
	}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[0] = "Minnggu Baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur Baru")
	// membuat arr baru yang di akhir nya ada "Libur Baru"
	fmt.Println(daysSlice2)
	fmt.Println(daysSlice1)
	fmt.Println(days)

}


