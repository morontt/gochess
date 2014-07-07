package main

import "gochess/engine"

func figureString(ch engine.Chessman) string {
    var fig string

    switch ch.Name {
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

    if ch.Color == engine.WHITE {
        fig += "w"
    } else {
        fig += "b"
    }

    return fig
}

