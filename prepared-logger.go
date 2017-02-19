package log

import (
	"fmt"
	"time"
	"runtime"
)

type PreparedLogger struct {
	fields []Field
}


func (pl *PreparedLogger) Clone() *PreparedLogger {
	fs := make([]Field, 0)
	return &PreparedLogger{
		fields: append(fs, pl.fields...),
	}
}

// Trace starts a trace & returns Traceable object to End + log.
// Example defer log.Trace(...).End()
func (pl *PreparedLogger) Trace(v ...interface{}) Traceable {

	t := WithFields(pl.fields...).Trace(v...)

	return t
}

func (pl *PreparedLogger) WithFields(f ...Field) LeveledLogger {
	pl.fields = append(pl.fields, f...)
	return pl
}

func (pl *PreparedLogger) WithError(err error) LeveledLogger {

	pl.fields = append(pl.fields, F(`err`,err))
	return pl
}

// StackTrace creates a new log Entry with pre-populated field with stack trace.
func (pl *PreparedLogger) StackTrace() LeveledLogger {

	trace := make([]byte, 1<<16)
	n := runtime.Stack(trace, true)
	if n > stackTraceLimit {
		n = stackTraceLimit
	}

	fields := append(pl.fields, []Field{F("stack trace", string(trace[:n])+"\n")}...)

	return newEntry(DebugLevel, "", fields, skipLevel)

}

//Tracef(msg string, v ...interface{}) Traceable
func (pl *PreparedLogger) Tracef(msg string, v ...interface{}) Traceable {
	t := Logger.tracePool.Get().(*TraceEntry)
	t.entry = newEntry(TraceLevel, fmt.Sprintf(msg, v...), pl.fields, skipLevel)
	t.start = time.Now().UTC()

	return t
}

// Debug level formatted message.
func (pl *PreparedLogger) Debug(v ...interface{}) {
	WithFields(pl.fields...).Debug(v...)
}

// Info level formatted message.
func (pl *PreparedLogger) Info(v ...interface{}) {
	WithFields(pl.fields...).Info(v...)
}

// Notice level formatted message.
func (pl *PreparedLogger) Notice(v ...interface{}) {
	WithFields(pl.fields...).Notice(v...)
}

// Warn level formatted message.
func (pl *PreparedLogger) Warn(v ...interface{}) {
	WithFields(pl.fields...).Warn(v...)
}

// Error level formatted message.
func (pl *PreparedLogger) Error(v ...interface{}) {
	WithFields(pl.fields...).Error(v...)
}

// Panic logs an Panic level formatted message and then panics
func (pl *PreparedLogger) Panic(v ...interface{}) {
	WithFields(pl.fields...).Panic(v...)
}

// Alert logs an Alert level formatted message and then panics
func (pl *PreparedLogger) Alert(v ...interface{}) {
	WithFields(pl.fields...).Alert(v...)
}

// Fatal level formatted message, followed by an exit.
func (pl *PreparedLogger) Fatal(v ...interface{}) {
	WithFields(pl.fields...).Fatal(v...)
}

// Debugf level formatted message.
func (pl *PreparedLogger) Debugf(msg string, v ...interface{}) {
	WithFields(pl.fields...).Debugf(msg, v...)
}

// Infof level formatted message.
func (pl *PreparedLogger) Infof(msg string, v ...interface{}) {
	WithFields(pl.fields...).Infof(msg, v...)
}

// Noticef level formatted message.
func (pl *PreparedLogger) Noticef(msg string, v ...interface{}) {
	WithFields(pl.fields...).Noticef(msg, v...)
}

// Warnf level formatted message.
func (pl *PreparedLogger) Warnf(msg string, v ...interface{}) {
	WithFields(pl.fields...).Warnf(msg, v...)
}

// Errorf level formatted message.
func (pl *PreparedLogger) Errorf(msg string, v ...interface{}) {
	WithFields(pl.fields...).Errorf(msg, v...)
}

// Panicf logs an Panic level formatted message and then panics
func (pl *PreparedLogger) Panicf(msg string, v ...interface{}) {
	WithFields(pl.fields...).Panicf(msg, v...)
}

// Alertf logs an Alert level formatted message and then panics
func (pl *PreparedLogger) Alertf(msg string, v ...interface{}) {
	WithFields(pl.fields...).Alertf(msg, v...)
}

// Fatalf level formatted message, followed by an exit.
func (pl *PreparedLogger) Fatalf(msg string, v ...interface{}) {
	WithFields(pl.fields...).Fatalf(msg, v...)
}