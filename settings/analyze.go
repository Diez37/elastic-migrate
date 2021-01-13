package settings

type Analyze struct {
	maxTokenCount uint32
}

func NewAnalyze(maxTokenCount uint32) *Analyze {
	return &Analyze{maxTokenCount: maxTokenCount}
}

func (analyze *Analyze) Source() (interface{}, error) {
	return map[string]interface{}{
		"max_token_count": analyze.maxTokenCount,
	}, nil
}
