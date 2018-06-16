package logger

import (
	"fmt"
	"io"
	"log"
	"time"
)

type LogLevel int

const (
	Trace   LogLevel = 0
	Info    LogLevel = 1
	Warning LogLevel = 2
	Error   LogLevel = 3
)

var logLevelStr = map[LogLevel]string{
	Trace:   "TRACE",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR"}

type Logger struct {
	logLevel LogLevel

	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
}

type logWriter struct {
	packageName string

	logLevel LogLevel
}

func (lw *logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("20-01-2006 15:04:05.99") + fmt.Sprintf("\t[%s]\t%s:\t%s", logLevelStr[lw.logLevel], lw.packageName, string(bytes)))
}

// Instantiate new logger for given package
func New(packageName string, logLevel LogLevel, writer io.Writer) Logger {
	var logger Logger
	logger.logLevel = logLevel

	logger.trace = log.New(writer,
		"",
		log.Ldate|log.Ltime)
	traceLogWriter := &logWriter{packageName: packageName, logLevel: Trace}
	logger.trace.SetOutput(traceLogWriter)

	logger.info = log.New(writer,
		"",
		log.Ldate|log.Ltime)
	infoLogWriter := &logWriter{packageName: packageName, logLevel: Info}
	logger.info.SetOutput(infoLogWriter)

	logger.warning = log.New(writer,
		"",
		log.Ldate|log.Ltime)
	warningLogWriter := &logWriter{packageName: packageName, logLevel: Warning}
	logger.warning.SetOutput(warningLogWriter)

	logger.error = log.New(writer,
		"",
		log.Ldate|log.Ltime)
	errorLogWriter := &logWriter{packageName: packageName, logLevel: Error}
	logger.error.SetOutput(errorLogWriter)

	return logger
}

func (l *Logger) Trace(log string) {
	if l.logLevel <= Trace {
		l.trace.Println(log)
	}
}

func (l *Logger) Tracef(format string, a ...interface{}) {
	if l.logLevel <= Trace {
		l.trace.Printf(format, a)
	}
}

func (l *Logger) Info(log string) {
	if l.logLevel <= Info {
		l.info.Println(log)
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	if l.logLevel <= Info {
		l.info.Printf(format, a)
	}
}

func (l *Logger) Warning(log string) {
	if l.logLevel <= Warning {
		l.warning.Println(log)
	}
}

func (l *Logger) Warningf(format string, a ...interface{}) {
	if l.logLevel <= Warning {
		l.warning.Printf(format, a)
	}
}

func (l *Logger) Error(log string) {
	if l.logLevel <= Error {
		l.error.Println(log)
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	if l.logLevel <= Error {
		l.error.Printf(format, a)
	}
}
