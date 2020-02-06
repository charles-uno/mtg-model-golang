package mtgbf


import (
    "math/rand"
    "time"
)


type card_list struct {
    cards []card
}


func Cards(card_names ...string) card_list {
    cards := []card{}
    for _, cn :=  range card_names {
        cards = append(cards, Card(cn))
    }
    return card_list{cards: cards}
}


func (self *card_list) shuffle() {
    // This shouldn't happen very often. Should be fine to re-seed every time
    rand.Seed(time.Now().UTC().UnixNano())
    cards_new := make([]card, len(self.cards))
    for i, j := range rand.Perm(len(self.cards)) {
        cards_new[i] = self.cards[j]
    }
    self.cards = cards_new
}

func (self *card_list) Pretty() string {
    card_names := []string{}
    for _, c := range self.cards {
        card_names = append(card_names, c.Pretty())
    }
    return tally_strings(card_names)
}
