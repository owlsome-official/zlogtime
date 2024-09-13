package zlogtime

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getTimeDuration(timeDuration time.Duration, unit string) int64 {
	switch unit {
	case "nano": // 10e-9
		return timeDuration.Nanoseconds()
	case "micro": // 10e-6
		return timeDuration.Microseconds()
	case "milli": // 10e-3
		return timeDuration.Milliseconds()
	default:
		return 0
	}
}

func getShortDurationUnit(unit string) string {
	switch unit {
	case "nano": // 10e-9
		return "ns"
	case "micro": // 10e-6
		return "us"
	case "milli": // 10e-3
		return "ms"
	default:
		return ""
	}
}

func getLogLevel(level string) *zerolog.Event {
	switch level {
	case "debug":
		return log.Logger.Debug()
	case "info":
		return log.Logger.Info()
	case "warn":
		return log.Logger.Warn()
	case "error":
		return log.Logger.Error()
	case "fatal":
		return log.Logger.Fatal()
	case "panic":
		return log.Logger.Panic()
	default:
		return log.Logger.Log()
	}
}
