package channel

import (
	"time"
)

//type SafeChannel interface {
//	WaitForClose()
//	Closed() bool
//}

type SafeFunctionChannel struct {
	channel chan func()
	closed bool
}

type SafeStringChannel struct {
	channel chan string
	closed bool
}

// TODO: Look into if its possible to D.R.Y. this code (probably a post generics release thing)

func NewSafeStringChannel(buffer int) *SafeStringChannel {
	return &SafeStringChannel{
		channel: make(chan string, buffer),
		closed: false,
	}
}

func (s *SafeStringChannel) WaitForClose() {
	if s.closed {
		return
	}

	s.closed = true

	go func() {
		for len(s.channel) != 0 {
			time.Sleep(time.Millisecond * 100)
		}

		close(s.channel)
	}()
}

func (s *SafeStringChannel) Offer(str string) {
	if s.closed {
		return
	}

	s.channel <- str
}

func (s *SafeStringChannel) ForEach(f func(str string)) {
	for v := range s.channel {
		f(v)
	}
}

func (s *SafeStringChannel) Closed() bool {
	return s.closed
}

func NewSafeFunctionChannel(buffer int) *SafeFunctionChannel {
	return &SafeFunctionChannel{
		channel: make(chan func(), buffer),
		closed: false,
	}
}

func (s *SafeFunctionChannel) WaitForClose() {
	if s.closed {
		return
	}

	s.closed = true

	go func() {
		for len(s.channel) != 0 {
			time.Sleep(time.Millisecond * 100)
		}

		close(s.channel)
	}()
}

func (s *SafeFunctionChannel) Offer(f func()) {
	if s.closed {
		return
	}

	s.channel <- f
}

func (s *SafeFunctionChannel) ForEach(f func(f func())) {
	for v := range s.channel {
		f(v)
	}
}

func (s *SafeFunctionChannel) Closed() bool {
	return s.closed
}
