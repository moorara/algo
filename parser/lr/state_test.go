package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildStateMap(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name             string
		C                ItemSetCollection
		expectedStateMap StateMap
	}{
		{
			name:             "OK",
			C:                NewItemSetCollection(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]),
			expectedStateMap: []ItemSet{s[0], s[1], s[9], s[11], s[10], s[6], s[8], s[7], s[2], s[4], s[5], s[3]},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stateMap := BuildStateMap(tc.C)
			assert.Equal(t, tc.expectedStateMap, stateMap)
		})
	}
}

func TestStateMap_Find(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name          string
		m             StateMap
		I             ItemSet
		expectedState State
	}{
		{
			name:          "OK",
			m:             []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			I:             s[7],
			expectedState: State(7),
		},
		{
			name: "Error",
			m:    []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			I: NewItemSet(
				&Item0{Production: prods[1][0], Start: starts[1], Dot: 1}, // E′ → E•
			),
			expectedState: ErrState,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedState, tc.m.Find(tc.I))
	}
}

func TestStateMap_States(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name           string
		m              StateMap
		expectedStates []State
	}{
		{
			name:           "OK",
			m:              []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			expectedStates: []State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedStates, tc.m.States())
	}
}

func TestStateMap_String(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name               string
		m                  StateMap
		expectedSubstrings []string
	}{
		{
			name: "OK",
			m:    []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			expectedSubstrings: []string{
				`┌──────[0]───────┐`,
				`│ E′ → •E        │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[1]───────┐`,
				`│ E′ → E•        │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[2]───────┐`,
				`│ E → T•         │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[3]───────┐`,
				`│ T → F•         │`,
				`└────────────────┘`,
				`┌──────[4]───────┐`,
				`│ F → "("•E ")"  │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[5]───────┐`,
				`│ F → "id"•      │`,
				`└────────────────┘`,
				`┌──────[6]───────┐`,
				`│ E → E "+"•T    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[7]───────┐`,
				`│ T → T "*"•F    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`└────────────────┘`,
				`┌──────[8]───────┐`,
				`│ F → "(" E•")"  │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[9]───────┐`,
				`│ E → E "+" T•   │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[10]──────┐`,
				`│ T → T "*" F•   │`,
				`└────────────────┘`,
				`┌──────[11]──────┐`,
				`│ F → "(" E ")"• │`,
				`└────────────────┘`,
			},
		},
	}

	for _, tc := range tests {
		s := tc.m.String()

		for _, expectedSubstring := range tc.expectedSubstrings {
			assert.Contains(t, s, expectedSubstring)
		}
	}
}
