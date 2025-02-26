package main
import "fmt"

func main() {
    counter := 1
    for counter <= 10 {
        j := 1
        for j <= counter {
            fmt.Print("* ")
            j++
        }
        fmt.Println()
        counter++
    }
}




