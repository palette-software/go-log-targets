package logging

import (
	"fmt"
	"time"
)

type TimedLogger struct {
	level   LogLevel
	message string
	start   time.Time
}

func NewTimedLogger(level LogLevel, format string, v ...interface{}) *TimedLogger {
	message := fmt.Sprintf(format, v...)
	timedLogger := TimedLogger{
		level:   level,
		message: message,
		start:   time.Now(),
	}
	Writef(level, message+" - START")
	return &timedLogger
}

func (tl *TimedLogger) LogEnd() {
	Writef(tl.level, fmt.Sprintf("%v - END  (elapsed %v ms)", tl.message, getElapsedTimeInMilliseconds(tl.start)))
}

func getElapsedTimeInMilliseconds(startTime time.Time) int64 {
	return (time.Since(startTime).Nanoseconds()) / 1e6
}
