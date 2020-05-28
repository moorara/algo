package radixsort

import "testing"

func TestLSDString(t *testing.T) {
	tests := []struct {
		a []string
		w int
	}{
		{
			a: []string{"BBB", "BBA", "BAB", "BAA", "ABB", "ABA", "AAB", "AAA"},
			w: 3,
		},
		{
			a: []string{"4PGC938", "2IYE230", "3CIO720", "1ICK750", "1OHV845", "4JZY524", "1ICK750", "3CIO720", "1OHV845", "1OHV845", "2RLA629", "2RLA629", "3ATW723"},
			w: 7,
		},
	}

	for _, tc := range tests {
		LSDString(tc.a, tc.w)

		if !isSortedString(tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}
