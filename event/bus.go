package event

type Cancellable interface {
	Cancelled() bool
	Cancel()
	Resume()
}

type Event struct {
	Name string
	Data interface{}
}

type Listener struct {
	IgnoreCancelled bool
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

		f.Function(event)
	}
}

func RegisterListener(listener Listener, target string) {
	if value, ok := registered[target]; ok {
		value = append(value, listener)
	} else {
		registered[target] = []Listener{listener}
	}
}
