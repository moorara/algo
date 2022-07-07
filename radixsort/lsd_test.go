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

		if !isSorted[string](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}

func TestLSDInt(t *testing.T) {
	tests := []struct {
		a []int
	}{
		{[]int{30, -20, 10, -40, 50}},
		{[]int{90, -80, 70, -60, 50, -40, 30, -20, 10}},
	}

	for _, tc := range tests {
		LSDInt(tc.a)

		if !isSorted[int](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}

func TestLSDUint(t *testing.T) {
	tests := []struct {
		a []uint
	}{
		{[]uint{30, 20, 10, 40, 50}},
		{[]uint{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		LSDUint(tc.a)

		if !isSorted[uint](tc.a) {
			t.Fatalf("%v is not sorted.", tc.a)
		}
	}
}
