package settings

type LogLevel struct {
	warn  *int8
	trace *int8
	debug *int8
	info  *int8
}

func NewLogLevel() *LogLevel {
	return &LogLevel{}
}

func (logLevel *LogLevel) SetWarn(warn int8) *LogLevel {
	logLevel.warn = &warn

	return logLevel
}

func (logLevel *LogLevel) SetTrace(trace int8) *LogLevel {
	logLevel.trace = &trace

	return logLevel
}

func (logLevel *LogLevel) SetDebug(debug int8) *LogLevel {
	logLevel.debug = &debug

	return logLevel
}

func (logLevel *LogLevel) SetInfo(info int8) *LogLevel {
	logLevel.info = &info

	return logLevel
}

func (logLevel *LogLevel) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if logLevel.warn != nil {
		source["warn"] = *logLevel.warn
	}

	if logLevel.trace != nil {
		source["trace"] = *logLevel.trace
	}

	if logLevel.debug != nil {
		source["debug"] = *logLevel.debug
	}

	if logLevel.info != nil {
		source["info"] = *logLevel.info
	}

	return source, nil
}
