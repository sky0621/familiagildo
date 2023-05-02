package log

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/familiagildo/app"
	"os"
	"strings"
)

func Trace(msg string) {
	log.Trace().Msg(msg)
}

func Tracef(format string, v ...any) {
	log.Trace().Msgf(format, v...)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Debugf(format string, v ...any) {
	log.Debug().Msgf(format, v...)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(format string, v ...any) {
	log.Info().Msgf(format, v...)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Warnf(format string, v ...any) {
	log.Warn().Msgf(format, v...)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func Errorf(format string, v ...any) {
	log.Error().Msgf(format, v...)
}

func ErrorWith(err error, msg string) {
	log.Err(errors.UnwrapAll(err)).Msg(msg)
}

func ErrorfWith(err error, format string, v ...any) {
	log.Err(errors.UnwrapAll(err)).Msgf(format, v...)
}

func ErrorSend(err error) {
	log.Err(errors.UnwrapAll(err)).Send()
}

func Fatal(msg string) {
	log.Fatal().Msg(msg)
}

func Fatalf(format string, v ...any) {
	log.Fatal().Msgf(format, v...)
}

func Panic(msg string) {
	log.Panic().Msg(msg)
}

func Panicf(format string, v ...any) {
	log.Panic().Msgf(format, v...)
}

func Logger(env app.Env) {
	if env.IsGCP() {
		gcpLogger()
		return
	}
	if env.IsAWS() {
		awsLogger()
		return
	}
	if env.IsLocal() {
		localLogger()
		return
	}
}

func gcpLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "severity"
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger()
}

func awsLogger() {
	// FIXME:
}

func localLogger() {
	// Pretty logging
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}
	log.Logger = zerolog.New(output).With().Stack().Timestamp().Caller().Logger()
}
