package event

import (
	"sort"
)

type Cancellable interface {
	Cancelled() bool
	Cancel()
	Resume()
}

type Event struct {
	Name string
	Data interface{}
}

type ListenerPriority int

const (
	LOWEST ListenerPriority = iota + 1
	LOW
	NORMAL
	HIGH
	HIGHEST
	ASAP
)

type Listener struct {
	IgnoreCancelled bool
	Async           bool
	Priority        ListenerPriority
	Function        func(event *Event)
}

var registered = make(map[string][]Listener)

func FireEvent(event *Event) {
	for _, f := range registered[event.Name] {
		var e interface{}
		e = *event

		switch t := e.(type) {
		case Cancellable:
			if t.Cancelled() && !f.IgnoreCancelled {
				continue
			}
		}

		if f.Async {
			go f.Function(event)
		} else {
			f.Function(event)
		}
	}
}

func RegisterListener(listener Listener, target string) {
	if value, ok := registered[target]; ok {
		registered[target] = append(value, listener)
	} else {
		registered[target] = []Listener{listener}
	}

	sort.Slice(registered[target], func(f, s int) bool {
		return registered[target][f].Priority > registered[target][s].Priority
	})
}
