package main
import "fmt"

func main() {
    pw := "Daffa"
	length := len(pw)

	switch {
		case length >= 8:
			fmt.Println("Password berhasil di ubah")
		case length <= 7:
			err := 8 - length
			fmt.Println("Password anda kurang", err, "element")
	}
}
