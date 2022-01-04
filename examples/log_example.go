package main

import (
	"github.com/stokito/go-logger-noalloc"
	"os"
)

// Log Prints everything to STDOUT but only if log level higher than INFO
var Log = &logger.Logger{Out: os.Stdout, LogLevel: logger.LOG_INFO}

func main() {
	// Log level must be concatenated as a prefix. The \n is required at end.
	Log.Printf(logger.PREF_LOG_DEBUG+"Trace logging of vars like arg[0]: %s\n", os.Args[0])
	// Try to avoid unnecessary calculations if the DEBUG is anyway disabled
	if Log.IsLoggable(logger.LOG_DEBUG) {
		// some heavy calculations
		pwd, _ := os.Getwd()
		Log.Printf(logger.PREF_LOG_DEBUG+"Started in %s\n", pwd)
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
