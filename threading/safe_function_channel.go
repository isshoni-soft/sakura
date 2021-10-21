package threading

import (
	"github.com/isshoni-soft/sakura/logging"
	"time"
)

var safeFunctionChannelLogger = logging.NewLogger("safe-function-channel", 2)

type SafeFunctionChannel struct {
	channel chan func()
	closed bool
}

func NewSafeFunctionChannel(buffer int) *SafeFunctionChannel {
	return &SafeFunctionChannel {
		channel: make(chan func(), buffer),
		closed: false,
	}
}

func (s SafeFunctionChannel) Close() {
	safeFunctionChannelLogger.Log("closing safe function channel")

	if s.closed {
		return
	}

	s.closed = true

	go func() {
		safeFunctionChannelLogger.Log("termination queued")

		for len(s.channel) != 0 {
			time.Sleep(time.Millisecond * 100)
		}

		safeFunctionChannelLogger.Log("channel empty, closing...")

		close(s.channel)
	}()
}

func (s SafeFunctionChannel) Offer(f func()) {
	if s.closed {
		return
	}

	s.channel <- f
}
