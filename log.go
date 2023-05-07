// SPDX-License-Identifier: 0BSD
package olog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Priority same as in log/syslog
type Priority int

const (
	LOG_EMERG Priority = iota + '0'
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (
	CRIT   = "<2>"
	ERR    = "<3>"
	WARN   = "<4>"
	NOTICE = "<5>"
	INFO   = "<6>"
	DEBUG  = "<7>"
)

var logStdOut = &Logger{Out: os.Stdout, LogLevel: LOG_INFO}
var logStdErr = &Logger{Out: os.Stderr, LogLevel: LOG_INFO}

type Logger struct {
	Out      io.Writer // destination for output
	LogLevel Priority
	mu       sync.Mutex // ensures atomic writes
}

func (l *Logger) Printf(format string, v ...any) {
	if len(format) < 4 {
		return
	}
	// Priority prefix looks like <6> we get the 6 from it
	prioChar := format[1]
	prio := Priority(prioChar)
	if l.IsLoggable(prio) {
		l.printfBlocking(format, v...)
	}
}

func (l *Logger) printfBlocking(format string, v ...any) {
	l.mu.Lock()
	_, _ = fmt.Fprintf(l.Out, format, v...)
	l.mu.Unlock()
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.printfBlocking(CRIT+format, v...)
	os.Exit(1)
}

func (l *Logger) IsLoggable(logLevel Priority) bool {
	// errors are always logged
	return logLevel <= l.LogLevel
}

// shorthand functions for stdout

func Printf(format string, v ...any) {
	logStdOut.Printf(format, v...)
}

func IsLoggable(logLevel Priority) bool {
	return logStdOut.IsLoggable(logLevel)
}

func SetLogLevel(logLevel Priority) {
	logStdOut.LogLevel = logLevel
}

func Fatalf(format string, v ...any) {
	logStdOut.Fatalf(format, v...)
}

// shorthand functions for stderr

func ErrPrintf(format string, v ...any) {
	logStdErr.Printf(format, v...)
}

func ErrIsLoggable(logLevel Priority) bool {
	return logStdErr.IsLoggable(logLevel)
}

func ErrSetLogLevel(logLevel Priority) {
	logStdErr.LogLevel = logLevel
}

func ErrFatalf(format string, v ...any) {
	logStdErr.Fatalf(format, v...)
}
