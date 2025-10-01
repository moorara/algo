package automata

// generateStatePermutations generates all permutations of a sequence of states using recursion and backtracking.
// Each permutation is passed to the provided yield function.
func generateStatePermutations(states []State, start, end int, yield func([]State) bool) bool {
	if start == end {
		return yield(states)
	}

	for i := start; i <= end; i++ {
		states[start], states[i] = states[i], states[start]
		cont := generateStatePermutations(states, start+1, end, yield)
		states[start], states[i] = states[i], states[start]

		if !cont {
			return false
		}
	}

	return true
}

// stateManager is used for keeping track of states when combining multiple automata.
type stateManager struct {
	last   State
	states map[int]map[State]State
}

func newStateManager(last State) *stateManager {
	return &stateManager{
		last:   last,
		states: map[int]map[State]State{},
	}
}

func (m *stateManager) GetOrCreateState(id int, s State) State {
	if _, ok := m.states[id]; !ok {
		m.states[id] = make(map[State]State)
	}

	if t, ok := m.states[id][s]; ok {
		return t
	}

	m.last++
	m.states[id][s] = m.last
	return m.last
}
