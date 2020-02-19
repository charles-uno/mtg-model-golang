package mtgbf


import (
    "errors"
)


type game_state_manager struct {
    states map[string]game_state
}


func UniqueGameStates(states ...game_state) unique_game_states {
    ugs := unique_game_states{states: make(map[string]game_state)}
    for _, gs := range states {
        ugs.add(gs)
    }
    return ugs
}


func (ugs *unique_game_states) pop() (game_state, error) {
    for key, gs := range ugs.states {
        delete(ugs.states, key)
        return gs, nil
    }
    return game_state{}, errors.New("Can't pop from empty UGS")
}


func (ugs *unique_game_states) get() (game_state, error) {
    for _, gs := range ugs.states { return gs, nil }
    return game_state{}, errors.New("Can't get from empty UGS")
}


func (ugs *unique_game_states) add(states ...game_state) {
    for _, gs := range states {
        ugs.states[gs.key()] = gs
    }
}


func (ugs *unique_game_states) iter() []game_state {
    states := []game_state{}
    for _, gs := range ugs.states { states = append(states, gs) }
    return states
}


func (ugs *unique_game_states) size() int {
    return len(ugs.states)
}
