package main
import "fmt"

func main() {

	// NoKtp alias String
	type NoKtp string

	var noKtpDaffa NoKtp = "10101010"

	var contoh string = "12345678"
	var ktpContoh NoKtp = NoKtp(contoh)

	fmt.Println(noKtpDaffa, "No KTP 1")
	fmt.Println(ktpContoh, "No KTP 2")
	fmt.Println(NoKtp("20202020 No KTP 3"))

}