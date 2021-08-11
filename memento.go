package main

// state field in memento stores the state of originator
type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}
