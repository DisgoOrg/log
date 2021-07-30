package log

import (
	"fmt"
	"log"
	"os"
)

var _ Logger = (*SimpleLogger)(nil)

var std *SimpleLogger = nil

// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// With the exception of the Lmsgprefix flag, there is no
// control over the order they appear (the order listed here)
// or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//	2009/01/23 01:23:23 message
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// Level are different levels at which the SimpleLogger can log
type Level int

// All Level(s) which SimpleLogger supports
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// String returns the name of the Level
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO "
	case LevelWarn:
		return "WARN "
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelPanic:
		return "PANIC"
	default:
		return ""
	}
}

//Default returns the default SimpleLogger
//goland:noinspection GoUnusedExportedFunction
func Default() *SimpleLogger {
	if std == nil {
		std = New(log.LstdFlags | log.Lmsgprefix)
	}
	return std
}

// New returns a new SimpleLogger implementation
//goland:noinspection GoUnusedExportedFunction
func New(flags int) *SimpleLogger {
	return &SimpleLogger{
		logger: log.New(os.Stderr, "", flags),
		level:  LevelInfo,
	}
}

// SimpleLogger is a wrapper for the std Logger
type SimpleLogger struct {
	logger *log.Logger
	level  Level
}

// SetLevel sets the lowest Level to log for
func (l *SimpleLogger) SetLevel(level Level) {
	l.level = level
}

// SetFlags sets the log flags like: Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC, Lmsgprefix,LstdFlags
func (l *SimpleLogger) SetFlags(flags int) {
	l.logger.SetFlags(flags)
}

func (l *SimpleLogger) log(level Level, args ...interface{}) {
	if level < l.level {
		return
	}
	l.logger.SetPrefix(level.String() + " ")
	switch level {
	case LevelFatal:
		l.logger.Fatal(args...)
	case LevelPanic:
		l.logger.Panic(args...)
	default:
		l.logger.Print(args...)
	}
}
func (l *SimpleLogger) logf(level Level, format string, args ...interface{}) {
	l.log(level, fmt.Sprintf(format, args...))
}

// Debug logs on the LevelDebug
func (l *SimpleLogger) Debug(args ...interface{}) {
	l.log(LevelDebug, args...)
}

// Debugf logs on the LevelDebug
func (l *SimpleLogger) Debugf(format string, args ...interface{}) {
	l.logf(LevelDebug, format, args...)
}

// Info logs on the LevelInfo
func (l *SimpleLogger) Info(args ...interface{}) {
	l.log(LevelInfo, args...)
}

// Infof logs on the LevelInfo
func (l *SimpleLogger) Infof(format string, args ...interface{}) {
	l.logf(LevelInfo, format, args...)
}

// Warn logs on the LevelWarn
func (l *SimpleLogger) Warn(args ...interface{}) {
	l.log(LevelWarn, args...)
}

// Warnf logs on the LevelWarn
func (l *SimpleLogger) Warnf(format string, args ...interface{}) {
	l.logf(LevelWarn, format, args...)
}

// Error logs on the LevelError
func (l *SimpleLogger) Error(args ...interface{}) {
	l.log(LevelError, args...)
}

// Errorf logs on the LevelError
func (l *SimpleLogger) Errorf(format string, args ...interface{}) {
	l.logf(LevelError, format, args...)
}

// Fatal logs on the LevelFatal
func (l *SimpleLogger) Fatal(args ...interface{}) {
	l.log(LevelFatal, args...)
}

// Fatalf logs on the LevelFatal
func (l *SimpleLogger) Fatalf(format string, args ...interface{}) {
	l.logf(LevelFatal, format, args...)
}

// Panic logs on the LevelPanic
func (l *SimpleLogger) Panic(args ...interface{}) {
	l.log(LevelPanic, args...)
}

// Panicf logs on the LevelPanic
func (l *SimpleLogger) Panicf(format string, args ...interface{}) {
	l.logf(LevelPanic, format, args...)
}

// SetLevel sets the Level of the default Logger
func SetLevel(level Level) {
	std.SetLevel(level)
}

// SetFlags sets the log flags like: Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC, Lmsgprefix,LstdFlags of the default Logger
func SetFlags(flags int) {
	std.SetFlags(flags)
}

// Debug logs on the LevelDebug with the default SimpleLogger
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf logs on the LevelDebug with the default SimpleLogger
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Info logs on the LevelInfo with the default SimpleLogger
func Info(args ...interface{}) {
	std.Info(args...)
}

// Infof logs on the LevelInfo with the default SimpleLogger
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warn logs on the LevelWarn with the default SimpleLogger
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Warnf logs on the Level with the default SimpleLogger
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Error logs on the LevelError with the default SimpleLogger
func Error(args ...interface{}) {
	std.Error(args...)
}

// Errorf logs on the LevelError with the default SimpleLogger
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Fatal logs on the LevelFatal with the default SimpleLogger
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Fatalf logs on the LevelFatal with the default SimpleLogger
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Panic logs on the LevelPanic with the default SimpleLogger
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Panicf logs on the LevelPanic with the default SimpleLogger
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}