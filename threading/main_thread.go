package threading

import (
	"github.com/isshoni-soft/sakura/logging"
	"runtime"
)

var mainThreadCallQueue *SafeFunctionChannel

var mainThreadLogger = logging.NewLogger("main-thread", 10)

func init() {
	runtime.LockOSThread()
}

func InitMainThread(cap int, run func()) {
	mainThreadCallQueue = NewSafeFunctionChannel(cap)

	go run()

	for f := range mainThreadCallQueue.channel {
		f()
	}
}

func RunMain(f func(), block ...bool) {
	if IsMainThreadRunning() {
		return
	}

	if len(block) == 1 && block[0] {
		done := make(chan bool, 1)
		mainThreadCallQueue.Offer(func() {
			f()
			done <- true
		})

		<- done
	} else if len(block) == 0 || block[0] {
		mainThreadCallQueue.Offer(f)
	}
}

func RunMainResult(f func() interface {}) interface {} {
	if IsMainThreadRunning() {
		return nil
	}

	done := make(chan interface {}, 1)

	mainThreadCallQueue.Offer(func() {
		done <- f()
	})

	return <- done
}

func ShutdownMainThread() {
	mainThreadLogger.Log("Shutting down main thread manager...")

	mainThreadCallQueue.Close()
}

func IsMainThreadRunning() bool {
	return mainThreadCallQueue.closed
}
