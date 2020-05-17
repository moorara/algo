package sort

import "testing"

func TestSelectionInt(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{20, 10, 30}},
		{[]interface{}{30, 20, 10, 40, 50}},
		{[]interface{}{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		Selection(tc.items, compareInt)

		if !sorted(tc.items, compareInt) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestSelectionString(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{"Milad", "Mona"}},
		{[]interface{}{"Alice", "Bob", "Alex", "Jackie"}},
		{[]interface{}{"Docker", "Kubernetes", "Go", "JavaScript", "Elixir", "React", "Redux", "Vue"}},
	}

	for _, tc := range tests {
		Selection(tc.items, compareString)

		if !sorted(tc.items, compareString) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}
