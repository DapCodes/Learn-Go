package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi error,", message)

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(true)
}
