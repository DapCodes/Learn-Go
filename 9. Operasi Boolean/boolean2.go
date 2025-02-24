package main
import "fmt"

func main() {
    var nilai = 90
	var absensi = 80

	var lulusNilai bool = nilai >= 80
	var lulusAbsensi bool = absensi >= 80

	var hasil = lulusAbsensi && lulusNilai
	fmt.Println(hasil)
}


