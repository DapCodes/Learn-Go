package main
import "fmt"

func main() {
    pw := "Daffa"
	length := len(pw)

	switch length >= 8 {
		case true:
			fmt.Println("Password berhasil di ubah")
		case false:
			err := 8 - length
			fmt.Println("Password anda kurang", err, "element")
	}
}
