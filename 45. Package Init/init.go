package main

import (
	"fmt"
	"go-module/database"
	_ "go-module/helper" // package ini tetap di jalankan dan tidak error meskipun tidak di pakai
)

func main() {
	fmt.Println(database.Getdatabase())
}
