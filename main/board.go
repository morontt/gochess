package main

import (
    "fmt"
    "gochess/engine"
)

func showBoard(p engine.Position) {
    str0 := "   +---+---+---+---+---+---+---+---+"
    str1 := "|   |   |   |   |   |   |   |   |"
    str2 := "     A   B   C   D   E   F   G   H"

    for i := 0; i < 9; i++ {
        fmt.Println(str0)
        if i != 8 {
            fmt.Printf(" %d %s\n", 8 - i, str1)
        }
    }

    fmt.Println(str2)
}

