package settings

type NodeLeft struct {
	delayedTimeout *Interval
}

func NewNodeLeft(delayedTimeout *Interval) *NodeLeft {
	return &NodeLeft{delayedTimeout: delayedTimeout}
}

func (nodeLeft *NodeLeft) Source() (interface{}, error) {
	intervalSource, err := nodeLeft.delayedTimeout.Source()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"delayed_timeout": intervalSource,
	}, nil
}
