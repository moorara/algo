package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultErrorFormat(t *testing.T) {
	tests := []struct {
		name           string
		errs           []error
		expectedString string
	}{
		{
			name:           "Nil",
			errs:           nil,
			expectedString: "",
		},
		{
			name:           "Zero",
			errs:           []error{},
			expectedString: "",
		},
		{
			name: "One",
			errs: []error{
				new(fooError),
			},
			expectedString: "error on foo\n",
		},
		{
			name: "Multiple",
			errs: []error{
				new(fooError),
				new(barError),
				new(bazError),
			},
			expectedString: "error on foo\nerror on bar\nerror on baz\nsomething failed\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := DefaultErrorFormat(tc.errs)

			assert.Equal(t, tc.expectedString, s)
		})
	}
}

func TestBulletErrorFormat(t *testing.T) {
	tests := []struct {
		name           string
		errs           []error
		expectedString string
	}{
		{
			name:           "Nil",
			errs:           nil,
			expectedString: "",
		},
		{
			name:           "Zero",
			errs:           []error{},
			expectedString: "",
		},
		{
			name: "One",
			errs: []error{
				new(fooError),
			},
			expectedString: "1 error occurred:\n\n  • error on foo\n",
		},
		{
			name: "Multiple",
			errs: []error{
				new(fooError),
				new(barError),
				new(bazError),
			},
			expectedString: "3 errors occurred:\n\n  • error on foo\n  • error on bar\n  • error on baz\n    something failed\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := BulletErrorFormat(tc.errs)

			assert.Equal(t, tc.expectedString, s)
		})
	}
}
