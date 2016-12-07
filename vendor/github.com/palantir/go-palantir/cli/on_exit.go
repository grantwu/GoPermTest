package cli

import (
	"fmt"
	"sort"
	"sync/atomic"
)

type OnExitID struct {
	id int32
}

type byID []OnExitID

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].id < a[j].id }

type OnExit interface {
	Register(f func()) OnExitID
	Unregister(id OnExitID)
	run()
}

func newOnExit() OnExit {
	return &onExit{
		functions: make(map[OnExitID]func()),
	}
}

type onExit struct {
	idCounter OnExitID
	functions map[OnExitID]func()
}

func (m *onExit) Register(f func()) OnExitID {
	id := OnExitID{
		id: atomic.AddInt32(&m.idCounter.id, 1),
	}
	m.functions[id] = f
	return id
}

func (m *onExit) Unregister(id OnExitID) {
	delete(m.functions, id)
}

func (m *onExit) run() {
	// sort keys in descending order (LIFO)
	keys := make([]OnExitID, 0, len(m.functions))
	for id := range m.functions {
		keys = append(keys, id)
	}
	sort.Sort(sort.Reverse(byID(keys)))

	for _, id := range keys {
		// invoke all registered onExit functions. Invoke in a manner that recovers from panics so that a
		// function panicking does not prevent the other functions from running.
		invokeAndRecover(m.functions[id])
	}
}

// Invokes the provided function. If the function panics, recovers and prints to console.
func invokeAndRecover(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in invokeAndRecover:", r)
		}
	}()
	f()
}
