package app

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

func Logger(env Env) {
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
	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
}
