package errors

import "errors"

// Append adds one or more errors to an existing error and returns a MultiError instance.
//
//   - If the input error e is nil, a new instance of MultiError is created.
//   - If e is not already an instance of MultiError, it will be wrapped into a new MultiError instance.
//   - If any of the provided errors (errs) are instances of MultiError, they are flattened into the result.
func Append(e error, errs ...error) *MultiError {
	var me *MultiError

	if e == nil || e == me {
		me = new(MultiError)
	} else {
		switch e := e.(type) {
		case *MultiError:
			me = e
		default:
			// Wrap the existing error e into a new MultiError instance.
			me = new(MultiError)
			me.errs = make([]error, 0, len(errs)+1)
			me.errs = append(me.errs, e)
		}
	}

	// Iterate through the provided errors and flatten any instances of MultiError.
	for _, err := range errs {
		if err != nil {
			switch err := err.(type) {
			case *MultiError:
				me.errs = append(me.errs, err.errs...)
			default:
				me.errs = append(me.errs, err)
			}
		}
	}

	// If no errors were accumulated, return nil.
	// This ensures the behavior of MultiError stays consistent with a single error when no errors are present.
	if len(me.errs) == 0 {
		return nil
	}

	return me
}

// MultiError represents an error type that aggregates multiple errors.
// It is used to accumulate and manage multiple errors, wrapping them into a single error instance.
type MultiError struct {
	errs   []error
	Format ErrorFormat
}

// // ErrorOrNil returns an error if the MultiError instance contains any errors, or nil if it has none.
// This method is useful for ensuring that a valid error value is returned after accumulating errors,
// indicating whether errors are present or not.
func (e *MultiError) ErrorOrNil() error {
	if e == nil || len(e.errs) == 0 {
		return nil
	}

	return e
}

// Error implements the error interface for MultiError.
// It formats the accumulated errors into a single string representation.
// If a custom ErrorFormat function is provided, it will be used;
// otherwise, the default ErrorFormat (DefaultErrorFormat) will be applied.
func (e *MultiError) Error() string {
	if e.Format == nil {
		e.Format = DefaultErrorFormat
	}

	return e.Format(e.errs)
}

// Is checks if any of the errors in the MultiError matches the target error.
func (e *MultiError) Is(target error) bool {
	for _, err := range e.errs {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

// As finds the first error in the MultiError's errors that matches target,
// and if one is found, sets target to that error value and returns true.
// Otherwise, it returns false.
func (e *MultiError) As(target any) bool {
	for _, err := range e.errs {
		if errors.As(err, target) {
			return true
		}
	}

	return false
}

// Unwrap implements the unwrap interface for MultiError.
// It returns the slice of accumulated errors wrapped in the MultiError instance.
// If there are no errors, it returns nil, indicating that e does not wrap any error.
func (e *MultiError) Unwrap() []error {
	if len(e.errs) == 0 {
		return nil
	}

	return e.errs
}
