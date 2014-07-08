package engine

import (
    "strings"
    "regexp"
)

type Field struct {
    X, Y byte
}

var StringToCoord = map[string]byte{
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

func Move(p *Position, a1, a2 string) bool {
    field1, success1 := ParseField(a1)
    field2, success2 := ParseField(a2)

    result := false

    if (success1 && success2) {
        for i := 0; i < 32; i++ {
            if !p.Figures[i].Dead && p.Figures[i].X == field1.X && p.Figures[i].Y == field1.Y {
                p.Figures[i].X = field2.X
                p.Figures[i].Y = field2.Y
                result = true

                break
            }
        }
    }

    return result
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

