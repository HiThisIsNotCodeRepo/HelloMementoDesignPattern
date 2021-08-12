# HelloMementoDesignPattern

> [Source](https://golangbyexample.com/memento-design-pattern-go/)

## Core element

1. Originator: we are saving this object's state

```
// createMemento save its internal state into memento
func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

// restoreMemento restore from memento
func (e *originator) restoreMemento(m *memento) {
	e.state = m.getSavedState()
}
```

2. Memento: it stores Originator's state

```
// state field in memento stores the state of originator 
type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}
```

3. Caretaker: this object saves multiple mementos, and use index to return corresponding memento.

```
type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}
```