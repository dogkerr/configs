package logger

import (
	"fmt"
	"os"
	"reflect"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New() *Logger {
	// var l zerolog.Level

	// switch strings.ToLower(level) {
	// case "error":
	// 	l = zerolog.ErrorLevel
	// case "warn":
	// 	l = zerolog.WarnLevel
	// case "info":
	// 	l = zerolog.InfoLevel
	// case "debug":
	// 	l = zerolog.DebugLevel
	// default:
	// 	l = zerolog.InfoLevel
	// }

	// zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

type LogMetadata struct {
	Clientid   string
	Method     string
	StatusCode int
	BodySize   int
	Path       string
	Latency    string
}

func getLogMetadata(args ...interface{}) LogMetadata {

	logMtd := reflect.ValueOf(args[0]).Interface().(LogMetadata)

	return logMtd
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {

	l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.msg("warn", message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	// if l.logger.GetLevel() == zerolog.DebugLevel {
	// 	l.Debug(message, args...)
	// }

	l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {

	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		logMetadata := getLogMetadata(args...)

		l.logger.Info().
			Str("client_id", logMetadata.Clientid).
			Str("method", logMetadata.Method).
			Str("status_code", fmt.Sprint(logMetadata.StatusCode)).
			Str("body_size", fmt.Sprint(logMetadata.BodySize)).
			Str("path", logMetadata.Path).
			Str("latency", logMetadata.Latency).
			Msgf(message, args...)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {

	switch msg := message.(type) {
	case error:
		if level == "error" {
			if len(args) == 0 {
				l.logger.Error().Stack().Err(msg).Msg(msg.Error())
			} else {
				logMetadata := getLogMetadata(args...)
				l.logger.Error().Stack().Err(msg).
					Str("client_id", logMetadata.Clientid).
					Str("method", logMetadata.Method).
					Str("status_code", fmt.Sprint(logMetadata.StatusCode)).
					Str("body_size", fmt.Sprint(logMetadata.BodySize)).
					Str("path", logMetadata.Path).
					Str("latency", logMetadata.Latency).Msg(msg.Error())
			}
		} 
		// else {
		// 	logMetadata := getLogMetadata(args...)
		// 	l.logger.Fatal().Stack().Err(msg).
		// 		Str("client_id", logMetadata.Clientid).
		// 		Str("method", logMetadata.Method).
		// 		Str("status_code", fmt.Sprint(logMetadata.StatusCode)).
		// 		Str("body_size", fmt.Sprint(logMetadata.BodySize)).
		// 		Str("path", logMetadata.Path).
		// 		Str("latency", logMetadata.Latency).Msg(msg.Error())
		// }
	case string:
		if level == "debug" {
			logMetadata := getLogMetadata(args...)
			l.logger.Debug().
				Str("client_id", logMetadata.Clientid).
				Str("method", logMetadata.Method).
				Str("status_code", fmt.Sprint(logMetadata.StatusCode)).
				Str("body_size", fmt.Sprint(logMetadata.BodySize)).
				Str("path", logMetadata.Path).
				Str("latency", logMetadata.Latency).Msg(msg)
		} else if level == "warn" {
			logMetadata := getLogMetadata(args...)
			l.logger.Warn().
				Str("client_id", logMetadata.Clientid).
				Str("method", logMetadata.Method).
				Str("status_code", fmt.Sprint(logMetadata.StatusCode)).
				Str("body_size", fmt.Sprint(logMetadata.BodySize)).
				Str("path", logMetadata.Path).
				Str("latency", logMetadata.Latency).Msg(msg)
		}

	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
