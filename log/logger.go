package log

import (
	"io"
	"log"
)

type SeverityLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

type SEVERITY int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

func NewDefaultSeverityLogger() SeverityLogger {
	w := log.Default().Writer()
	return NewSeverityLogger(w)
}

func NewSeverityLogger(file io.Writer) SeverityLogger {
	logger := SeverityLogger{
		debugLogger: log.New(file, "DEBUG: ", log.LstdFlags),
		infoLogger:  log.New(file, "INFO: ", log.LstdFlags),
		warnLogger:  log.New(file, "WARNING: ", log.LstdFlags),
		errorLogger: log.New(file, "ERROR: ", log.LstdFlags),
		fatalLogger: log.New(file, "FATAL: ", log.LstdFlags),
	}

	return logger
}

func (s *SeverityLogger) GetInternalLogger(severity SEVERITY) *log.Logger {
	switch severity {
	case DEBUG:
		return s.debugLogger
	case INFO:
		return s.infoLogger
	case WARN:
		return s.warnLogger
	case ERROR:
		return s.errorLogger
	case FATAL:
		return s.fatalLogger
	default:
		return s.infoLogger
	}
}

func (s *SeverityLogger) Log(severity SEVERITY, msg string, args ...interface{}) {
	msg += "\n"
	switch severity {
	case DEBUG:
		s.debugLogger.Printf(msg, args...)
	case INFO:
		s.infoLogger.Printf(msg, args...)
	case WARN:
		s.warnLogger.Printf(msg, args...)
	case ERROR:
		s.errorLogger.Printf(msg, args...)
	case FATAL:
		s.fatalLogger.Printf(msg, args...)
	default:
		s.infoLogger.Printf(msg, args...)
	}
}

func (s *SeverityLogger) Debug(msg string, args ...interface{}) {
	s.Log(DEBUG, msg, args...)
}

func (s *SeverityLogger) Info(msg string, args ...interface{}) {
	s.Log(INFO, msg, args...)
}

func (s *SeverityLogger) Warn(msg string, args ...interface{}) {
	s.Log(WARN, msg, args...)
}

func (s *SeverityLogger) Error(msg string, args ...interface{}) {
	s.Log(ERROR, msg, args...)
}

func (s *SeverityLogger) Fatal(msg string, args ...interface{}) {
	s.Log(FATAL, msg, args...)
}
