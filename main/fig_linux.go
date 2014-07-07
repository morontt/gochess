package main

import "gochess/engine"

func figureString(ch engine.Chessman) string {
    var fig, colorOpen string

    colorClose := "\033[00m "

    if ch.Color == engine.WHITE {
        colorOpen = " \033[01;33m"
    } else {
        colorOpen = " \033[01;34m"
    }

    switch ch.Name {
        case engine.PAWN:
            fig = "P"
        case engine.BISHOP:
            fig = "B"
        case engine.KNIGHT:
            fig = "K"
        case engine.ROOK:
            fig = "R"
        case engine.QUEEN:
            fig = "Q"
        case engine.KING:
            fig = "X"
        default:
            fig = "*"
    }

    return colorOpen + fig + colorClose
}

