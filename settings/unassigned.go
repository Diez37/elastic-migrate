package settings

type Unassigned struct {
	nodeLeft *NodeLeft
}

func NewUnassigned(nodeLeft *NodeLeft) *Unassigned {
	return &Unassigned{nodeLeft: nodeLeft}
}

func (unassigned *Unassigned) Source() (interface{}, error) {
	src, err := unassigned.nodeLeft.Source()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"node_left": src,
	}, nil
}
