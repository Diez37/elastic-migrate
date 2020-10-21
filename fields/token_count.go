package fields

import "github.com/diez37/elastic-migrate/settings"

type TokenCount struct {
    enablePositionIncrements *bool
    docValues                *bool
    index                    *bool
    store                    *bool
    boost                    *float64
    analyzer                 *settings.AnalyzerName
    nullValue                interface{}
}

func NewTokenCount() *TokenCount {
    return &TokenCount{}
}

func (field *TokenCount) SetEnablePositionIncrements(enablePositionIncrements bool) *TokenCount {
    field.enablePositionIncrements = &enablePositionIncrements

    return field
}

func (field *TokenCount) SetDocValues(docValues bool) *TokenCount {
    field.docValues = &docValues

    return field
}

func (field *TokenCount) SetIndex(index bool) *TokenCount {
    field.index = &index

    return field
}

func (field *TokenCount) SetStore(store bool) *TokenCount {
    field.store = &store

    return field
}

func (field *TokenCount) SetBoost(boost float64) *TokenCount {
    field.boost = &boost

    return field
}

func (field *TokenCount) SetAnalyzer(analyzer settings.AnalyzerName) *TokenCount {
    field.analyzer = &analyzer

    return field
}

func (field *TokenCount) SetNullValue(nullValue interface{}) *TokenCount {
    field.nullValue = nullValue

    return field
}

func (field *TokenCount) GetType() Type {
    return TypeTokenCount
}

func (field *TokenCount) Source() (interface{}, error) {
    // {
    //  "type":     "token_count",
    //  "analyzer": "standard",
    //  "enable_position_increments": true,
    //  "boost": 2,
    //  "doc_values": false,
    //  "index": true,
    //  "null_value": "NULL",
    //  "store": true
    // }

    source := map[string]interface{}{}

    source["type"] = field.GetType()

    if field.analyzer != nil {
        source["analyzer"] = *field.analyzer
    }

    if field.enablePositionIncrements != nil {
        source["enable_position_increments"] = *field.enablePositionIncrements
    }

    if field.boost != nil {
        source["boost"] = *field.boost
    }

    if field.docValues != nil {
        source["doc_values"] = *field.docValues
    }

    if field.index != nil {
        source["index"] = *field.index
    }

    if field.nullValue != nil {
        source["null_value"] = field.nullValue
    }

    if field.store != nil {
        source["store"] = *field.store
    }

    return source, nil
}
