
package valakut

import (
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
    for _, card := range uniques(s.hand) {
        states = append(states, s.clone_play(card)...)
    }
    states = append(states, s.clone_pass_turn()...)
    return states
}

// ---------------------------------------------------------------------

func (s *state) clone_pass_turn() []state {
    clone := s.clone()
    return clone.pass_turn()
}

func (s *state) clone_play(card string) []state {
    if !s.in_hand(card) { return []state{} }
    if is_land(card) && s.lands == 0 { return []state{} }
    cost := get_cost(card)
    if !s.can_pay(cost) { return []state{} }
    clone := s.clone()
    clone.note("Playing " + slug(card))
    clone.hand = remove(clone.hand, card)
    if is_land(card) { clone.lands -= 1 } else { clone.pay(cost) }
    switch card {
        case "Explore":
            return clone.play_explore()
        case "Forest":
            return clone.play_forest()
        case "Primeval Titan":
            return clone.play_primeval_titan()
        case "Search for Tomorrow":
            return clone.play_search_for_tomorrow()
        case "Shefet Monitor":
            return clone.play_shefet_monitor()
        case "Sleight of Hand":
            return clone.play_sleight_of_hand()
        case "Summoner's Pact":
            return clone.play_summoners_pact()
        case "Through the Breach":
            return clone.play_through_the_breach()
    }
    return []state{}
}

func get_cost(card string) int {
    switch card {
        case "Explore":
            return 2
        case "Primeval Titan":
            return 6
        case "Sakura-Tribe Elder":
            return 2
        // Treat SFT as a 1-cost spell with kicker 2.
        case "Search for Tomorrow":
            return 1
        case "Shefet Monitor":
            return 4
        case "Sleight of Hand":
            return 1
        case "Summoner's Pact":
            return 6
        case "Through the Breach":
            return 5
        default:
            return 0
    }
}





// ---------------------------------------------------------------------

func (s *state) pass_turn() []state {
    s.turn += 1
    s.lands = 1
    // Reset mana pool
    s.mana = 0
    for _, card := range s.board {
        if is_land(card) { s.mana += 1 }
    }
    s.note("\nTurn " + strconv.Itoa(s.turn))
    s.note("Hand: " + tally(s.hand))
    s.note("Board: " + tally(s.board))
    // Handle anything suspended
    for i, card := range s.board {
        if is_land(card) { continue }
        // Remove a counter from anything suspended.
        if card[:1] == "." {
            s.note("Ticking down " + card)
            s.board[i] = s.board[i][1:]
            card = s.board[i]
        }
        // Cast anything that's out of counters.
        if card[:1] != "." {
            s.play_suspended(card)
        }
    }
    if s.turn > 1 || !s.play { s.draw(1) }
    return []state{*s}
}

func (s *state) play_suspended(card string) {
    s.note("Playing " + card + " from suspend")
    s.board = remove(s.board, card)
    switch card {
        case "Search for Tomorrow":
            s.board = append(s.board, "Forest")
            s.mana += 1
    }
}

func (s *state) play_forest() []state {
    s.board = append(s.board, "Forest")
    s.mana += 1
    return []state{*s}
}

func (s *state) play_explore() []state {
    s.lands += 1
    s.draw(1)
    return []state{*s}
}

func (s *state) play_primeval_titan() []state {
    s.done = true
    return []state{*s}
}

func (s *state) play_sakura_tribe_elder() []state {
    s.board = append(s.board, "Forest")
    return []state{*s}
}

func (s *state) play_search_for_tomorrow() []state {
    // Treat suspending as the default case, kicker 2 to get it now.
    states := []state{}
    if s.can_pay(2) {
        clone := s.clone()
        clone.pay(2)
        clone.mana += 1
        clone.board = append(clone.board, "Forest")
        states = append(states, clone)
    }
    s.unnote()
    s.note("Suspending " + slug("Search for Tomorrow"))
    s.board = append(s.board, "..Search for Tomorrow")
    states = append(states, *s)
    return states
}

func (s *state) play_shefet_monitor() []state {
    s.board = append(s.board, "Forest")
    s.draw(1)
    s.mana += 1
    return []state{*s}
}

func (s *state) play_sleight_of_hand() []state {
    choices := s.deck[:2]
    s.deck = s.deck[2:]
    clone := s.clone()
    s.note("Taking " + slug(choices[0]) + " over " + slug(choices[1]))
    s.hand = append(s.hand, choices[0])
    s.deck = append(s.deck, choices[1])
    clone.note("Taking " + slug(choices[1]) + " over " + slug(choices[0]))
    clone.hand = append(clone.hand, choices[0])
    clone.deck = append(clone.deck, choices[1])
    return []state{*s, clone}
}

func (s *state) play_summoners_pact() []state {
    return s.play_primeval_titan()
}

func (s *state) play_through_the_breach() []state {
    if !s.in_hand("Primeval Titan") && !s.in_hand("Summoner's Pact") {
        return []state{}
    }
    s.done = true
    return []state{*s}
}



// ---------------------------------------------------------------------

func (s *state) can_pay(cost int) bool { return cost <= s.mana }

func (s *state) pay(cost int) { s.mana = s.mana - cost }

// ---------------------------------------------------------------------

func (s *state) note(line string) { s.Log = s.Log + "\n" + line }

func (s *state) unnote() {
    lines := strings.Split(s.Log, "\n")
    s.Log = strings.Join(lines[:len(lines)-1], "\n")
}

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
    s.note("Drawing " + tally(s.deck[:n]))
    s.deck, s.hand = s.deck[n:], append(s.hand, s.deck[:n]...)
}

func (s *state) mill(n int) {
    s.note("Milling " + tally(s.deck[:n]))
    s.deck = append(s.deck[n:], s.deck[:n]...)
}

func (s *state) in_hand(card string) bool {
    for _, c := range s.hand {
        if c == card { return true }
    }
    return false
}
