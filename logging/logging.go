package logging

import "fmt"

var defaultLogger = NewLogger("engine", 16)

type Logger struct {
	loggerChannel chan string

	prefix string
}

func GetLogger() *Logger {
	return defaultLogger
}

func NewLogger(prefix string, buffer int) *Logger {
	result := new(Logger)
	result.prefix =  "sakura:" + prefix + "| "
	result.loggerChannel = make(chan string, buffer)

	go result.loggerTick()

	return result
}

func (l Logger) loggerTick() {
	for str := range l.loggerChannel {
		fmt.Println(str)
	}
}

func (l Logger) Shutdown() {
	close(l.loggerChannel)
}

func (l Logger) SetPrefix(str string) {
	l.prefix = str
}

func (l Logger) Format(str ...string) (result string) {
	result = l.prefix

	for _, s := range str {
		result = result + s
	}

	return
}

func (l Logger) Log(str ...string) {
	l.loggerChannel <- l.Format(str...)
}
