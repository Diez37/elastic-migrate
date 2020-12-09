package settings

type Search struct {
	throttled *bool
	slowlog   *Slowlog
	idle      *Idle
}

func NewSearch() *Search {
	return &Search{}
}

func (search *Search) SetThrottled(throttled bool) *Search {
	search.throttled = &throttled
	
	return search
}

func (search *Search) SetSlowlog(slowlog *Slowlog) *Search {
	search.slowlog = slowlog
	
	return search
}

func (search *Search) SetIdle(idle *Idle) *Search {
	search.idle = idle
	
	return search
}

func (search *Search) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if search.throttled != nil {
		source["throttled"] = *search.throttled
	}

	if search.slowlog != nil {
		slowlog, err := search.slowlog.Source()
		if err != nil {
			return nil, err
		}

		source["slowlog"] = slowlog
	}

	if search.idle != nil {
		idle, err := search.idle.Source()
		if err != nil {
			return nil, err
		}

		source["idle"] = idle
	}

	return source, nil
}
