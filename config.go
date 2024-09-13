package zlogtime

type Config struct {

	// Optional. Default: false
	Hidden bool

	// Optional. Default: "info"
	LogLevel string

	// Optional. Default: "milli". Possible Value: ["nano", "micro", "milli"]
	ElapsedTimeUnit string
}

var ConfigDefault = Config{
	Hidden:          false,
	LogLevel:        "info",
	ElapsedTimeUnit: "milli",
}

func configDefault(config ...Config) ZLogTime {
	// Return default config if nothing provided
	if len(config) < 1 {
		return &ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	// Note: cfg.Hidden: it's false by default.

	if cfg.LogLevel == "" {
		cfg.LogLevel = ConfigDefault.LogLevel
	}

	if cfg.ElapsedTimeUnit == "" {
		cfg.ElapsedTimeUnit = ConfigDefault.ElapsedTimeUnit
	}

	return &cfg
}
