package main

import "fmt"

func main() {
	aCaretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}

	anOriginator := &originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", anOriginator.getState())
	aCaretaker.addMemento(anOriginator.createMemento())

	anOriginator.setState("B")
	fmt.Printf("Originator Current State: %s\n", anOriginator.getState())

	aCaretaker.addMemento(anOriginator.createMemento())
	anOriginator.setState("C")

	fmt.Printf("Originator Current State: %s\n", anOriginator.getState())
	aCaretaker.addMemento(anOriginator.createMemento())

	anOriginator.restoreMemento(aCaretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", anOriginator.getState())

	anOriginator.restoreMemento(aCaretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", anOriginator.getState())
}
