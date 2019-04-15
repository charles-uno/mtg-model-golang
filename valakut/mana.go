
package valakut

import (
    "strconv"
)
// ---------------------------------------------------------------------

type mana struct {
    green int
    total int
}

func (m *mana) empty() {
    m.total = 0
    m.green = 0
}

func (m *mana) can_pay(s string) bool {
    dm := get_mana(s)
    return m.total >= dm.total && m.green >= dm.green
}

func (m *mana) pay(s string) {
    dm := get_mana(s)
    m.total -= dm.total
    m.green -= dm.green
    // If we have 2GG and we need to pay 3, end up with G.
    if m.green > m.total { m.green = m.total }
}

func (m *mana) add(s string) {
    dm := get_mana(s)
    m.total += dm.total
    m.green += dm.green
}

func get_mana(s string) mana {
    m := mana{}
    for _, c := range s {
        if c == 'G' { m.total += 1; m.green += 1 }
        if c == 'R' { m.total += 1 }
        if '0' <= c && c <= '9' { m.total += int(c - '0') }
    }
    return m
}

func (m *mana) show() string {
    s := ""
    if m.total > m.green {
        s += strconv.Itoa(m.total - m.green)
    }
    for i := 0; i < m.green; i++ { s += "G" }
    return s
}
