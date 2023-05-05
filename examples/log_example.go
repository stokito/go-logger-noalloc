package main

import (
	"github.com/stokito/go-logger-noalloc"
	"os"
)

// Log Prints everything to STDOUT but only if log level higher than INFO
var Log = &olog.Logger{Out: os.Stdout, LogLevel: olog.LOG_INFO}

func main() {
	// globally write to STDOUT
	olog.Printf(olog.DEBUG + "Debug message to STDOUT\n")
	olog.ErrPrintf(olog.ERR + "Error message to STDERR\n")

	// Log level must be concatenated as a prefix. The \n is required at end.
	Log.Printf(olog.DEBUG+"Trace logging of vars like arg[0]: %s\n", os.Args[0])
	// Try to avoid unnecessary calculations if the DEBUG is anyway disabled
	if Log.IsLoggable(olog.LOG_DEBUG) {
		// some heavy calculations
		pwd, _ := os.Getwd()
		Log.Printf("Started in %s\n", pwd)
		Log.Printf(olog.DEBUG+"Started in %s\n", pwd)
	}
	Log.Printf(olog.INFO + "Describe execution step or the app sends/received a request from external system, minor error occurred like a timeout\n")
	Log.Printf(olog.WARN + "Something suspicious happened, used deprecated API or an error occurred because a request is invalid\n")
	Log.Printf(olog.NOTICE + "Application did something important: processed a request, finished processing\n")
	Log.Printf(olog.ERR + "Unexpected internal error occurred: invalid request format\n")
	Log.Printf(olog.CRIT + "App can't do something: a port is already taken, missing config etc, fatal panic\n")
	// you can't log EMERG: leave it for OS
	// Now try to disable logs...
	Log.LogLevel = olog.LOG_EMERG
	Log.Printf(olog.ERR + "...and any error will be anyway printed\n")
	Log.Printf(olog.INFO + "But INFO now won't be printed\n")
	olog.ErrFatalf("%s\n", "fatal error")
}
