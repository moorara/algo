package sort

import (
	"testing"

	"github.com/moorara/algo/generic"
)

func TestSelection_int(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{[]int{}},
		{[]int{20, 10, 30}},
		{[]int{30, 20, 10, 40, 50}},
		{[]int{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		cmp := generic.NewCompareFunc[int]()
		Selection[int](tc.items, cmp)

		if !isSorted(tc.items, cmp) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestSelection_string(t *testing.T) {
	tests := []struct {
		items []string
	}{
		{[]string{}},
		{[]string{"Milad", "Mona"}},
		{[]string{"Alice", "Bob", "Alex", "Jackie"}},
		{[]string{"Docker", "Kubernetes", "Go", "JavaScript", "Elixir", "React", "Redux", "Vue"}},
	}

	for _, tc := range tests {
		cmp := generic.NewCompareFunc[string]()
		Selection[string](tc.items, cmp)

		if !isSorted(tc.items, cmp) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}
