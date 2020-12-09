package settings

type Threshold struct {
	fetch *LogLevel
	query *LogLevel
}

func NewThreshold() *Threshold {
	return &Threshold{}
}

func (threshold *Threshold) SetFetch(fetch *LogLevel) *Threshold {
	threshold.fetch = fetch
	
	return threshold
}

func (threshold *Threshold) SetQuery(query *LogLevel) *Threshold {
	threshold.query = query
	
	return threshold
}

func (threshold *Threshold) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if threshold.fetch != nil {
		fetchSource, err := threshold.fetch.Source()
		if err != nil {
			return nil, err
		}

		source["fetch"] = fetchSource
	}

	if threshold.query != nil {
		querySource, err := threshold.query.Source()
		if err != nil {
			return nil, err
		}

		source["query"] = querySource
	}

	return source, nil
}
