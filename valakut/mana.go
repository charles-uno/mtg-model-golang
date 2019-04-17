
package valakut

import (
    "strconv"
)

// ---------------------------------------------------------------------

type mana struct {
    green int
    total int
}

func Mana(ms string) mana {
    m := mana{}
    m.add(ms)
    return m
}

// ---------------------------------------------------------------------

func (m *mana) empty() {
    m.total = 0
    m.green = 0
}

func (m *mana) can_pay(ms string) bool {
    dm := parse_mana_string(ms)
    return m.total >= dm.total && m.green >= dm.green
}

func (m *mana) pay(ms string) {
    dm := parse_mana_string(ms)
    m.total -= dm.total
    m.green -= dm.green
    // If we have 2GG and we need to pay 3, end up with G.
    if m.green > m.total { m.green = m.total }
}

func (m *mana) add(ms string) {
    dm := parse_mana_string(ms)
    m.total += dm.total
    m.green += dm.green
}

func parse_mana_string(ms string) mana {
    m := mana{}
    for _, c := range ms {
        if c == 'G' { m.total += 1; m.green += 1 }
        if c == 'R' { m.total += 1 }
        if '0' <= c && c <= '9' { m.total += int(c - '0') }
    }
    return m
}

func (m *mana) show() string {
    ms := ""
    if m.total > m.green {
        ms += strconv.Itoa(m.total - m.green)
    }
    for i := 0; i < m.green; i++ { ms += "G" }
    return ms
}
