package errors

import "sync"

// Group is used for aggregating multiple errors from multiple goroutines.
// It provides thread-safe methods to run functions concurrently and collect any errors they return.
type Group struct {
	wg  sync.WaitGroup
	mu  sync.Mutex
	err *MultiError
}

// Go starts a new goroutine to execute the given function f.
// If f returns an error, it will be safely added to the aggregated errors.
func (e *Group) Go(f func() error) {
	e.wg.Add(1)

	go func() {
		defer e.wg.Done()

		if err := f(); err != nil {
			e.mu.Lock()
			e.err = Append(e.err, err)
			e.mu.Unlock()
		}
	}()
}

// Wait blocks until all goroutines started with the Go method have finished.
// It returns the aggregated errors from all goroutines.
// If no errors were returned, it returns nil.
func (e *Group) Wait() *MultiError {
	e.wg.Wait()
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.err
}
