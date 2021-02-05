package fields

import "github.com/diez37/elastic-migrate/settings"

type Text struct {
	name                     FieldName
	fieldData                *bool
	eagerGlobalOrdinals      *bool
	indexPhrases             *bool
	index                    *bool
	norms                    *bool
	store                    *bool
	positionIncrementGap     *int
	boost                    *float64
	analyzer                 *settings.AnalyzerName
	searchAnalyzer           *settings.AnalyzerName
	searchQuoteAnalyzer      *settings.AnalyzerName
	fieldDataFrequencyFilter *FieldDataFrequencyFilter
	indexOptions             *IndexOption
	indexPrefixes            *IndexPrefixes
	similarity               *settings.SimilarityName
	termVector               *TermVector
	fields                   []*Field
	meta                     []*Meta
}

func NewText(name string) *Text {
	return &Text{name: FieldName(name)}
}

func (field *Text) Name() FieldName {
	return field.name
}

func (field *Text) SetStore(store bool) *Text {
	field.store = &store

	return field
}

func (field *Text) SetIndex(index bool) *Text {
	field.index = &index

	return field
}

func (field *Text) SetEagerGlobalOrdinals(eagerGlobalOrdinals bool) *Text {
	field.eagerGlobalOrdinals = &eagerGlobalOrdinals

	return field
}

func (field *Text) SetNorms(norms bool) *Text {
	field.norms = &norms

	return field
}

func (field *Text) SetFieldData(fieldData bool) *Text {
	field.fieldData = &fieldData

	return field
}

func (field *Text) SetIndexPhrases(indexPhrases bool) *Text {
	field.indexPhrases = &indexPhrases

	return field
}

func (field *Text) SetPositionIncrementGap(positionIncrementGap int) *Text {
	field.positionIncrementGap = &positionIncrementGap

	return field
}

func (field *Text) SetBoost(boost float64) *Text {
	field.boost = &boost

	return field
}

func (field *Text) SetAnalyzer(analyzer settings.AnalyzerName) *Text {
	field.analyzer = &analyzer

	return field
}

func (field *Text) SetSearchAnalyzer(searchAnalyzer settings.AnalyzerName) *Text {
	field.searchAnalyzer = &searchAnalyzer

	return field
}

func (field *Text) SetSearchQuoteAnalyzer(searchQuoteAnalyzer settings.AnalyzerName) *Text {
	field.searchQuoteAnalyzer = &searchQuoteAnalyzer

	return field
}

func (field *Text) SetFieldDataFrequencyFilter(fieldDataFrequencyFilter *FieldDataFrequencyFilter) *Text {
	field.fieldDataFrequencyFilter = fieldDataFrequencyFilter

	return field
}

func (field *Text) SetIndexOptions(indexOptions IndexOption) *Text {
	field.indexOptions = &indexOptions

	return field
}

func (field *Text) SetIndexPrefixes(indexPrefixes *IndexPrefixes) *Text {
	field.indexPrefixes = indexPrefixes

	return field
}

func (field *Text) SetSimilarity(similarity settings.SimilarityName) *Text {
	field.similarity = &similarity

	return field
}

func (field *Text) SetTermVector(termVector TermVector) *Text {
	field.termVector = &termVector

	return field
}

func (field *Text) Fields(fields ...*Field) *Text {
	field.fields = append(field.fields, fields...)

	return field
}

func (field *Text) Meta(metas ...*Meta) *Text {
	field.meta = append(field.meta, metas...)

	return field
}

func (field *Text) GetType() Type {
	return TypeText
}

func (field *Text) Source() (interface{}, error) {
	// {
	//  "type":  "text",
	//  "analyzer":"my_analyzer",
	//  "search_analyzer":"my_stop_analyzer",
	//  "search_quote_analyzer":"my_analyzer",
	//  "boost": 2,
	//  "eager_global_ordinals": true,
	//  "fieldData": true,
	//  "fielddata_frequency_filter": {
	//    "min": 0.001,
	//    "max": 0.1,
	//    "min_segment_size": 500
	//  },
	//  "fields": {
	//    "raw": {
	//      "type":  "keyword"
	//    }
	//  },
	//  "index": true,
	//  "index_options": "offsets",
	//  "index_prefixes": {
	//    "min_chars" : 1,
	//    "max_chars" : 10
	//  },
	//  "index_phrases": true,
	//  "norms": false,
	//  "position_increment_gap": 0,
	//  "store": true,
	//  "similarity": "boolean",
	//  "term_vector": "with_positions_offsets",
	//  "meta": {
	//    "unit": "ms"
	//  }
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	if field.fieldData != nil {
		source["fieldData"] = *field.fieldData
	}

	if field.eagerGlobalOrdinals != nil {
		source["eager_global_ordinals"] = *field.eagerGlobalOrdinals
	}

	if field.indexPhrases != nil {
		source["index_phrases"] = *field.indexPhrases
	}

	if field.index != nil {
		source["index"] = *field.index
	}

	if field.norms != nil {
		source["norms"] = *field.norms
	}

	if field.store != nil {
		source["store"] = *field.store
	}

	if field.positionIncrementGap != nil {
		source["position_increment_gap"] = *field.positionIncrementGap
	}

	if field.boost != nil {
		source["boost"] = *field.boost
	}

	if field.analyzer != nil {
		source["analyzer"] = *field.analyzer
	}

	if field.searchAnalyzer != nil {
		if field.analyzer == nil {
			return nil, ErrorSearchAnalyzer
		}

		source["search_analyzer"] = *field.searchAnalyzer
	}

	if field.searchQuoteAnalyzer != nil {
		if field.analyzer == nil || field.searchAnalyzer == nil {
			return nil, ErrorSearchQuoteAnalyzer
		}

		source["search_quote_analyzer"] = *field.searchQuoteAnalyzer
	}

	if field.fieldDataFrequencyFilter != nil {
		body, err := field.fieldDataFrequencyFilter.Source()
		if err != nil {
			return nil, err
		}

		source["fielddata_frequency_filter"] = body
	}

	if field.indexOptions != nil {
		source["index_options"] = *field.indexOptions
	}

	if field.indexPrefixes != nil {
		body, err := field.indexPrefixes.Source()
		if err != nil {
			return nil, err
		}

		source["index_prefixes"] = body
	}

	if field.similarity != nil {
		source["similarity"] = *field.similarity
	}

	if field.termVector != nil {
		source["term_vector"] = *field.termVector
	}

	if len(field.fields) > 0 {
		fields := map[FieldName]interface{}{}

		for _, field := range field.fields {
			s, err := field.Source()
			if err != nil {
				return nil, err
			}

			fields[field.Name()] = s
		}

		source["fields"] = fields
	}

	if len(field.meta) > 0 {
		body := map[string]interface{}{}

		for _, meta := range field.meta {
			body[meta.name] = meta.value
		}

		source["meta"] = body
	}

	return source, nil
}
