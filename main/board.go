package main

import (
    "fmt"
    "gochess/engine"
)

func showBoard(p engine.Position) {
    var (
        i, j byte
        fig  string
    )

    fmt.Println(" GoChess 0.0.9\n")

    str0 := "   +---+---+---+---+---+---+---+---+"
    str1 := "     A   B   C   D   E   F   G   H"

    for i = 0; i < 9; i++ {
        fmt.Println(str0)
        if i != 8 {
            fmt.Printf(" %d |", 8 - i)
            for j = 0; j < 8; j++ {
                fig = "   "
                for k := 0; k < 32; k++ {
                    if !p.Figures[k].Dead && p.Figures[k].X == j && p.Figures[k].Y == 7 - i {
                        fig = figureString(p.Figures[k])
                    }
                }
                fmt.Printf("%s|", fig)
            }
            fmt.Print("\n")
        }
    }

    fmt.Println(str1)
}

