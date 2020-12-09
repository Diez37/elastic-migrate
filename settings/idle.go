package settings

type Idle struct {
	after *Interval
}

func NewIdle(after *Interval) *Idle {
	return &Idle{after: after}
}

func (idle *Idle) Source() (interface{}, error) {
	sourceAfter, err := idle.after.Source()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"after": sourceAfter,
	}, nil
}
