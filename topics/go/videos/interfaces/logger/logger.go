package logger

import (
	"fmt"
	"io"
)

type Syncer interface {
	Sync() error
}

type Logger struct {
	w io.Writer
	s Syncer
}

func New(w io.Writer) *Logger {
	log := Logger{
		w: w,
		s: NopSyncer{},
	}

	if s, ok := w.(Syncer); ok {
		log.s = s
	}

	return &log
}

func (l *Logger) Info(msg string, args ...any) {
	fmt.Fprintf(l.w, "INFO: ")
	fmt.Fprintf(l.w, msg, args...)
	fmt.Fprintln(l.w)
	l.s.Sync()
}

type NopSyncer struct{}

func (NopSyncer) Sync() error { return nil }
