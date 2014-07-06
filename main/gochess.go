package main

import (
    "fmt"
    "gochess/engine"
)

func main() {
    f := engine.InitPosition()

    fmt.Println()
    showBoard(f)
    fmt.Println()
}

