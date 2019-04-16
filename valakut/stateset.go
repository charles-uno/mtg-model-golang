
package valakut

// ---------------------------------------------------------------------

// Pretend to be a Python-style set of states.
type state_set struct {
    states map[string]state
}

func set(states ...state) state_set {
    ss := state_set{states: make(map[string]state)}
    for _, s := range states {
        ss.add(s)
    }
    return ss
}

func (ss *state_set) pop() state {
    for k, s := range ss.states { delete(ss.states, k); return s }
    return state{}
}

func (ss *state_set) get() state {
    for _, s := range ss.states { return s }
    return state{}
}

func (ss *state_set) add(s state) {
    ss.states[s.key()] = s
}
