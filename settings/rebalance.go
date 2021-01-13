package settings

const (
	RebalanceAll       RebalanceVal = "all"
	RebalancePrimaries RebalanceVal = "primaries"
	RebalanceReplicas  RebalanceVal = "replicas"
	RebalanceNone      RebalanceVal = "none"
)

type RebalanceVal string

type Rebalance struct {
	enable RebalanceVal
}

func NewRebalance(enable RebalanceVal) *Rebalance {
	return &Rebalance{enable: enable}
}

func (rebalance *Rebalance) Source() (interface{}, error) {
	return map[string]interface{}{
		"enable": rebalance.enable,
	}, nil
}
