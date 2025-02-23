package main
import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai16, "int16")
	fmt.Println(nilai32, "int32")
	fmt.Println(nilai64, "int64")


	var name = "Daffa"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(e, "| Dalam bentuk Byte")
	fmt.Println(eString, "| Dalam bentuk String")
}
