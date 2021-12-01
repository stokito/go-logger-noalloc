# go-logger-noalloc
Essential logger that avoids memory allocation. Intended for a high performance systems.
That's why a stack trace is not printed.
It prints in syslog format so later it can be easily be processed by journald from systemd.
For example:
```go
Log.Printf(logger.PREF_LOG_INFO + "Executing something\n")
```
So info record will be printed as `<6>Executing something` and the `<6>` here is a syslog prefix for INFO level.

## Install

    go get -u github.com/stokito/go-logger-noalloc@v1.0.0

## Usage

```go
package main

import (
	"github.com/stokito/go-logger-noalloc"
	"os"
)

func main() {
	// Log Prints everything to STDOUT but only if log level higher than INFO 
    Log := &logger.Logger{Out: os.Stdin, LogLevel: logger.LOG_INFO}

	Log.Printf(logger.PREF_LOG_DEBUG+ "Trace logging of vars like arg[0]: %s\n", os.Args[0])
	// Try to avoid unnecessary calculations if the DEBUG is anyway disabled
	if Log.IsLoggable(logger.LOG_DEBUG) {
		// some heavy calculations
		pwd, _ := os.Getwd()
		Log.Printf(logger.PREF_LOG_DEBUG+ "Started in %s\n", pwd)
	}
	Log.Printf(logger.PREF_LOG_INFO + "Describe execution step or the app sends/received a request from external system, minor error occurred like a timeout\n")
	Log.Printf(logger.PREF_LOG_WARNING + "Something suspicious happened, used deprecated API or an error occurred because a request is invalid\n")
	Log.Printf(logger.PREF_LOG_NOTICE + "Application did something important: processed a request, finished processing\n")
	Log.Printf(logger.PREF_LOG_ERR + "Unexpected internal error occurred: invalid request format\n")
	Log.Printf(logger.PREF_LOG_CRIT + "App can't do something: a port is already taken, missing config etc, fatal panic\n")
	// you can't log EMERG: leave it for OS
	// Now try to disable logs...
	Log.LogLevel = logger.LOG_EMERG
	Log.Printf(logger.PREF_LOG_ERR + "...and any error will be anyway printed\n")
	Log.Printf(logger.PREF_LOG_INFO + "But INFO now won't be printed\n")
}
```

See [example](examples/log_example.go)

