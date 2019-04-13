
package valakut

import (

//    "fmt"

    "strings"
    "strconv"
)

// ---------------------------------------------------------------------

type state struct {
    play bool
    deck []string
    hand []string
    board []string
    mana int
    lands int
    turn int
    Log string
    done bool
}

// ---------------------------------------------------------------------

func (s *state) next() []state {

    states := []state{}

    states = append(states, s.clone_play_forest()...)
    states = append(states, s.clone_cast_explore()...)
    states = append(states, s.clone_cast_primeval_titan()...)
    states = append(states, s.clone_pass_turn()...)

    return states
}

// ---------------------------------------------------------------------

func (s state) clone_pass_turn() []state {
    s.pass_turn()
    return []state{s}
}

func (s *state) clone_play_forest() []state {
    card := "Forest"
    if !s.in_hand(card) || s.lands == 0 { return []state{} }
    clone := s.clone()
    clone.note("Playing " + card)
    clone.play_forest()
    return []state{clone}
}

func (s *state) clone_cast_explore() []state {
    card := "Explore"
    cost := 2
    if !s.in_hand(card) || !s.can_pay(cost) { return []state{} }
    clone := s.clone()
    clone.note("Casting " + card)
    clone.cast_explore()
    return []state{clone}
}

func (s *state) clone_cast_primeval_titan() []state {
    card := "Titan"
    cost := 6
    if !s.in_hand(card) || !s.can_pay(cost) { return []state{} }
    clone := s.clone()
    clone.note("Casting " + card)
    clone.cast_primeval_titan()
    return []state{clone}
}

// ---------------------------------------------------------------------

func (s *state) pass_turn() {
    s.turn += 1
    s.lands = 1
    // Reset mana pool
    s.mana = len(s.board)
    s.note("\nTurn " + strconv.Itoa(s.turn))
    s.note("Hand: " + tally(s.hand))
    s.note("Board: " + tally(s.board))
    if s.turn > 1 || !s.play { s.draw(1) }
}

func (s *state) play_forest() {
    s.hand = remove(s.hand, "Forest")
    s.board = append(s.board, "Forest")
    s.mana += 1
    s.lands -= 1
}

func (s *state) cast_explore() {
    s.hand = remove(s.hand, "Explore")
    s.pay(2)
    s.lands += 1
    s.draw(1)
}

func (s *state) cast_primeval_titan() {
    s.hand = remove(s.hand, "Titan")
    s.pay(6)
    s.done = true
}




// ---------------------------------------------------------------------

func (s *state) can_pay(cost int) bool { return cost <= s.mana }

func (s *state) pay(cost int) { s.mana = s.mana - cost }

// ---------------------------------------------------------------------

func (s *state) note(line string) { s.Log += "\n" + line }

func (s state) clone() state {
    // Deep copy the slices.
    s.deck = append([]string{}, s.deck...)
    s.hand = append([]string{}, s.hand...)
    s.board = append([]string{}, s.board...)
    return s
}

func (s *state) flip() {
    s.play = flip()
    if s.play { s.note("On the play") } else { s.note("On the draw") }
}

func (s *state) shuffle() {
    s.note("Shuffling")
    s.deck, s.hand = shuffled(append(s.deck, s.hand...)), []string{}
}

func (s *state) draw(n int) {
    s.note("Drawing " + strings.Join(s.deck[:n], " "))
    s.deck, s.hand = s.deck[n:], append(s.hand, s.deck[:n]...)
}

func (s *state) mill(n int) {
    s.note("Milling " + s.deck[0])
    s.deck = append(s.deck[n:], s.deck[:n]...)
}

func (s *state) in_hand(card string) bool {
    for _, c := range s.hand {
        if c == card { return true }
    }
    return false
}
