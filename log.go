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
	LOG_EMERG Priority = iota
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

type Logger struct {
	Out      io.Writer // destination for output
	LogLevel Priority
	mu       sync.Mutex // ensures atomic writes
}

func (l *Logger) Printf(format string, v ...any) {
	// Priority prefix looks like <6> we get the 6 from it
	prioChar := format[1:2][0]
	prio := Priority(prioChar - '0')
	if !l.IsLoggable(prio) {
		return
	}
	l.mu.Lock()
	fmt.Fprintf(l.Out, format, v...)
	l.mu.Unlock()
}

func (l *Logger) Fatal(v ...any) {
	l.Printf(CRIT+"%s\n", v...)
	os.Exit(1)
}

func (l *Logger) IsLoggable(logLevel Priority) bool {
	// errors are always logged
	return logLevel <= LOG_ERR ||
		logLevel <= l.LogLevel
}

func Printf(format string, v ...any) {
	logStdOut.Printf(format, v)
}

func IsLoggable(logLevel Priority) bool {
	return logStdOut.IsLoggable(logLevel)
}

func Fatal(format string, v ...any) {
	logStdOut.Fatal(format, v)
}
