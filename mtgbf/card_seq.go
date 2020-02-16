package mtgbf


import (
    "math/rand"
    "time"
)


type card_seq struct {
    cards []card_obj
}


func Cards(card_names ...string) card_seq {
    cards := []card_obj{}
    for _, cn :=  range card_names {
        cards = append(cards, Card(cn))
    }
    return card_seq{cards: cards}
}


func (self *card_seq) Plus(other card_seq) card_seq {
    cards := self.cards
    for _, c :=  range other.cards {
        cards = append(cards, c)
    }
    return card_seq{cards: cards}
}


func (self *card_seq) shuffle() {
    // This shouldn't happen very often. Should be fine to re-seed every time
    rand.Seed(time.Now().UTC().UnixNano())
    cards_new := make([]card_obj, len(self.cards))
    for i, j := range rand.Perm(len(self.cards)) {
        cards_new[i] = self.cards[j]
    }
    self.cards = cards_new
}

func (self *card_seq) Pretty() string {
    card_names := []string{}
    for _, c := range self.cards {
        card_names = append(card_names, c.Pretty())
    }
    return tally_strings(card_names)
}


func (self *card_seq) top(i int) card_seq {
    // In Python, self[:n] for positive or negative n
    if i >= 0 {
        return card_seq{cards: self.cards[:i]}
    } else {
        return card_seq{cards: self.cards[:len(self.cards) + i]}
    }
}


func (self *card_seq) after_top(i int) card_seq {
    // In Python, self[n:] for positive or negative n
    if i >= 0 {
        return card_seq{cards: self.cards[i:]}
    } else {
        return card_seq{cards: self.cards[len(self.cards) + i:]}
    }
}
