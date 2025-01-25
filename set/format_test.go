package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultStringFormat(t *testing.T) {
	tests := []struct {
		name           string
		members        []int
		expectedString string
	}{
		{
			name:           "Nil",
			members:        nil,
			expectedString: "{}",
		},
		{
			name:           "Zero",
			members:        []int{},
			expectedString: "{}",
		},
		{
			name:           "One",
			members:        []int{2},
			expectedString: "{2}",
		},
		{
			name:           "Multiple",
			members:        []int{2, 4, 8, 16, 32, 64},
			expectedString: "{2, 4, 8, 16, 32, 64}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := defaultStringFormat(tc.members)

			assert.Equal(t, tc.expectedString, s)
		})
	}
}
