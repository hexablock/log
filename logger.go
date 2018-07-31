package log

import (
	"fmt"
	"log"
	"os"
)

// Logger is a default implementation of the Logger interface.
type Logger struct {
	stdout *log.Logger
	stderr *log.Logger

	debug     bool
	calldepth int
}

// NewDefaultLogger creates a default logger that prints to stderr or errors
// with everything else going to stdout
func NewDefaultLogger() *Logger {
	return &Logger{
		calldepth: 2,
		stderr:    log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds),
		stdout:    log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
	}
}

// SetCallDepth sets the call depth of the logger
func (l *Logger) SetCallDepth(d int) {
	l.calldepth = d
}

// EnableDebug enables debug messages to print.
func (l *Logger) EnableDebug(v bool) {
	if v {
		flg := log.LstdFlags | log.Lmicroseconds | log.Lshortfile

		l.stderr.SetFlags(flg)
		l.stdout.SetFlags(flg)
	}
	l.debug = v
}

// Debug implements the Logger interface.
func (l *Logger) Debug(v ...interface{}) {
	if l.debug {
		l.stderr.Output(l.calldepth, header("DBG", fmt.Sprint(v...)))
	}
}

// Debugf implements the Logger interface.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.stderr.Output(l.calldepth, header("DBG", fmt.Sprintf(format, v...)))
	}
}

// Info implements the Logger interface.
func (l *Logger) Info(v ...interface{}) {
	l.stdout.Output(l.calldepth, header("INF", fmt.Sprint(v...)))
}

// Infof implements the Logger interface.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.stdout.Output(l.calldepth, header("INF", fmt.Sprintf(format, v...)))
}

// Error implements the Logger interface.
func (l *Logger) Error(v ...interface{}) {
	l.stderr.Output(l.calldepth, header("ERR", fmt.Sprint(v...)))
}

// Errorf implements the Logger interface.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.stderr.Output(l.calldepth, header("ERR", fmt.Sprintf(format, v...)))
}

// Fatal implements the Logger interface.
func (l *Logger) Fatal(v ...interface{}) {
	l.stderr.Output(l.calldepth, header("FTL", fmt.Sprint(v...)))
	os.Exit(1)
}

// Fatalf implements the Logger interface.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.stderr.Output(l.calldepth, header("FTL", fmt.Sprintf(format, v...)))
	os.Exit(1)
}

// Panic implements the Logger interface.
func (l *Logger) Panic(v ...interface{}) {
	l.stderr.Panic(v)
}

// Panicf implements the Logger interface.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.stderr.Panicf(format, v...)
}

func header(lvl, msg string) string {
	return lvl + ": " + msg
}
