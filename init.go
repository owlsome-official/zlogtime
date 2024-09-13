package zlogtime

import (
	"time"

	"github.com/rs/zerolog"
)

var (
	TimeFieldFormat string = time.RFC3339Nano
)

func init() {
	zerolog.TimeFieldFormat = TimeFieldFormat
	zerolog.TimestampFieldName = "timestamp"
}

func New(conf ...Config) ZLogTime {
	// set default config
	cfg := configDefault(conf...)

	return cfg
}
