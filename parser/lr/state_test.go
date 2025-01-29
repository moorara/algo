package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildStateMap(t *testing.T) {
	tests := []struct {
		name             string
		C                ItemSetCollection
		expectedStateMap StateMap
	}{
		{
			name: "OK",
			C: NewItemSetCollection(
				LR0ItemSets[0],
				LR0ItemSets[1],
				LR0ItemSets[2],
				LR0ItemSets[3],
				LR0ItemSets[4],
				LR0ItemSets[5],
				LR0ItemSets[6],
				LR0ItemSets[7],
				LR0ItemSets[8],
				LR0ItemSets[9],
				LR0ItemSets[10],
				LR0ItemSets[11],
			),
			expectedStateMap: statemaps[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := BuildStateMap(tc.C)

			assert.Equal(t, tc.expectedStateMap, m)
		})
	}
}

func TestStateMap_Item(t *testing.T) {
	tests := []struct {
		name         string
		m            StateMap
		s            State
		i            int
		expectedItem Item
	}{
		{
			name:         "OK",
			m:            statemaps[0],
			s:            9,
			i:            3,
			expectedItem: &Item0{Production: prods[2][5], Start: starts[2], Dot: 0},
		},
	}

	for _, tc := range tests {
		item := tc.m.Item(tc.s, tc.i)

		assert.True(t, item.Equal(tc.expectedItem))
	}
}

func TestStateMap_ItemSet(t *testing.T) {
	tests := []struct {
		name            string
		m               StateMap
		s               State
		expectedItemSet ItemSet
	}{
		{
			name:            "OK",
			m:               statemaps[0],
			s:               5,
			expectedItemSet: LR0ItemSets[6],
		},
	}

	for _, tc := range tests {
		itemSet := tc.m.ItemSet(tc.s)

		assert.True(t, itemSet.Equal(tc.expectedItemSet))
	}
}

func TestStateMap_FindItem(t *testing.T) {
	tests := []struct {
		name          string
		m             StateMap
		s             State
		item          Item
		expectedIndex int
	}{
		{
			name:          "Found",
			m:             statemaps[0],
			s:             State(7),
			item:          &Item0{Production: prods[2][3], Start: starts[2], Dot: 2}, // T → T *•F
			expectedIndex: 0,
		},
		{
			name:          "NotFound",
			m:             statemaps[0],
			s:             State(7),
			item:          &Item0{Production: prods[2][5], Start: starts[2], Dot: 1}, // F → (•E ),
			expectedIndex: -1,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedIndex, tc.m.FindItem(tc.s, tc.item))
	}
}

func TestStateMap_FindItemSet(t *testing.T) {
	tests := []struct {
		name          string
		m             StateMap
		I             ItemSet
		expectedState State
	}{
		{
			name:          "Found",
			m:             statemaps[0],
			I:             LR0ItemSets[2],
			expectedState: State(8),
		},
		{
			name:          "NotFound",
			m:             statemaps[0],
			I:             NewItemSet(),
			expectedState: ErrState,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedState, tc.m.FindItemSet(tc.I))
	}
}

func TestStateMap_States(t *testing.T) {
	tests := []struct {
		name           string
		m              StateMap
		expectedStates []State
	}{
		{
			name:           "OK",
			m:              statemaps[0],
			expectedStates: []State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedStates, tc.m.States())
	}
}

func TestStateMap_All(t *testing.T) {
	tests := []struct {
		name           string
		m              StateMap
		expectedYields []ItemSet
	}{
		{
			name: "OK",
			m:    statemaps[0],
			expectedYields: []ItemSet{
				LR0ItemSets[0],
				LR0ItemSets[1],
				LR0ItemSets[9],
				LR0ItemSets[11],
				LR0ItemSets[10],
				LR0ItemSets[6],
				LR0ItemSets[8],
				LR0ItemSets[7],
				LR0ItemSets[2],
				LR0ItemSets[4],
				LR0ItemSets[5],
				LR0ItemSets[3],
			},
		},
	}

	for _, tc := range tests {
		for s, I := range tc.m.All() {
			assert.True(t, I.Equal(tc.expectedYields[s]))
		}
	}
}

func TestStateMap_String(t *testing.T) {
	tests := []struct {
		name               string
		m                  StateMap
		expectedSubstrings []string
	}{
		{
			name: "OK",
			m:    statemaps[0],
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
				`│ E → E "+" T•   │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[3]───────┐`,
				`│ F → "(" E ")"• │`,
				`└────────────────┘`,
				`┌──────[4]───────┐`,
				`│ T → T "*" F•   │`,
				`└────────────────┘`,
				`┌──────[5]───────┐`,
				`│ E → E "+"•T    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[6]───────┐`,
				`│ F → "(" E•")"  │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[7]───────┐`,
				`│ T → T "*"•F    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`└────────────────┘`,
				`┌──────[8]───────┐`,
				`│ E → T•         │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[9]───────┐`,
				`│ F → "("•E ")"  │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[10]──────┐`,
				`│ F → "id"•      │`,
				`└────────────────┘`,
				`┌──────[11]──────┐`,
				`│ T → F•         │`,
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
