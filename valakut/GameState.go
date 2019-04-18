
package valakut

import (
    "strings"
    "strconv"
)

// ---------------------------------------------------------------------

type game_state struct {
    board []string
    deck []string
    done bool
    exile []string
    fast bool
    hand []string
    lands int
    Log string
    mulls int
    play bool
    pool mana
    turn int
}

func GameState(deck []string) game_state {
    gs := game_state{deck: deck}
    gs.flip()
    gs.shuffle()
    gs.draw(7)
    return gs
}

func (gs *game_state) clone_mulligan() []game_state {
    ctop := gs.clone()
    ctop.mulls += 1
    ctop.shuffle()
    ctop.draw(7 - ctop.mulls)
    cbot := ctop.clone()
    cbot.mill(1)
    return []game_state{ctop, cbot}
}

// ---------------------------------------------------------------------

func (gs *game_state) key() string {
    // We don't care about order for hand, exile, or board. But we do
    // care for deck.
    return strings.Join(
        []string{
            tally(gs.hand),
            tally(gs.exile),
            tally(gs.board),
            gs.pool.show(),
            strconv.Itoa(gs.lands),
            strconv.FormatBool(gs.done),
            strings.Join(gs.deck, " "),
        },
        ";",
    )
}

func (gs *game_state) Line() string {
    return strings.Join(
        []string{
            strconv.Itoa(gs.turn),
            b2i2a(gs.play),
            strconv.Itoa(gs.mulls),
            b2i2a(gs.fast),
        },
        ",",
    )
}

// ---------------------------------------------------------------------

func (gs *game_state) next() []game_state {
    states := []game_state{}
    for _, card := range unique_strings(gs.hand) {
        states = append(states, gs.clone_play(card)...)
    }
    states = append(states, gs.clone_pass_turn()...)
    return states
}

// ---------------------------------------------------------------------

func (gs *game_state) clone_pass_turn() []game_state {
    clone := gs.clone()
    return clone.pass_turn()
}

func (gs *game_state) clone_play(card string) []game_state {
    if !gs.in_hand(card) { return []game_state{} }
    // Don't bail for spell lands like Sheltered Thicket.
    if !is_spell(card) && gs.lands == 0 { return []game_state{} }
    cost := get_cost(card)
    if !gs.pool.can_pay(cost) { return []game_state{} }
    clone := gs.clone()
    clone.note(slug(card))
    clone.hand = remove(clone.hand, card)
    // Careful -- we may accidentally over-play from Sheltered Thicket.
    if is_land(card) { clone.lands -= 1 } else { clone.pool.pay(cost) }
    switch card {
        case "Arboreal Grazer":
            return clone.play_arboreal_grazer()
        case "Cinder Glade":
            return clone.play_cinder_glade()
        case "Deadshot Minotaur":
            return clone.play_deadshot_minotaur()
        case "Desperate Ritual":
            return clone.play_desperate_ritual()
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
    return []game_state{}
}

// ---------------------------------------------------------------------

func (gs *game_state) pass_turn() []game_state {
    gs.turn += 1
    gs.lands = 1
    // Reset mana pool
    gs.pool.empty()
    for _, card := range gs.board {
        gs.pool.add(land_output(card))
    }
    gs.note("")
    turn := "[T" + strconv.Itoa(gs.turn) + "] "
    if gs.turn > 1 || !gs.play {
        gs.note(turn + "Hand: " + tally(gs.hand) + ", drawing " + slug(gs.deck[0]))
        gs.draw(1)
    } else {
        gs.note(turn + "Hand: " + tally(gs.hand))
    }
    if len(gs.board) > 0 { gs.note(turn + "Board: " + tally(gs.board)) }
    if len(gs.exile) > 0 { gs.note(turn + "Exile: " + tally(gs.exile) + ", ticking down") }
    // Handle anything suspended
    for i, card := range gs.exile {
        if card[0] == '.' { gs.exile[i] = gs.exile[i][1:] }
    }
    for _, card := range gs.exile {
        if card[0] != '.' { gs.play_suspended(card) }

    }
    return []game_state{*gs}
}

func (gs *game_state) play_suspended(card string) {
    gs.note(slug(card) + " from exile")
    gs.exile = remove(gs.exile, card)
    switch card {
        case "Search for Tomorrow":
            gs.board = append(gs.board, "Forest")
            gs.pool.add("G")
    }
}

// ---------------------------------------------------------------------

func (gs *game_state) play_arboreal_grazer() []game_state {
    // We put the land in tapped...
    states := []game_state{}
    for _, card := range gs.hand {
        if is_land(card) {
            clone := gs.clone()
            clone.hand = remove(clone.hand, card)
            clone.board = append(clone.board, card)
            clone.unnote()
            clone.note(slug("Arboreal Grazer") + ", " + slug(card))
            states = append(states, clone)
        }
    }
    return states
}

func (gs *game_state) play_cinder_glade() []game_state {
    gs.board = append(gs.board, "Taiga")
    nbasics := 0
    for _, card := range gs.board {
        if card == "Forest" || card == "Mountain" { nbasics++ }
    }
    gs.unnote()
    if nbasics > 1 {
        gs.note(slug("Cinder Glade") + " (untapped)")
        gs.pool.add("G")
    } else {
        gs.note(slug("Cinder Glade") + " (tapped)")
    }
    return []game_state{*gs}
}

func (gs *game_state) play_deadshot_minotaur() []game_state {
    gs.unnote()
    gs.note("Cycling " + slug("Deadshot Minotaur") + ", drawing " + slug(gs.deck[0]))
    gs.draw(1)
    return []game_state{*gs}
}

func (gs *game_state) play_desperate_ritual() []game_state {
    gs.pool.add("RRR")
    return []game_state{*gs}
}

func (gs *game_state) play_explore() []game_state {
    gs.lands += 1
    gs.unnote()
    gs.note(slug("Explore") + ", drawing " + slug(gs.deck[0]))
    gs.draw(1)
    return []game_state{*gs}
}

func (gs *game_state) play_forest() []game_state {
    gs.board = append(gs.board, "Forest")
    gs.pool.add("G")
    return []game_state{*gs}
}

func (gs *game_state) play_mountain() []game_state {
    gs.board = append(gs.board, "Mountain")
    gs.pool.add("R")
    return []game_state{*gs}
}

func (gs *game_state) play_primeval_titan() []game_state {
    gs.done = true
    return []game_state{*gs}
}

func (gs *game_state) play_sakura_tribe_elder() []game_state {
    gs.board = append(gs.board, "Forest")
    return []game_state{*gs}
}

func (gs *game_state) play_search_for_tomorrow() []game_state {
    // Treat suspending as the default case, kicker 2 to get it now.
    states := []game_state{}
    if gs.pool.can_pay("2") {
        clone := gs.clone()
        clone.pool.pay("2")
        clone.pool.add("G")
        clone.board = append(clone.board, "Forest")
        states = append(states, clone)
    }
    gs.unnote()
    gs.note("Suspending " + slug("Search for Tomorrow"))
    gs.exile = append(gs.exile, "..Search for Tomorrow")
    states = append(states, *gs)
    return states
}

func (gs *game_state) play_shefet_monitor() []game_state {
    gs.board = append(gs.board, "Forest")
    gs.unnote()
    gs.note("Cycling " + slug("Shefet Monitor") + ", drawing " + slug(gs.deck[0]))
    gs.draw(1)
    gs.pool.add("G")
    return []game_state{*gs}
}

func (gs *game_state) play_sheltered_thicket() []game_state {
    states := []game_state{}
    if gs.pool.can_pay("2") {
        clone := gs.clone()
        // Un-play the land
        clone.lands += 1
        clone.pool.pay("2")
        clone.unnote()
        clone.note("Cycling " + slug("Sheltered Thicket") + ", drawing " + slug(clone.deck[0]))
        clone.draw(1)
        states = append(states, clone)
    }
    if gs.lands >= 0 {
        gs.board = append(gs.board, "Taiga")
        states = append(states, *gs)
    }
    return states
}

func (gs *game_state) play_simian_spirit_guide() []game_state {
    gs.pool.add("R")
    return []game_state{*gs}
}

func (gs *game_state) play_sleight_of_hand() []game_state {
    choices := gs.deck[:2]
    gs.deck = gs.deck[2:]
    clone := gs.clone()
    gs.note("Taking " + slug(choices[0]) + " over " + slug(choices[1]))
    gs.hand = append(gs.hand, choices[0])
    gs.deck = append(gs.deck, choices[1])
    clone.note("Taking " + slug(choices[1]) + " over " + slug(choices[0]))
    clone.hand = append(clone.hand, choices[0])
    clone.deck = append(clone.deck, choices[1])
    return []game_state{*gs, clone}
}

func (gs *game_state) play_stomping_ground() []game_state {
    gs.board = append(gs.board, "Taiga")
    gs.pool.add("G")
    return []game_state{*gs}
}

func (gs *game_state) play_summoners_pact() []game_state {
    return gs.play_primeval_titan()
}

func (gs *game_state) play_through_the_breach() []game_state {
    if !gs.in_hand("Primeval Titan") && !gs.in_hand("Summoner's Pact") {
        return []game_state{}
    }
    gs.done = true
    gs.fast = true
    return []game_state{*gs}
}

func (gs *game_state) play_valakut_the_molten_pinnacle() []game_state {
    gs.board = append(gs.board, "Valakut, the Molten Pinnacle")
    return []game_state{*gs}
}

func (gs *game_state) play_wooded_foothills() []game_state {
    gs.board = append(gs.board, "Forest")
    gs.pool.add("G")
    return []game_state{*gs}
}

// ---------------------------------------------------------------------

func (gs *game_state) note(line string) {
    gs.Log = gs.Log + "\n" + line
}

func (gs *game_state) unnote() {
    lines := strings.Split(gs.Log, "\n")
    gs.Log = strings.Join(lines[:len(lines)-1], "\n")

}

func (gs game_state) clone() game_state {
    // Deep copy the slices.
    gs.deck = append([]string{}, gs.deck...)
    gs.hand = append([]string{}, gs.hand...)
    gs.board = append([]string{}, gs.board...)
    gs.exile = append([]string{}, gs.exile...)
    return gs
}

func (gs *game_state) flip() {
    gs.play = flip()
    if gs.play { gs.note("On the play") } else { gs.note("On the draw") }
}

func (gs *game_state) shuffle() {
    gs.note("Shuffling")
    gs.deck, gs.hand = shuffled(append(gs.deck, gs.hand...)), []string{}
}

func (gs *game_state) draw(n int) {
    gs.deck, gs.hand = gs.deck[n:], append(gs.hand, gs.deck[:n]...)
}

func (gs *game_state) mill(n int) {
    gs.note("Milling " + tally(gs.deck[:n]))
    gs.deck = append(gs.deck[n:], gs.deck[:n]...)
}

func (gs *game_state) in_hand(card string) bool {
    for _, c := range gs.hand {
        if c == card { return true }
    }
    return false
}
