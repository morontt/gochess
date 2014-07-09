package main

import (
    "fmt"
    "flag"
    "gochess/engine"
)

var (
    computerFirst, blackMove bool
)

func init() {
    flag.BoolVar(&computerFirst, "comp", false, "first turn the computer")
    flag.BoolVar(&blackMove, "black", false, "black's move")
}

func main() {
    var (
        move1, move2, cursor string
    )

    flag.Parse()

    pos := engine.InitPosition()

    showBoard(pos)
    fmt.Println()

    gameOver := false
    for !gameOver {
        if blackMove {
            cursor = "BLACK"
        } else {
            cursor = "WHITE"
        }

        fmt.Print(cursor + " > ")
        _, err := fmt.Scanln(&move1, &move2)

        if move1 == "exit" || move1 == "quit" {
            break
        }

        if err != nil {
            fmt.Println("error")
        } else {
            if engine.Move(&pos, blackMove, move1, move2) {
                showBoard(pos)
                fmt.Println()
                blackMove = !blackMove
            } else {
                fmt.Println("error")
            }
        }
    }
}

