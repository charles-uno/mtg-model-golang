
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
    exile []string
    pool mana
    lands int
    turn int
    Log string
    done bool
    mulls int
}

func (s *state) id() string {
    // We don't care about order for hand, exile, or board. But we do
    // care for deck.
    return tally(s.hand) + ";" + tally(s.exile) + ";" + tally(s.board) + ";" + s.pool.show() + ";" + strconv.Itoa(s.lands) + ";" + strconv.FormatBool(s.done) + ";" + strings.Join(s.deck, " ")
}

func (s *state) Line() string {
    play := strconv.FormatBool(s.play)[:1]
    return "t:" + strconv.Itoa(s.turn) + " m:" + strconv.Itoa(s.mulls) + " p:" + play + "\n"
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

func (s *state) clone_mulligan() []state {
    ctop := s.clone()
    ctop.mulls += 1
    ctop.shuffle()
    ctop.draw(7 - ctop.mulls)
    cbot := ctop.clone()
    cbot.mill(1)
    return []state{ctop, cbot}
}

func (s *state) clone_play(card string) []state {
    if !s.in_hand(card) { return []state{} }
    // Don't bail for spell lands like Sheltered Thicket.
    if !is_spell(card) && s.lands == 0 { return []state{} }
    cost := get_cost(card)
    if !s.pool.can_pay(cost) { return []state{} }
    clone := s.clone()
    clone.note(slug(card))
    clone.hand = remove(clone.hand, card)
    // Careful -- we may accidentally over-play from Sheltered Thicket.
    if is_land(card) { clone.lands -= 1 } else { clone.pool.pay(cost) }
    switch card {
        case "Cinder Glade":
            return clone.play_cinder_glade()
        case "Explore":
            return clone.play_explore()
        case "Forest":
            return clone.play_forest()
        case "Mountain":
            return clone.play_mountain()
        case "Primeval Titan":
            return clone.play_primeval_titan()
        case "Sakura-Tribe Elder":
            return clone.play_sakura_tribe_elder()
        case "Search for Tomorrow":
            return clone.play_search_for_tomorrow()
        case "Shefet Monitor":
            return clone.play_shefet_monitor()
        case "Sheltered Thicket":
            return clone.play_sheltered_thicket()
        case "Simian Spirit Guide":
            return clone.play_simian_spirit_guide()
        case "Sleight of Hand":
            return clone.play_sleight_of_hand()
        case "Stomping Ground":
            return clone.play_stomping_ground()
        case "Summoner's Pact":
            return clone.play_summoners_pact()
        case "Through the Breach":
            return clone.play_through_the_breach()
        case "Valakut, the Molten Pinnacle":
            return clone.play_valakut_the_molten_pinnacle()
        case "Wooded Foothills":
            return clone.play_wooded_foothills()
    }
    return []state{}
}

func get_cost(card string) string {
    switch card {
        case "Explore":
            return "1G"
        case "Primeval Titan":
            return "4GG"
        case "Sakura-Tribe Elder":
            return "1G"
        // Treat SFT as a 1-cost spell with kicker 2.
        case "Search for Tomorrow":
            return "G"
        case "Shefet Monitor":
            return "3G"
        case "Sleight of Hand":
            return "G"
        case "Summoner's Pact":
            return "4GG"
        case "Through the Breach":
            return "4R"
        default:
            return ""
    }
}

// ---------------------------------------------------------------------

func (s *state) pass_turn() []state {
    s.turn += 1
    s.lands = 1
    // Reset mana pool
    s.pool.empty()
    for _, card := range s.board {
        s.pool.add(land_output(card))
    }
    s.note("")
    turn := "[T" + strconv.Itoa(s.turn) + "] "
    if s.turn > 1 || !s.play {
        s.note(turn + "Hand: " + tally(s.hand) + ", drawing " + slug(s.deck[0]))
        s.draw(1)
    } else {
        s.note(turn + "Hand: " + tally(s.hand))
    }
    if len(s.board) > 0 { s.note(turn + "Board: " + tally(s.board)) }
    if len(s.exile) > 0 { s.note(turn + "Exile: " + tally(s.exile) + ", ticking down") }
    // Handle anything suspended
    for i, card := range s.exile {
        if card[0] == '.' { s.exile[i] = s.exile[i][1:] }
    }
    for _, card := range s.exile {
        if card[0] != '.' { s.play_suspended(card) }

    }
    return []state{*s}
}

func (s *state) play_suspended(card string) {
    s.note(slug(card) + " from exile")
    s.exile = remove(s.exile, card)
    switch card {
        case "Search for Tomorrow":
            s.board = append(s.board, "Forest")
            s.pool.add("G")
    }
}

// ---------------------------------------------------------------------

func (s *state) play_cinder_glade() []state {
    s.board = append(s.board, "Taiga")
    nbasics := 0
    for _, card := range s.board {
        if card == "Forest" || card == "Mountain" { nbasics++ }
    }
    s.unnote()
    if nbasics > 1 {
        s.note("Playing " + slug("Cinder Glade") + " untapped")
        s.pool.add("G")
    } else {
        s.note("Playing " + slug("Cinder Glade") + " tapped")
    }
    return []state{*s}
}

func (s *state) play_explore() []state {
    s.lands += 1
    s.unnote()
    s.note(slug("Explore") + ", drawing " + slug(s.deck[0]))
    s.draw(1)
    return []state{*s}
}

func (s *state) play_forest() []state {
    s.board = append(s.board, "Forest")
    s.pool.add("G")
    return []state{*s}
}

func (s *state) play_mountain() []state {
    s.board = append(s.board, "Mountain")
    s.pool.add("R")
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
    if s.pool.can_pay("2") {
        clone := s.clone()
        clone.pool.pay("2")
        clone.pool.add("G")
        clone.board = append(clone.board, "Forest")
        states = append(states, clone)
    }
    s.unnote()
    s.note("Suspending " + slug("Search for Tomorrow"))
    s.exile = append(s.exile, "..Search for Tomorrow")
    states = append(states, *s)
    return states
}

func (s *state) play_shefet_monitor() []state {
    s.board = append(s.board, "Forest")
    s.unnote()
    s.note("Cycling " + slug("Shefet Monitor") + ", drawing " + slug(s.deck[0]))
    s.draw(1)
    s.pool.add("G")
    return []state{*s}
}

func (s *state) play_sheltered_thicket() []state {
    states := []state{}
    if s.pool.can_pay("2") {
        clone := s.clone()
        // Un-play the land
        clone.lands += 1
        clone.pool.pay("2")
        clone.unnote()
        clone.note("Cycling " + slug("Sheltered Thicket") + ", drawing " + slug(clone.deck[0]))
        clone.draw(1)
        states = append(states, clone)
    }
    if s.lands >= 0 {
        s.board = append(s.board, "Taiga")
        states = append(states, *s)
    }
    return states
}

func (s *state) play_simian_spirit_guide() []state {
    s.pool.add("R")
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

func (s *state) play_stomping_ground() []state {
    s.board = append(s.board, "Taiga")
    s.pool.add("G")
    return []state{*s}
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

func (s *state) play_valakut_the_molten_pinnacle() []state {
    s.board = append(s.board, "Valakut, the Molten Pinnacle")
    return []state{*s}
}

func (s *state) play_wooded_foothills() []state {
    s.board = append(s.board, "Forest")
    s.pool.add("G")
    return []state{*s}
}

// ---------------------------------------------------------------------

func (s *state) note(line string) {
    s.Log = s.Log + "\n" + line
}

func (s *state) unnote() {
    lines := strings.Split(s.Log, "\n")
    s.Log = strings.Join(lines[:len(lines)-1], "\n")

}

func (s state) clone() state {
    // Deep copy the slices.
    s.deck = append([]string{}, s.deck...)
    s.hand = append([]string{}, s.hand...)
    s.board = append([]string{}, s.board...)
    s.exile = append([]string{}, s.exile...)
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
