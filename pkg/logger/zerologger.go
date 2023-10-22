package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	timestampFieldName = "@timestamp"
	timeFieldFormat    = time.RFC3339

	debugLvl = "debug"
	infoLvl  = "info"
	warnLvl  = "warn"
	errorLvl = "error"
	fatalLvl = "fatal"
)

type ZeroLogger struct {
	logger  *zerolog.Logger
	writers []io.Writer
}

var (
	_ Logger = (*ZeroLogger)(nil) // check if gin comply with Logger
)

func New(level string, develMode bool, additionalWriters ...io.Writer) *ZeroLogger {
	var multiWriter zerolog.LevelWriter

	var writers []io.Writer

	writers = append(writers, additionalWriters...)

	var consoleWriter zerolog.ConsoleWriter
	if develMode {
		// use human friendly version of console gin
		consoleWriter = zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: "15:04:05"}
		writers = append(writers, consoleWriter)
	} else {
		// production json gin
		writers = append(writers, os.Stdout)
	}

	// get global level
	lvl := getLevel(level)
	zerolog.SetGlobalLevel(lvl)

	// set the final writer
	multiWriter = zerolog.MultiLevelWriter(writers...)
	innerLog := zerolog.New(multiWriter).With().Timestamp().Logger().Level(lvl)

	zerolog.TimestampFieldName = timestampFieldName
	zerolog.TimeFieldFormat = timeFieldFormat

	return &ZeroLogger{
		logger:  &innerLog,
		writers: writers,
	}
}

func (l *ZeroLogger) GetLevel() string {
	return l.logger.GetLevel().String()
}

func getLevel(level string) zerolog.Level {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case errorLvl:
		l = zerolog.ErrorLevel
	case warnLvl:
		l = zerolog.WarnLevel
	case infoLvl:
		l = zerolog.InfoLevel
	case debugLvl:
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}
	return l
}

func (l *ZeroLogger) WithHook(h zerolog.Hook) *ZeroLogger {
	lh := l.logger.Hook(h)
	return &ZeroLogger{
		logger: &lh,
	}
}

func (l *ZeroLogger) Debug(message any, args ...any) {
	l.msg(debugLvl, message, args...)
}

func (l *ZeroLogger) Info(message string, args ...any) {
	l.msg(infoLvl, message, args...)
}

func (l *ZeroLogger) Warn(message string, args ...any) {
	l.msg(warnLvl, message, args...)
}

func (l *ZeroLogger) Error(message any, args ...any) {
	l.msg(errorLvl, message, args...)
}

func (l *ZeroLogger) Fatal(message any, args ...any) {
	l.msg(fatalLvl, message, args...)

	os.Exit(1)
}

func (l *ZeroLogger) getEventAtLevel(level string) *zerolog.Event {
	var e *zerolog.Event

	switch strings.ToLower(level) {
	case errorLvl:
		e = l.logger.Error()
	case warnLvl:
		e = l.logger.Warn()
	case debugLvl:
		e = l.logger.Debug()
	default: // default covers the info level
		e = l.logger.Info()
	}

	return e
}

func (l *ZeroLogger) log(level, message string, args ...any) {
	if len(args) == 0 {
		l.getEventAtLevel(level).Msg(message)
		return
	}
	l.getEventAtLevel(level).Msgf(message, args...)

}

func (l *ZeroLogger) msg(level string, message any, args ...any) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(warnLvl, fmt.Sprintf("[%s] message %v has unknown type %v", level, message, msg), args...)
	}
}

func (l *ZeroLogger) ZeroLogger() *zerolog.Logger {
	return l.logger
}
