package sort

import "testing"

func TestSelectInt(t *testing.T) {
	tests := []struct {
		items         []interface{}
		expectedItems []interface{}
	}{
		{[]interface{}{}, nil},
		{[]interface{}{20, 10, 30}, []interface{}{10, 20, 30}},
		{[]interface{}{20, 10, 30, 40, 50}, []interface{}{10, 20, 30, 40, 50}},
		{[]interface{}{20, 10, 30, 40, 50, 80, 60, 70, 90}, []interface{}{10, 20, 30, 40, 50, 60, 70, 80, 90}},
	}

	for _, tc := range tests {
		for k := 0; k < len(tc.items); k++ {
			item := Select(tc.items, k, compareInt)

			if item != tc.expectedItems[k] {
				t.Fatalf("expected selection: %d, actual selection: %d", tc.expectedItems[k], item)
			}
		}
	}
}

func TestQuickInt(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{20, 10, 30}},
		{[]interface{}{30, 20, 10, 40, 50}},
		{[]interface{}{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		Quick(tc.items, compareInt)

		if !sorted(tc.items, compareInt) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestQuickString(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{"Milad", "Mona"}},
		{[]interface{}{"Alice", "Bob", "Alex", "Jackie"}},
		{[]interface{}{"Docker", "Kubernetes", "Go", "JavaScript", "Elixir", "React", "Redux", "Vue"}},
	}

	for _, tc := range tests {
		Quick(tc.items, compareString)

		if !sorted(tc.items, compareString) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestQuick3WayInt(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{20, 10, 10, 20, 30, 30, 30}},
		{[]interface{}{30, 20, 30, 20, 10, 40, 40, 40, 50, 50}},
		{[]interface{}{90, 10, 80, 20, 70, 30, 60, 40, 50, 50, 40, 60, 30, 70, 20, 80, 10, 90}},
	}

	for _, tc := range tests {
		Quick3Way(tc.items, compareInt)

		if !sorted(tc.items, compareInt) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestQuick3WayString(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{"Milad", "Mona", "Milad", "Mona"}},
		{[]interface{}{"Alice", "Bob", "Alex", "Jackie", "Jackie", "Alex", "Bob", "Alice"}},
		{[]interface{}{"Docker", "Kubernetes", "Docker", "Go", "JavaScript", "Go", "React", "Redux", "Vue", "Redux", "React"}},
	}

	for _, tc := range tests {
		Quick3Way(tc.items, compareString)

		if !sorted(tc.items, compareString) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}
