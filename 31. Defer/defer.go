package main

import "fmt"

func logging() {
	fmt.Println("Function selesai di laksanakan")
}

func runApp() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApp()
}
