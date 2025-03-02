package main

import "fmt"

func ups() interface{} {
	return "Hallo"
}

func yaps() any {
	return false
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)

	var empty any = yaps()
	fmt.Println(empty)
}
