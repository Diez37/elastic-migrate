package settings

type Analysis struct {
	analyzer   []Analyzer
	normalizer []Normalizer
	filter     []Filter
	charFilter []CharFilter
	tokenizer  []Tokenizer
}

func NewAnalysis() *Analysis {
	return &Analysis{}
}

func (analysis *Analysis) AddAnalyzer(analyzer ...Analyzer) *Analysis {
	analysis.analyzer = append(analysis.analyzer, analyzer...)

	return analysis
}

func (analysis *Analysis) AddNormalizer(normalizer ...Normalizer) *Analysis {
	analysis.normalizer = append(analysis.normalizer, normalizer...)

	return analysis
}

func (analysis *Analysis) AddFilter(filter ...Filter) *Analysis {
	analysis.filter = append(analysis.filter, filter...)

	return analysis
}

func (analysis *Analysis) AddCharFilter(charFilter ...CharFilter) *Analysis {
	analysis.charFilter = append(analysis.charFilter, charFilter...)

	return analysis
}

func (analysis *Analysis) AddTokenizer(tokenizer ...Tokenizer) *Analysis {
	analysis.tokenizer = append(analysis.tokenizer, tokenizer...)

	return analysis
}

func (analysis *Analysis) Source() (interface{}, error) {
	source := map[string]interface{}{}

	if len(analysis.analyzer) > 0 {
		items := map[AnalyzerName]interface{}{}

		for _, item := range analysis.analyzer {
			itemSource, err := item.Source()
			if err != nil {
				return nil, err
			}

			items[item.Name()] = itemSource
		}

		source["analyzer"] = items
	}

	if len(analysis.normalizer) > 0 {
		items := map[NormalizerName]interface{}{}

		for _, item := range analysis.normalizer {
			itemSource, err := item.Source()
			if err != nil {
				return nil, err
			}

			items[item.Name()] = itemSource
		}

		source["normalizer"] = items
	}

	if len(analysis.filter) > 0 {
		items := map[FilterName]interface{}{}

		for _, item := range analysis.filter {
			itemSource, err := item.Source()
			if err != nil {
				return nil, err
			}

			items[item.Name()] = itemSource
		}

		source["filter"] = items
	}

	if len(analysis.charFilter) > 0 {
		items := map[CharFilterName]interface{}{}

		for _, item := range analysis.charFilter {
			itemSource, err := item.Source()
			if err != nil {
				return nil, err
			}

			items[item.Name()] = itemSource
		}

		source["char_filter"] = items
	}

	if len(analysis.tokenizer) > 0 {
		items := map[TokenizerName]interface{}{}

		for _, item := range analysis.tokenizer {
			itemSource, err := item.Source()
			if err != nil {
				return nil, err
			}

			items[item.Name()] = itemSource
		}

		source["tokenizer"] = items
	}

	return source, nil
}
