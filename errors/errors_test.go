package errors

type fooError struct{}

func (e *fooError) Error() string {
	return "error on foo"
}

type barError struct{}

func (e *barError) Error() string {
	return "error on bar"
}

type bazError struct{}

func (e *bazError) Error() string {
	return "error on baz\nsomething failed"
}
