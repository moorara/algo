package errors_test

import (
	"fmt"

	"github.com/moorara/algo/errors"
)

func ExampleAppend() {
	var err error
	err = errors.Append(err, fmt.Errorf("error on foo: %d", 1))
	err = errors.Append(err, fmt.Errorf("error on bar: %d", 2))
	err = errors.Append(err, fmt.Errorf("error on baz: %d", 3))

	fmt.Println(err)
}

func ExampleAppend_withCustomFormat() {
	err := &errors.MultiError{
		Format: errors.BulletErrorFormat,
	}

	err = errors.Append(err, fmt.Errorf("error on foo: %d", 1))
	err = errors.Append(err, fmt.Errorf("error on bar: %d", 2))
	err = errors.Append(err, fmt.Errorf("error on baz: %d", 3))

	fmt.Println(err)
}

func ExampleMultiError_ErrorOrNil() {
	err := new(errors.MultiError)
	err = err.ErrorOrNil()

	fmt.Println(err)
}

func ExampleMultiError_Unwrap() {
	var err *errors.MultiError
	err = errors.Append(err, fmt.Errorf("error on foo: %d", 1))
	err = errors.Append(err, fmt.Errorf("error on bar: %d", 2))
	err = errors.Append(err, fmt.Errorf("error on baz: %d", 3))

	for _, err := range err.Unwrap() {
		fmt.Println(err)
	}
}
