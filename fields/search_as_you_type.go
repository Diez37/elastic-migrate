package fields

import (
	"errors"
	"github.com/diez37/elastic-migrate/settings"
)

var ErrorSearchAnalyzer = errors.New("analyzer must be set when search_analyzer is set")
var ErrorShingleSizeMaximum = errors.New("max_shingle_size maximum value = 4")
var ErrorShingleSizeMinimum = errors.New("max_shingle_size minimum value = 2")

type SearchAsYouType struct {
	name                FieldName
	index               *bool
	norms               *bool
	store               *bool
	maxShingleSize      *int
	analyzer            *settings.AnalyzerName
	searchAnalyzer      *settings.AnalyzerName
	searchQuoteAnalyzer *settings.AnalyzerName
	indexOptions        *IndexOption
	similarity          *settings.SimilarityName
	termVector          *TermVector
}

func NewSearchAsYouType(name string) *SearchAsYouType {
	return &SearchAsYouType{name: FieldName(name)}
}

func (field *SearchAsYouType) Name() FieldName {
	return field.name
}

func (field *SearchAsYouType) SetMaxShingleSize(maxShingleSize int) *SearchAsYouType {
	field.maxShingleSize = &maxShingleSize

	return field
}

func (field *SearchAsYouType) SetIndex(index bool) *SearchAsYouType {
	field.index = &index

	return field
}

func (field *SearchAsYouType) SetNorms(norms bool) *SearchAsYouType {
	field.norms = &norms

	return field
}

func (field *SearchAsYouType) SetStore(store bool) *SearchAsYouType {
	field.store = &store

	return field
}

func (field *SearchAsYouType) SetAnalyzer(analyzer settings.AnalyzerName) *SearchAsYouType {
	field.analyzer = &analyzer

	return field
}

func (field *SearchAsYouType) SetSearchAnalyzer(searchAnalyzer settings.AnalyzerName) *SearchAsYouType {
	field.searchAnalyzer = &searchAnalyzer

	return field
}

func (field *SearchAsYouType) SetSearchQuoteAnalyzer(searchQuoteAnalyzer settings.AnalyzerName) *SearchAsYouType {
	field.searchQuoteAnalyzer = &searchQuoteAnalyzer

	return field
}

func (field *SearchAsYouType) SetIndexOptions(indexOptions IndexOption) *SearchAsYouType {
	field.indexOptions = &indexOptions

	return field
}

func (field *SearchAsYouType) SetSimilarity(similarity settings.SimilarityName) *SearchAsYouType {
	field.similarity = &similarity

	return field
}

func (field *SearchAsYouType) SetTermVector(termVector TermVector) *SearchAsYouType {
	field.termVector = &termVector

	return field
}

func (field *SearchAsYouType) GetType() Type {
	return TypeSearchAsYouType
}

func (field *SearchAsYouType) Source() (interface{}, error) {
	// {
	//  "type": "search_as_you_type",
	//  "max_shingle_size": 3,
	//  "analyzer": "my_analyzer",
	//  "search_analyzer": "my_stop_analyzer",
	//  "search_quote_analyzer": "my_analyzer",
	//  "index": true,
	//  "index_options": "offsets",
	//  "norms": false,
	//  "store": true,
	//  "similarity": "boolean",
	//  "term_vector": "with_positions_offsets"
	// }

	source := map[string]interface{}{}

	source["type"] = field.GetType()

	if field.maxShingleSize != nil {
		if *field.maxShingleSize > ShingleSizeMaximum {
			return nil, ErrorShingleSizeMaximum
		}

		if *field.maxShingleSize < ShingleSizeMinimum {
			return nil, ErrorShingleSizeMinimum
		}

		source["max_shingle_size"] = *field.maxShingleSize
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

	if field.indexOptions != nil {
		source["index_options"] = *field.indexOptions
	}

	if field.similarity != nil {
		source["similarity"] = *field.similarity
	}

	if field.termVector != nil {
		source["term_vector"] = *field.termVector
	}

	return source, nil
}
