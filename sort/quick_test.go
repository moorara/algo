package sort

import (
	"testing"

	"github.com/moorara/algo/common"
)

func TestSelect_int(t *testing.T) {
	tests := []struct {
		items         []int
		expectedItems []int
	}{
		{[]int{}, nil},
		{[]int{20, 10, 30}, []int{10, 20, 30}},
		{[]int{20, 10, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
		{[]int{20, 10, 30, 40, 50, 80, 60, 70, 90}, []int{10, 20, 30, 40, 50, 60, 70, 80, 90}},
	}

	for _, tc := range tests {
		for k := 0; k < len(tc.items); k++ {
			cmp := common.NewCompareFunc[int]()
			item := Select[int](tc.items, k, cmp)

			if item != tc.expectedItems[k] {
				t.Fatalf("expected selection: %d, actual selection: %d", tc.expectedItems[k], item)
			}
		}
	}
}

func TestQuick_int(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{[]int{}},
		{[]int{20, 10, 30}},
		{[]int{30, 20, 10, 40, 50}},
		{[]int{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		cmp := common.NewCompareFunc[int]()
		Quick[int](tc.items, cmp)

		if !isSorted(tc.items, cmp) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestQuick_string(t *testing.T) {
	tests := []struct {
		items []string
	}{
		{[]string{}},
		{[]string{"Milad", "Mona"}},
		{[]string{"Alice", "Bob", "Alex", "Jackie"}},
		{[]string{"Docker", "Kubernetes", "Go", "JavaScript", "Elixir", "React", "Redux", "Vue"}},
	}

	for _, tc := range tests {
		cmp := common.NewCompareFunc[string]()
		Quick[string](tc.items, cmp)

		if !isSorted(tc.items, cmp) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestQuick3Way_int(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{[]int{}},
		{[]int{20, 10, 10, 20, 30, 30, 30}},
		{[]int{30, 20, 30, 20, 10, 40, 40, 40, 50, 50}},
		{[]int{90, 10, 80, 20, 70, 30, 60, 40, 50, 50, 40, 60, 30, 70, 20, 80, 10, 90}},
	}

	for _, tc := range tests {
		cmp := common.NewCompareFunc[int]()
		Quick3Way[int](tc.items, cmp)

		if !isSorted(tc.items, cmp) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestQuick3Way_string(t *testing.T) {
	tests := []struct {
		items []string
	}{
		{[]string{}},
		{[]string{"Milad", "Mona", "Milad", "Mona"}},
		{[]string{"Alice", "Bob", "Alex", "Jackie", "Jackie", "Alex", "Bob", "Alice"}},
		{[]string{"Docker", "Kubernetes", "Docker", "Go", "JavaScript", "Go", "React", "Redux", "Vue", "Redux", "React"}},
	}

	for _, tc := range tests {
		cmp := common.NewCompareFunc[string]()
		Quick3Way[string](tc.items, cmp)

		if !isSorted(tc.items, cmp) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}
