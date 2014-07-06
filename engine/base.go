package engine

const (
    PAWN   byte = 1
    BISHOP byte = 2
    KNIGHT byte = 3
    ROOK   byte = 4
    QUEEN  byte = 5
    KING   byte = 6

    BLACK  int8 = -1
    WHITE  int8 = 1
)

type Chessman struct {
    name byte
    color int8
    X, Y byte
    dead bool
}

type Position struct {
    Figures [32]Chessman
}

func InitPosition() Position {
    var (
        res   Position
        color int8
        h     byte
    )

    for i := 0; i < 2; i++ {
        delta := 8 * i
        
        if i == 0 {
            color = WHITE
            h = 0
        } else {
            color = BLACK
            h = 7
        }
        
        res.Figures[0 + delta] = Chessman{ROOK, color, 0, h, false}
        res.Figures[1 + delta] = Chessman{KNIGHT, color, 1, h, false}
        res.Figures[2 + delta] = Chessman{BISHOP, color, 2, h, false}
        res.Figures[3 + delta] = Chessman{QUEEN, color, 3, h, false}
        res.Figures[4 + delta] = Chessman{KING, color, 4, h, false}
        res.Figures[5 + delta] = Chessman{BISHOP, color, 5, h, false}
        res.Figures[6 + delta] = Chessman{KNIGHT, color, 6, h, false}
        res.Figures[7 + delta] = Chessman{ROOK, color, 7, h, false}
    }
    
    return res
}

