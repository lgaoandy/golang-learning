package main

import "fmt"

// go mod tidy - locates and downloads rsc.io/quote module from import
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
}