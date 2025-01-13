package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	var me *MultiError

	tests := []struct {
		name               string
		e                  error
		errs               []error
		expectedMultiError *MultiError
	}{
		{
			name:               "I",
			e:                  nil,
			errs:               nil,
			expectedMultiError: nil,
		},
		{
			name:               "II",
			e:                  nil,
			errs:               []error{},
			expectedMultiError: nil,
		},
		{
			name: "III",
			e:    nil,
			errs: []error{
				new(barError),
				new(bazError),
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "IV",
			e:    nil,
			errs: []error{
				&MultiError{
					errs: []error{
						new(barError),
						new(bazError),
					},
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "V",
			e:    new(fooError),
			errs: nil,
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
		},
		{
			name: "VI",
			e:    new(fooError),
			errs: []error{},
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
		},
		{
			name: "VII",
			e:    new(fooError),
			errs: []error{
				new(barError),
				new(bazError),
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "VIII",
			e:    new(fooError),
			errs: []error{
				&MultiError{
					errs: []error{
						new(barError),
						new(bazError),
					},
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name:               "IX",
			e:                  me,
			errs:               nil,
			expectedMultiError: nil,
		},
		{
			name:               "X",
			e:                  me,
			errs:               []error{},
			expectedMultiError: nil,
		},
		{
			name: "XI",
			e:    me,
			errs: []error{
				new(barError),
				new(bazError),
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "XII",
			e:    me,
			errs: []error{
				&MultiError{
					errs: []error{
						new(barError),
						new(bazError),
					},
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name:               "XIII",
			e:                  &MultiError{},
			errs:               nil,
			expectedMultiError: nil,
		},
		{
			name:               "XIV",
			e:                  &MultiError{},
			errs:               []error{},
			expectedMultiError: nil,
		},
		{
			name: "XV",
			e:    &MultiError{},
			errs: []error{
				new(barError),
				new(bazError),
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "XVI",
			e:    &MultiError{},
			errs: []error{
				&MultiError{
					errs: []error{
						new(barError),
						new(bazError),
					},
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "XVII",
			e: &MultiError{
				errs: []error{},
			},
			errs:               nil,
			expectedMultiError: nil,
		},
		{
			name: "XVIII",
			e: &MultiError{
				errs: []error{},
			},
			errs:               []error{},
			expectedMultiError: nil,
		},
		{
			name: "XIX",
			e: &MultiError{
				errs: []error{},
			},
			errs: []error{
				new(barError),
				new(bazError),
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "XX",
			e: &MultiError{
				errs: []error{},
			},
			errs: []error{
				&MultiError{
					errs: []error{
						new(barError),
						new(bazError),
					},
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "XXI",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			errs: nil,
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
		},
		{
			name: "XXII",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			errs: []error{},
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
		},
		{
			name: "XXIII",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			errs: []error{
				new(barError),
				new(bazError),
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
		},
		{
			name: "XXIV",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			errs: []error{
				&MultiError{
					errs: []error{
						new(barError),
						new(bazError),
					},
				},
			},
			expectedMultiError: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := Append(tc.e, tc.errs...)
			assert.Equal(t, tc.expectedMultiError, err)
		})
	}
}

func TestMultiError_ErrorOrNil(t *testing.T) {
	tests := []struct {
		name        string
		e           *MultiError
		expectError bool
	}{

		{
			name: "Nil",
			e: &MultiError{
				errs: nil,
			},
			expectError: false,
		},
		{
			name: "Zero",
			e: &MultiError{
				errs: []error{},
			},
			expectError: false,
		},
		{
			name: "One",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			expectError: true,
		},
		{
			name: "Multiple",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.e.ErrorOrNil()

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMultiError_Error(t *testing.T) {
	tests := []struct {
		name          string
		e             *MultiError
		expectedError string
	}{

		{
			name: "Nil",
			e: &MultiError{
				errs: nil,
			},
			expectedError: "",
		},
		{
			name: "Zero",
			e: &MultiError{
				errs: []error{},
			},
			expectedError: "",
		},
		{
			name: "One",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			expectedError: "error on foo\n",
		},
		{
			name: "Multiple",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
			expectedError: "error on foo\nerror on bar\nerror on baz\nsomething failed\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualError(t, tc.e, tc.expectedError)
		})
	}
}

func TestMultiError_Is(t *testing.T) {
	tests := []struct {
		name       string
		e          *MultiError
		target     error
		expectedIs bool
	}{
		{
			name: "False",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
				},
			},
			target:     &bazError{},
			expectedIs: false,
		},
		{
			name: "True",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
			target:     &bazError{},
			expectedIs: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			is := tc.e.Is(tc.target)
			assert.Equal(t, tc.expectedIs, is)
		})
	}
}

func TestMultiError_As(t *testing.T) {
	var bazErr *bazError

	tests := []struct {
		name       string
		e          *MultiError
		target     any
		expectedAs bool
	}{
		{
			name: "False",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
				},
			},
			target:     &bazErr,
			expectedAs: false,
		},
		{
			name: "True",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
			target:     &bazErr,
			expectedAs: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			as := tc.e.As(tc.target)
			assert.Equal(t, tc.expectedAs, as)
		})
	}
}

func TestMultiError_Unwrap(t *testing.T) {
	tests := []struct {
		name           string
		e              *MultiError
		expectedErrors []error
	}{
		{
			name:           "Nil",
			e:              &MultiError{},
			expectedErrors: nil,
		},
		{
			name: "Zero",
			e: &MultiError{
				errs: []error{},
			},
			expectedErrors: nil,
		},
		{
			name: "One",
			e: &MultiError{
				errs: []error{
					new(fooError),
				},
			},
			expectedErrors: []error{
				new(fooError),
			},
		},
		{
			name: "Multiple",
			e: &MultiError{
				errs: []error{
					new(fooError),
					new(barError),
					new(bazError),
				},
			},
			expectedErrors: []error{
				new(fooError),
				new(barError),
				new(bazError),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := tc.e.Unwrap()
			assert.Equal(t, tc.expectedErrors, errs)
		})
	}
}
