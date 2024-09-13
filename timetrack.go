package zlogtime

import (
	"strconv"
	"time"
)

type ZLogTime interface {
	TimeTrack(name string, start time.Time)
}

func (cfg *Config) TimeTrack(name string, start time.Time) {
	eventName := "TimeTrack"
	if name != "" {
		eventName = name
	}

	logTime := time.Now()
	elapsedTime := getTimeDuration(time.Since(start), cfg.ElapsedTimeUnit)
	shortUnit := getShortDurationUnit(cfg.ElapsedTimeUnit)

	logger := getLogLevel(cfg.LogLevel)
	logger = logger.
		Interface("event", map[string]interface{}{
			"name":                    eventName,
			"start":                   start.Format(TimeFieldFormat),
			"end":                     logTime.Format(TimeFieldFormat),
			"elapsed_time":            elapsedTime,
			"elapsed_time_unit":       cfg.ElapsedTimeUnit,
			"elapsed_time_unit_short": shortUnit,
			"elapsed_time_string":     strconv.FormatInt(elapsedTime, 10) + shortUnit,
		})

	logger.Msg("TimeTrack at: " + name)
}
