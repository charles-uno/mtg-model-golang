
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

func play_turns(ss state_set) (state, error) {
    for turn := 1; turn < 7; turn++ {
        ss = next_turn(ss)
        // Prefer states with fewer mulligans.
        done_state := state{}
        done_mulls := 3
        for _, s := range ss.states {
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

func next_turn(ss state_set) state_set {
    // Take some states at turn N. Return states at turn N+1.
    ss_old := ss
    ss_new := set()
    current_turn := ss.get().turn
    for len(ss_old.states) > 0 {
        s_old := ss_old.pop()
        for _, s := range s_old.next() {
            if s.turn > current_turn || s.done {
                ss_new.add(s)
            } else {
                ss_old.add(s)
            }
        }
    }
    return ss_new
}

// ---------------------------------------------------------------------

func turn_zero(deck []string) state_set {
    // Resolve all mulligans and return states ready to go.
    seven := state{deck: deck}
    seven.flip()
    seven.shuffle()
    seven.draw(7)
    sixes := seven.clone_mulligan()
    fives := sixes[0].clone_mulligan()
    ss := set(seven, sixes[0], sixes[1], fives[0], fives[1])
    for _, s := range ss.states { s.pass_turn() }
    return ss
}
