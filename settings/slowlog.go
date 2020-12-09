package settings

const (
	SlowLogWarn  SlowLogLevel = "WARN"
	SlowLogTrace SlowLogLevel = "TRACE"
	SlowLogDebug SlowLogLevel = "DEBUG"
	SlowLogInfo  SlowLogLevel = "INFO"
)

type SlowLogLevel string

type Slowlog struct {
	reformat  *bool
	source    *uint32
	level     *SlowLogLevel
	threshold *Threshold
}

func NewSlowlog() *Slowlog {
	return &Slowlog{}
}

func (slowlog *Slowlog) SetReformat(reformat bool) *Slowlog {
	slowlog.reformat = &reformat
	
	return slowlog
}

func (slowlog *Slowlog) SetSource(source uint32) *Slowlog {
	slowlog.source = &source
	
	return slowlog
}

func (slowlog *Slowlog) SetLevel(level SlowLogLevel) *Slowlog {
	slowlog.level = &level
	
	return slowlog
}

func (slowlog *Slowlog) SetThreshold(threshold *Threshold) *Slowlog {
	slowlog.threshold = threshold
	
	return slowlog
}

func (slowlog *Slowlog) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if slowlog.reformat != nil {
		source["reformat"] = *slowlog.reformat
	}

	if slowlog.source != nil {
		source["source"] = *slowlog.source
	}

	if slowlog.level != nil {
		source["level"] = *slowlog.level
	}

	if slowlog.threshold != nil {
		thresholdSource, err := slowlog.threshold.Source()
		if err != nil {
			return nil, err
		}

		source["threshold"] = thresholdSource
	}

	return source, nil
}
