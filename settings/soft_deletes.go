package settings

type SoftDeletes struct {
	enabled        *bool
	retention      *Retention
	retentionLease *RetentionLease
}

func NewSoftDeletes() *SoftDeletes {
	return &SoftDeletes{}
}

func (softDeletes *SoftDeletes) SetEnabled(enabled bool) *SoftDeletes {
	softDeletes.enabled = &enabled
	
	return softDeletes
}

func (softDeletes *SoftDeletes) SetRetention(retention *Retention) *SoftDeletes {
	softDeletes.retention = retention

	return softDeletes
}

func (softDeletes *SoftDeletes) SetRetentionLease(retentionLease *RetentionLease) *SoftDeletes {
	softDeletes.retentionLease = retentionLease

	return softDeletes
}

func (softDeletes *SoftDeletes) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if softDeletes.enabled != nil {
		source["enabled"] = *softDeletes.enabled
	}

	if softDeletes.retention != nil {
		src, err := softDeletes.retention.Source()
		if err != nil {
			return nil, err
		}

		source["retention"] = src
	}

	if softDeletes.retentionLease != nil {
		src, err := softDeletes.retentionLease.Source()
		if err != nil {
			return nil, err
		}

		source["retention_lease"] = src
	}

	return source, nil
}
