package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/hinotora/sse-notification-service/internal/config"
)

type Logger struct {
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
}

var Instance *Logger = nil

func New() *Logger {
	return &Logger{
		info:  log.New(os.Stderr, "[system] INFO\t", log.Ldate|log.Ltime|log.Lmsgprefix),
		debug: log.New(os.Stderr, "[system] DEBUG\t", log.Ldate|log.Ltime|log.Lmsgprefix),
		err:   log.New(os.Stderr, "[system] ERROR\t", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
}

func (l *Logger) SetPrefix(prefix string) {
	l.info.SetPrefix(fmt.Sprintf("%s INFO\t", prefix))
	l.debug.SetPrefix(fmt.Sprintf("%s DEBUG\t", prefix))
	l.err.SetPrefix(fmt.Sprintf("%s ERROR\t", prefix))
}

func (l *Logger) Info(v any) {
	l.info.Println(v)
}

func (l *Logger) Debug(v any) {
	if config.GetInstance().App.Mode == "debug" {
		l.debug.Println(v)
	}
}

func (l *Logger) Error(v any) {
	l.err.Println(v)
}

func (l *Logger) Fatal(v any) {
	l.err.Println(v)
	os.Exit(0)
}
