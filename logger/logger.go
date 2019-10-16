package logger

import (
	"io"
	"time"
)

type (
	Logger struct {
		writers        map[string]io.Writer
		fallbackWriter io.Writer
	}
	BaseRecord struct {
		path []string
		tags []string
	}
	IORecord struct {
		startedAt  *time.Time
		finishedAt *time.Time
	}
	ProfilingRecord struct {
	}
	DebugRecord struct {
	}
	Error struct {
	}
	Request struct {
	}
)

func (l *Logger) IO() {

}

func (l *Logger) Profiling() {

}

func (l *Logger) Debug() {

}

func (l *Logger) Error() {

}

func (l *Logger) Request() {

}
