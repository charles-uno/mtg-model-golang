package mtgbf


import (
    "errors"
    "strconv"
)


type mana struct {
    green int
    total int
}


func Mana(expr string) mana {
    // We only track green and total, but let's still accept expressions with
    // other mana symbols.
    green := count_runes(expr, 'G')
    total := green
    // Figure out colored mana. Alto tolerate C for colorless
    for _, r := range "WUBRC" {
        total += count_runes(expr, r)
    }
    // Figure out generic mana. Each digit is counted independently
    for i, r := range "0123456789" {
        total += count_runes(expr, r)*i
    }
    return mana{green: green, total: total}
}


func (self *mana) plus(other mana) mana {
    return mana{
        green: self.green + other.green,
        total: self.total + other.total,
    }
}


func (self *mana) minus(other mana) (mana, error) {
    if self.total < other.total || self.green < other.green {
        return mana{}, errors.New("can't subtract " + self.Pretty() + " - " + other.Pretty())
    }
    total := self.total - other.total
    green := self.green - other.green
    if green > total {
        return mana{total: total, green: total}, nil
    } else {
        return mana{total: total, green: green}, nil
    }
}



func (m *mana) Pretty() string {
    if m.total == 0 {
        return "0"
    }
    expr := ""
    if m.total > m.green {
        expr += strconv.Itoa(m.total - m.green)
    }
    for i := 0; i < m.green; i++ {
        expr += "G"
    }
    return expr
}


func count_runes(s string, r rune) int {
    count := 0
    for _, char := range s {
        if char == r { count += 1 }
    }
    return count
}
