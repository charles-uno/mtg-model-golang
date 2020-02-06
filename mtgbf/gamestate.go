package mtgbf


type game_state struct {
    battlefield card_list
    hand card_list
    library card_list
    mana_pool mana
    on_the_play bool
    report string
}


func InitialGameState(library card_list, on_the_play bool) game_state {
    library.shuffle()
    return game_state{
        library: library,
        on_the_play: on_the_play,
    }
}


func (self *game_state) Pretty() string {
    return "hand: " + self.hand.Pretty() + "\n" +
        "battlefield: " + self.battlefield.Pretty() + "\n" +
        "pool: " + self.mana_pool.Pretty()
}
