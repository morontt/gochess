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

    str0 := "   +---+---+---+---+---+---+---+---+"
    str1 := "     A   B   C   D   E   F   G   H"

    for i = 0; i < 9; i++ {
        fmt.Println(str0)
        if i != 8 {
            fmt.Printf(" %d |", 8 - i)
            for j = 0; j < 8; j++ {
                fig = "   "
                for k := 0; k < 32; k++ {
                    if p.Figures[k].X == j && p.Figures[k].Y == 7 - i {
                        switch p.Figures[k].Name {
                            case engine.PAWN:
                                fig = " P"
                            case engine.BISHOP:
                                fig = " B"
                            case engine.KNIGHT:
                                fig = " K"
                            case engine.ROOK:
                                fig = " R"
                            case engine.QUEEN:
                                fig = " Q"
                            case engine.KING:
                                fig = " X"
                            default:
                                fig = " *"
                        }

                        if p.Figures[k].Color == engine.WHITE {
                            fig += "w"
                        } else {
                            fig += "b"
                        }
                    }
                }
                fmt.Printf("%s|", fig)
            }
            fmt.Print("\n")
        }
    }

    fmt.Println(str1)
}

