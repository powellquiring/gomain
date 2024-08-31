package main

import (
    "fmt"

    "github.com/powellquiring/gomod"
)

func main() {
    // Get a greeting message and print it.
    message := gomod.Hello("Gladys")
    fmt.Println(message)
}

