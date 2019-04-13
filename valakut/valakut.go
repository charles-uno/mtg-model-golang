
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
        for _, s := range states {
            if s.done { return s, nil }
        }
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
    return states_new
}

// ---------------------------------------------------------------------

func turn_zero(deck []string) []state {
    // Resolve all mulligans and return states ready to go.
    seven := state{deck: deck}
    seven.flip()
    seven.shuffle()
    seven.draw(7)
    six := seven.clone()
    six.shuffle()
    six.draw(6)
    five := six.clone()
    five.shuffle()
    sixb := six.clone()
    sixb.mill(1)
    five.draw(5)
    fiveb := five.clone()
    fiveb.mill(1)
    seven.pass_turn()
    six.pass_turn()
    sixb.pass_turn()
    five.pass_turn()
    fiveb.pass_turn()
    return []state{seven, six, sixb, five, fiveb}
}
