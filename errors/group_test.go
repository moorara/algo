package errors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroupError(t *testing.T) {
	tests := []struct {
		name               string
		funcs              []func() error
		expectedMultiError *MultiError
	}{
		{
			name: "WithoutErrors",
			funcs: []func() error{
				func() error {
					time.Sleep(20 * time.Millisecond)
					return nil
				},
				func() error {
					time.Sleep(10 * time.Millisecond)
					return nil
				},
				func() error {
					time.Sleep(30 * time.Millisecond)
					return nil
				},
			},
			expectedMultiError: nil,
		},
		{
			name: "WithErrors",
			funcs: []func() error{
				func() error {
					time.Sleep(20 * time.Millisecond)
					return new(fooError)
				},
				func() error {
					time.Sleep(10 * time.Millisecond)
					return new(barError)
				},
				func() error {
					time.Sleep(30 * time.Millisecond)
					return new(bazError)
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(fooError),
					new(bazError),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := new(Group)

			for _, f := range tc.funcs {
				g.Go(f)
			}

			err := g.Wait()
			assert.Equal(t, tc.expectedMultiError, err)
		})
	}
}
