
package valakut

import (
    "errors"
)

// ---------------------------------------------------------------------

func Simulate(name string) (state, error) {
    // Accept a deck list. Simulate a game and return the final state.
    deck, err := load(name)
    if err != nil { return state{}, err }
    states := turn_zero(deck)
    return play_turns(states)
}

// ---------------------------------------------------------------------

func play_turns(states []state) (state, error) {
    for turn := 1; turn < 7; turn++ {
        states = next_turn(states)
        // Prefer states with fewer mulligans.
        done_state := state{}
        done_mulls := 3
        for _, s := range states {
            if s.done && s.mulls < done_mulls {
                done_mulls = s.mulls
                done_state = s
            }
        }
        if done_state.turn > 0 { return done_state, nil }
    }
    return state{}, errors.New("Failed to finish")
}

// ---------------------------------------------------------------------

func next_turn(states []state) []state {
    // Take some states at turn N. Return states at turn N+1.
    states_old := states
    states_new := []state{}
    for len(states_old) > 0 {
        for _, s := range states_old[0].next() {
            if s.turn > states_old[0].turn || s.done {
                states_new = append(states_new, s)
            } else {
                states_old = append(states_old, s)
            }
        }
        states_old = states_old[1:]
    }
    return unique_states(states_new)
}

// ---------------------------------------------------------------------

func unique_states(states []state) []state {
    // Use map keys to get rid of duplicates. There are a lot of them.
    tracker := make(map[string]state)
    for _, s := range states {
        tracker[s.id()] = s
    }
    new_states := []state{}
    for _, s := range tracker {
        new_states = append(new_states, s)
    }
    return new_states
}

// ---------------------------------------------------------------------

func turn_zero(deck []string) []state {
    // Resolve all mulligans and return states ready to go.
    seven := state{deck: deck}
    seven.flip()
    seven.shuffle()
    seven.draw(7)
    sixes := seven.clone_mulligan()
    fives := sixes[0].clone_mulligan()
    states := []state{seven, sixes[0], sixes[1], fives[0], fives[1]}
    for _, s := range states { s.pass_turn() }
    return states
}
