package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestItemSetStringer(t *testing.T) {
	s := getTestLR0ItemSets()
	state := State(0)

	tests := []struct {
		name               string
		ss                 *itemSetStringer
		expectedSubstrings []string
	}{
		{
			name: "WithoutState",
			ss: &itemSetStringer{
				items: generic.Collect1(s[0].All()),
			},
			expectedSubstrings: []string{
				`┌────────────────┐`,
				`│ E′ → •E        │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
			},
		},
		{
			name: "WithState",
			ss: &itemSetStringer{
				state: &state,
				items: generic.Collect1(s[0].All()),
			},
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
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.ss.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, str, expectedSubstring)
			}
		})
	}
}

func TestItemSetCollectionStringer(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name               string
		cs                 *itemSetCollectionStringer
		expectedSubstrings []string
	}{
		{
			name: "OK",
			cs: &itemSetCollectionStringer{
				sets: []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			},
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
		t.Run(tc.name, func(t *testing.T) {
			str := tc.cs.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, str, expectedSubstring)
			}
		})
	}
}
