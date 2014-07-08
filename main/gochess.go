package main

import (
    "fmt"
    "gochess/engine"
)

func main() {
    var (
        move1, move2 string
    )

    pos := engine.InitPosition()

    showBoard(pos)
    fmt.Println()

    gameOver := false
    for !gameOver {
        fmt.Print("> ")
        _, err := fmt.Scanln(&move1, &move2)

        if move1 == "exit" || move1 == "quit" {
            break
        }

        if err != nil {
            fmt.Println("error")
        } else {
            if engine.Move(&pos, move1, move2) {
                showBoard(pos)
                fmt.Println()
            } else {
                fmt.Println("error")
            }
        }
    }
}

