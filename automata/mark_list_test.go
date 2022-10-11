package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var eqFunc = func(r, s rune) bool {
	return r == s
}

func TestNewMarkList(t *testing.T) {
	list := newMarkList[rune](eqFunc)
	assert.NotNil(t, list)
}

func TestMarkList_Values(t *testing.T) {
	tests := []struct {
		name         string
		list         *markList[rune]
		expectedVals []rune
	}{
		{
			name: "OK",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', true},
					{'c', true},
					{'d', false},
					{'e', false},
					{'f', false},
				},
			},
			expectedVals: []rune{'a', 'b', 'c', 'd', 'e', 'f'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			vals := tc.list.Values()
			assert.Equal(t, tc.expectedVals, vals)
		})
	}
}

func TestMarkList_AddUnmarked(t *testing.T) {
	tests := []struct {
		name          string
		list          *markList[rune]
		val           rune
		expectedIndex int
	}{
		{
			name: "OK",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', false},
				},
			},
			val:           'c',
			expectedIndex: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			index := tc.list.AddUnmarked(tc.val)
			assert.Equal(t, tc.expectedIndex, index)
		})
	}
}

func TestMarkList_GetUnmarked(t *testing.T) {
	tests := []struct {
		name          string
		list          *markList[rune]
		expectedVal   rune
		expectedIndex int
	}{
		{
			name: "WithUnmarked",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', false},
				},
			},
			expectedVal:   'b',
			expectedIndex: 1,
		},
		{
			name: "WithoutUnmarked",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', true},
				},
			},
			expectedVal:   0,
			expectedIndex: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val, index := tc.list.GetUnmarked()
			assert.Equal(t, tc.expectedVal, val)
			assert.Equal(t, tc.expectedIndex, index)
		})
	}
}

func TestMarkList_Contains(t *testing.T) {
	tests := []struct {
		name          string
		list          *markList[rune]
		val           rune
		expectedIndex int
	}{
		{
			name: "Yes",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', false},
				},
			},
			val:           'b',
			expectedIndex: 1,
		},
		{
			name: "No",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', false},
				},
			},
			val:           'c',
			expectedIndex: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			index := tc.list.Contains(tc.val)
			assert.Equal(t, tc.expectedIndex, index)
		})
	}
}

func TestMarkList_MarkByIndex(t *testing.T) {
	tests := []struct {
		name  string
		list  *markList[rune]
		index int
	}{
		{
			name: "OK",
			list: &markList[rune]{
				eq: eqFunc,
				list: []markEntry[rune]{
					{'a', true},
					{'b', false},
				},
			},
			index: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.MarkByIndex(tc.index)
			_, i := tc.list.GetUnmarked()
			assert.Equal(t, -1, i)
		})
	}
}
