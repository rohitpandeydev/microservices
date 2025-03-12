package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	logger *log.Logger
	level  Level
}

func New(level Level) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", 0),
		level:  level,
	}
}

func (l *Logger) log(level Level, format string, args ...any) {
	if level < l.level {
		return
	}
	_, file, line, _ := runtime.Caller(2)
	timestamp := time.Now().Format("2026-01-02 15:04:30")
	levelStr := [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[level]
	message := fmt.Sprintf(format, args...)

	l.logger.Printf("[%s] %s %s:%d: %s", levelStr, timestamp, file, line, message)
}

func (l *Logger) Debug(format string, args ...any) {
	l.log(DEBUG, format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.log(INFO, format, args...)
}

func (l *Logger) Warn(format string, args ...any) {
	l.log(WARN, format, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.log(ERROR, format, args...)
}
