package engine

import (
    "strings"
    "regexp"
)

type Field struct {
    X, Y int8
}

type Vector struct {
    X, Y int8
}

func (v *Vector) Create(f1, f2 Field) {
    v.X = f2.X - f1.X
    v.Y = f2.Y - f1.Y
}

var StringToCoord = map[string]int8{
    "a": 0,
    "b": 1,
    "c": 2,
    "d": 3,
    "e": 4,
    "f": 5,
    "g": 6,
    "h": 7,
    "1": 0,
    "2": 1,
    "3": 2,
    "4": 3,
    "5": 4,
    "6": 5,
    "7": 6,
    "8": 7,
}

func Move(p *Position, blackMove bool, a1, a2 string) bool {
    var (
        idx1, idx2 int
        vect       Vector
    )

    field1, success1 := ParseField(a1)
    field2, success2 := ParseField(a2)

    search := false
    kick := false

    if (success1 && success2) {
        vect.Create(field1, field2)
        for i := 0; i < 32; i++ {
            if !p.Figures[i].Dead && p.Figures[i].X == field1.X && p.Figures[i].Y == field1.Y {
                idx1 = i
                search = true
            }

            if !p.Figures[i].Dead && p.Figures[i].X == field2.X && p.Figures[i].Y == field2.Y {
                idx2 = i
                kick = true
            }
        }

        if search && (blackMove && p.Figures[idx1].Color == WHITE || !blackMove && p.Figures[idx1].Color == BLACK) {
            search = false
        }

        if (search && isLegalMove(p, kick, idx1, idx2, vect)) {
            p.Figures[idx1].X = field2.X
            p.Figures[idx1].Y = field2.Y

            if kick {
                p.Figures[idx2].Dead = true
            }
        } else {
            search = false
        }
    }

    return search
}

func ParseField(s string) (Field, bool) {
    s = strings.ToLower(s)

    result := Field{0, 0}
    success := false

    validString := regexp.MustCompile("^[a-h][1-8]$")
    if (validString.MatchString(s)) {
        result.X = StringToCoord[string(s[0])]
        result.Y = StringToCoord[string(s[1])]
        success = true
    }

    return result, success
}

func isLegalMove(p *Position, kick bool, idx1, idx2 int, vect Vector) bool {
    var (
        pureVector bool = false
    )

    if kick && p.Figures[idx1].Color == p.Figures[idx2].Color {
        return false
    }

    switch p.Figures[idx1].Name {
        case KNIGHT:
            if abs(vect.X) == 1 && abs(vect.Y) == 2 {
                pureVector = true
            }
            if abs(vect.X) == 2 && abs(vect.Y) == 1 {
                pureVector = true
            }
        case BISHOP:
            if abs(vect.X) == abs(vect.Y) {
                pureVector = true
            }
        case ROOK:
            if vect.X == 0 || vect.Y == 0 {
                pureVector = true
            }
        case QUEEN:
            if abs(vect.X) == abs(vect.Y) {
                pureVector = true
            }
            if vect.X == 0 || vect.Y == 0 {
                pureVector = true
            }
        case KING:
            if (abs(vect.X) == 1 || vect.X == 0) && (abs(vect.Y) == 1 || vect.Y == 0) {
                pureVector = true
            }
    }

    if !pureVector {
        return false
    }

    return true
}

func abs(x int8) int8 {
    switch {
        case x < 0:
            return -x
        case x == 0:
            return 0
    }

    return x
}

