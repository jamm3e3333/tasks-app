package logger

import (
	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(message any, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message any, args ...any)
	Fatal(message any, args ...any)
	WithHook(h zerolog.Hook) *ZeroLogger
}
