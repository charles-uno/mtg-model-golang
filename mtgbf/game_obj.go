package mtgbf


type game_obj struct {
    battlefield card_seq
    hand card_seq
    library card_seq
    mana_pool mana_obj
    report string
}


func InitialGameState(deck_name string) game_obj {
    library, err := LoadList(deck_name)
    if err != nil { panic(err) }
    library.shuffle()
    initial_game_state := game_obj{
        hand: card_seq{},
        library: library,
        mana_pool: Mana(""),
        report: "turn 0",
    }
    return initial_game_state.draw(7)
}


func (self *game_obj) Pretty() string {
    return "hand: " + self.hand.Pretty() + "\n" +
        "battlefield: " + self.battlefield.Pretty() + "\n" +
        "pool: " + self.mana_pool.Pretty() + "\n" +
        self.report
}


func (self *game_obj) draw(n int) game_obj {
    cards := self.library.top(n)
    return game_obj{
        battlefield: self.battlefield,
        hand: self.hand.Plus(cards),
        library: self.library.after_top(n),
        mana_pool: self.mana_pool,
        report: self.report + ", draw " + cards.Pretty(),
    }
}
