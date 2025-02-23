package main
import "fmt"

func main() {
	var name string 
	var umur int
	var ipk = 3.7
	
	alamat := "Bandung"
	universitas := "ITB"
	jumlahSKS := 19

	name = "Daffa Ramadhan"
	umur = 17

	fmt.Println(
		"Hai nama saya", name,
		"usia saya", umur,
		"saya tinggal di", alamat,
		"saya kuliah di", universitas,
		"mengambil SKS sebanyak", jumlahSKS,
		"dan IPK saya", ipk,
	);
}