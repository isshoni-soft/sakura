package logging

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var logFileChannel = make(chan string, 5)
var defaultLogger = NewLogger("engine", 16)
var dateLayout = "01-02-2006|15:04:05"
var logFileName = "Sakura-" + time.Now().Format(dateLayout) + ".log"

type Logger struct {
	loggerChannel chan string

	prefix string
}

func init() {
	go logFileTick()
}

func GetLogger() *Logger {
	return defaultLogger
}

func NewLogger(prefix string, buffer int) *Logger {
	result := new(Logger)
	result.prefix =  "[" + time.Now().Format(dateLayout) +"]: sakura:" + prefix + "| "
	result.loggerChannel = make(chan string, buffer)

	go result.loggerTick()

	return result
}

func (l Logger) loggerTick() {
	for str := range l.loggerChannel {
		fmt.Println(str)

		logFileChannel <- str // now that we've logged the line lets queue it for adding to logfile
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

func logFileTick() {
	num := 1

	for _, err := os.Stat(logFileName); os.IsExist(err); {
		logFileName ="Sakura-" + time.Now().Format(dateLayout) + "-" + strconv.FormatInt(int64(num), 10) + ".log"
	}

	f, err := os.Create(logFileName)

	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()

		if err != nil {
			panic(err)
		}
	}(f)

	for str := range logFileChannel {
		_, err := f.WriteString(str + "\n")
		if err != nil {
			return
		}
	}
}