package grammar

// MockWriter is an implementation of io.Writer for testing purposes.
type MockWriter struct {
	WriteIndex int
	WriteMocks []WriteMock
}

type WriteMock struct {
	InBytes  []byte
	OutN     int
	OutError error
}

func (m *MockWriter) Write(b []byte) (int, error) {
	i := m.WriteIndex
	m.WriteIndex++
	m.WriteMocks[i].InBytes = b
	return m.WriteMocks[i].OutN, m.WriteMocks[i].OutError
}
