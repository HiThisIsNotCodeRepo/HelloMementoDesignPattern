package main

type originator struct {
	state string
}

// createMemento save its internal state into memento
func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

// restoreMemento restore from memento
func (e *originator) restoreMemento(m *memento) {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}
