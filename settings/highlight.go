package settings

type Highlight struct {
	maxAnalyzedOffset uint32
}

func NewHighlight(maxAnalyzedOffset uint32) *Highlight {
	return &Highlight{maxAnalyzedOffset: maxAnalyzedOffset}
}

func (highlight *Highlight) Source() (interface{}, error) {
	return map[string]interface{}{
		"max_analyzed_offset": highlight.maxAnalyzedOffset,
	}, nil
}
