package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "Hello, World!"
    text = strings.ToUpper(text)
    fmt.Println(text) // Outputs: HELLO, WORLD!
}
