
package valakut

import (
    "errors"
)

// ---------------------------------------------------------------------

func Simulate(name string) (game_state, error) {
    // Accept a deck list. Simulate a game and return the final state.
    deck, err := LoadDeck(name)
    if err != nil { return game_state{}, err }
    states := turn_zero(deck)
    return play_turns(states)
}

// ---------------------------------------------------------------------

func play_turns(ugs unique_game_states) (game_state, error) {
    var err error
    for turn := 1; turn < 7; turn++ {
        ugs, err = next_turn(ugs)
        if err != nil { return game_state{}, err }
        // Prefer states with fewer mulligans.
        done_state := game_state{}
        done_mulls := 3
        for _, gs := range ugs.iter() {
            if gs.done && (gs.mulls < done_mulls || gs.fast) {
                done_mulls = gs.mulls
                done_state = gs
            }
        }
        if done_state.turn > 0 { return done_state, nil }
    }
    return game_state{}, errors.New("Failed to finish")
}

// ---------------------------------------------------------------------

func next_turn(ugs unique_game_states) (unique_game_states, error) {
    // Take some states at turn N. Return states at turn N+1.
    ugs_old := ugs
    ugs_new := UniqueGameStates()
    gs, err := ugs_old.get()
    if err != nil { return UniqueGameStates(), err }
    current_turn := gs.turn
    for ugs_old.size() > 0 {
        gs_old, err := ugs_old.pop()
        if err != nil { return UniqueGameStates(), err }
        for _, gs := range gs_old.next() {
            if gs.turn > current_turn || gs.done {
                ugs_new.add(gs)
            } else {
                ugs_old.add(gs)
            }
        }
    }
    return ugs_new, nil
}

// ---------------------------------------------------------------------

func turn_zero(deck []string) unique_game_states {
    // Resolve all mulligans and return states ready to go.
    seven := GameState(deck)
//    sixes := seven.clone_mulligan()
//    fives := sixes[0].clone_mulligan()
    ugs := UniqueGameStates(seven)
//    ugs.add(sixes...)
//    ugs.add(fives...)
    for _, gs := range ugs.iter() { gs.pass_turn() }
    return ugs
}
