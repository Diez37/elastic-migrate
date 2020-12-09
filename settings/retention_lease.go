package settings

type RetentionLease struct {
	period *Interval
}

func NewRetentionLease(period *Interval) *RetentionLease {
	return &RetentionLease{period: period}
}

func (retentionLease *RetentionLease) Source() (interface{}, error) {
	intervalSource, err := retentionLease.period.Source()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"period": intervalSource,
	}, nil
}
