package mtgbf


import (
    "strconv"
)


type game_state struct {
    battlefield []string
    hand []string
    land_drops int
    library []string
    mana_pool mana
    report string
    turn int
}


func InitialGameState(deck_name string) game_state {
    gs := game_state{
        battlefield: []string{},
        hand: []string{},
        land_drops: 0,
        library: load_list(deck_name),
        mana_pool: Mana(""),
        report: "turn 0",
        turn: 0,
    }
    gs = gs.draw(7)

    gs.pass_turn()

    gs.battlefield = append(gs.battlefield, "Amulet of Vigor")
    gs.battlefield = append(gs.battlefield, "Amulet of Vigor")

    for _, card := range gs.hand {
        if is_land(card) {
            gs = gs.play(card)
        }
    }
    return gs
}


func (self *game_state) Pretty() string {
    return "hand: " + pretty(self.hand...) + "\n" +
        "battlefield: " + pretty(self.battlefield...) + "\n" +
        "pool: " + self.mana_pool.Pretty() + "\n" +
        self.report
}


// ----------------------------------------------------------------------------


func (copy game_state) add_mana(m mana) game_state {
    copy.mana_pool = copy.mana_pool.plus(m)
    copy.report += ", " + copy.mana_pool.Pretty() + " in pool"
    return copy
}


func (self *game_state) cast(card []string) game_state {
    return game_state{}
}


func (copy game_state) draw(n int) game_state {
    copy.hand = append(copy.hand, copy.library[:n]...)
    copy.report += ", draw " + pretty(copy.library[:n]...)
    copy.library = copy.library[n:]
    return copy
}


func (copy game_state) pass_turn() game_state {
    copy.land_drops = 1
    copy.mana_pool = Mana("")
    copy.turn += 1
    copy.report += "\nturn " + strconv.Itoa(copy.turn)
    return copy.draw(1)
}


func (self *game_state) play(card string) game_state {
    if enters_tapped(card) {
        return self.play_tapped(card)
    } else {
        return self.play_untapped(card)
    }
}


func (copy game_state) play_tapped(card string) game_state {
    copy.battlefield = append(copy.battlefield, card)
    copy.hand = discard(copy.hand, card)
    copy.report += "\nplay " + pretty(card)
    n_amulets := count_strings(copy.battlefield, "Amulet of Vigor")
    for i := 0; i < n_amulets; i++ {
        copy = copy.add_mana(taps_for(card))
    }
    return copy
}


func (copy game_state) play_untapped(card string) game_state {
    copy.battlefield = append(copy.battlefield, card)
    copy.hand = discard(copy.hand, card)
    copy.report += "\nplay " + pretty(card)
    return copy.add_mana(taps_for(card))
}

// ----------------------------------------------------------------------------


func (copy game_state) cast_explore() game_state {
    copy.land_drops += 1
    return copy.draw(1)
}
