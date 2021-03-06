package sort

import (
	"testing"

	"github.com/moorara/algo/compare"
)

func TestMergeInt(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{20, 10, 30}},
		{[]interface{}{30, 20, 10, 40, 50}},
		{[]interface{}{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		Merge(tc.items, compare.Int)

		if !sorted(tc.items, compare.Int) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestMergeString(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{"Milad", "Mona"}},
		{[]interface{}{"Alice", "Bob", "Alex", "Jackie"}},
		{[]interface{}{"Docker", "Kubernetes", "Go", "JavaScript", "Elixir", "React", "Redux", "Vue"}},
	}

	for _, tc := range tests {
		Merge(tc.items, compare.String)

		if !sorted(tc.items, compare.String) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestMergeRecInt(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{20, 10, 30}},
		{[]interface{}{30, 20, 10, 40, 50}},
		{[]interface{}{90, 80, 70, 60, 50, 40, 30, 20, 10}},
	}

	for _, tc := range tests {
		MergeRec(tc.items, compare.Int)

		if !sorted(tc.items, compare.Int) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}

func TestMergeRecString(t *testing.T) {
	tests := []struct {
		items []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{"Milad", "Mona"}},
		{[]interface{}{"Alice", "Bob", "Alex", "Jackie"}},
		{[]interface{}{"Docker", "Kubernetes", "Go", "JavaScript", "Elixir", "React", "Redux", "Vue"}},
	}

	for _, tc := range tests {
		MergeRec(tc.items, compare.String)

		if !sorted(tc.items, compare.String) {
			t.Fatalf("%v is not sorted.", tc.items)
		}
	}
}
