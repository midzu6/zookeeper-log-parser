package models

import (
	"fmt"
	"strings"
	"time"
)

type LogLevel string

func (el LogLevel) String() string {
	return string(el)
}

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelError LogLevel = "ERROR"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelTrace LogLevel = "TRACE"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     LogLevel  `json:"level"`
	Thread    string    `json:"thread"`
	Logger    string    `json:"logger"`
	Message   string    `json:"message"`
}

func NewLogEntry(
	timestamp time.Time,
	level LogLevel,
	thread string,
	logger string,
	message string) *LogEntry {

	return &LogEntry{
		Timestamp: timestamp,
		Level:     level,
		Thread:    strings.TrimSpace(thread),
		Logger:    strings.TrimSpace(logger),
		Message:   strings.TrimSpace(message),
	}
}

func (e *LogEntry) String() string {
	return fmt.Sprintf("[%s] %-5s %s - %s",
		e.Timestamp.Format("15:04:05.000"),
		e.Level,
		e.Thread,
		e.Message,
	)
}

func (e *LogEntry) IsInfo() bool {
	return e.Level == LogLevelInfo
}
func (e *LogEntry) IsWarn() bool {
	return e.Level == LogLevelWarn
}
func (e *LogEntry) IsError() bool {
	return e.Level == LogLevelError
}
func (e *LogEntry) IsDebug() bool {
	return e.Level == LogLevelDebug
}

func (e *LogEntry) IsTrace() bool {
	return e.Level == LogLevelTrace
}
