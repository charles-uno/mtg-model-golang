package mtgbf


import (
    "strconv"
)


type mana_obj struct {
    wubrg [5]int
    total int
}


func Mana(ms string) mana_obj {
    m := mana_obj{}
    // Figure out colored mana
    for i, r := range "WUBRG" {
        m.wubrg[i] += count_runes(ms, r)
    }
    // Figure out generic mana. Each digit is counted independently
    for i, r := range "0123456789" {
        m.total += count_runes(ms, r)*i
    }
    // Total includes colored mana as well
    for _, n := range m.wubrg {
        m.total += n
    }
    return m
}


func count_runes(s string, r rune) int {
    count := 0
    for _, char := range s {
        if char == r { count += 1 }
    }
    return count
}


func (m *mana_obj) Pretty() string {
    ms := ""
    mana_color := 0
    for i, r := range "WUBRG" {
        for j := 0; j < m.wubrg[i]; j++ {
            ms += string(r)
        }
        mana_color += m.wubrg[i]
    }
    mana_leftover := m.total - mana_color
    if mana_leftover > 0 {
        ms = strconv.Itoa(mana_leftover) + ms
    }
    return ms
}
