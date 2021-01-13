package settings

type Routing struct {
	allocation *Allocation
	rebalance  *Rebalance
}

func NewRouting() *Routing {
	return &Routing{}
}

func (routing *Routing) SetAllocation(allocation *Allocation) *Routing {
	routing.allocation = allocation
	
	return routing
}

func (routing *Routing) SetRebalance(rebalance *Rebalance) *Routing {
	routing.rebalance = rebalance
	
	return routing
}

func (routing *Routing) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if routing.allocation != nil {
		allocationSource, err := routing.allocation.Source()
		if err != nil {
			return nil, err
		}

		source["allocation"] = allocationSource
	}

	if routing.rebalance != nil {
		rebalanceSource, err := routing.rebalance.Source()
		if err != nil {
			return nil, err
		}

		source["rebalance"] = rebalanceSource
	}

	return source, nil
}
