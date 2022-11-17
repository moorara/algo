package automata

// stateFactory is used for keeping track of states when combining multiple finite automata.
type stateFactory struct {
	last   State
	states map[int]map[State]State
}

func newStateFactory() *stateFactory {
	return &stateFactory{
		last:   0,
		states: map[int]map[State]State{},
	}
}

func (f *stateFactory) StateFor(id int, s State) State {
	m, ok := f.states[id]
	if !ok {
		m = map[State]State{}
		f.states[id] = m
	}

	t, ok := m[s]
	if !ok {
		f.last++
		t = f.last
		m[s] = t
	}

	return t
}
