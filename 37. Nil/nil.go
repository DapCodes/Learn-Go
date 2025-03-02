package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("Daffa")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}
}
