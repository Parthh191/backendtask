package logger

import (
	"fmt"
	"time"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	fmt.Printf("[%s] INFO: %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	fmt.Printf("[%s] ERROR: %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(msg, args...))
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	fmt.Printf("[%s] DEBUG: %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(msg, args...))
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	fmt.Printf("[%s] WARN: %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(msg, args...))
}
