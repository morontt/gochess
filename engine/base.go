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
    Name  byte
    Color int8
    X, Y  byte
    Dead  bool
}

type Position struct {
    Figures [32]Chessman
}

func InitPosition() Position {
    var (
        res    Position
        color  int8
        hf, hp byte
        deltaP int
    )

    for i := 0; i < 2; i++ {
        delta := 8 * i

        if i == 0 {
            color = WHITE
            hf = 0
            hp = 1
            deltaP = 16
        } else {
            color = BLACK
            hf = 7
            hp = 6
            deltaP = 24
        }

        res.Figures[0 + delta] = Chessman{ROOK, color, 0, hf, false}
        res.Figures[1 + delta] = Chessman{KNIGHT, color, 1, hf, false}
        res.Figures[2 + delta] = Chessman{BISHOP, color, 2, hf, false}
        res.Figures[3 + delta] = Chessman{QUEEN, color, 3, hf, false}
        res.Figures[4 + delta] = Chessman{KING, color, 4, hf, false}
        res.Figures[5 + delta] = Chessman{BISHOP, color, 5, hf, false}
        res.Figures[6 + delta] = Chessman{KNIGHT, color, 6, hf, false}
        res.Figures[7 + delta] = Chessman{ROOK, color, 7, hf, false}

        for k := 0; k < 8; k++ {
            res.Figures[deltaP + k] = Chessman{PAWN, color, byte(k), hp, false}
        }
    }

    return res
}

