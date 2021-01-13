package settings

const (
    NormalizerNameLowercase NormalizerName = "lowercase"

    NormalizerTypeCustom NormalizerType = "custom"
)

type NormalizerType string
type NormalizerName string

type Normalizer interface {
    Type() NormalizerType
    Name() NormalizerName
    Source() (interface{}, error)
}

// see https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-normalizers.html
type NormalizerCustom struct {
    name       NormalizerName
    charFilter []CharFilterName
    filter     []FilterName
}

func NewNormalizerCustom(name string) *NormalizerCustom {
    return &NormalizerCustom{name: NormalizerName(name)}
}

func (normalizer *NormalizerCustom) AddCharFilter(charFilter ...CharFilterName) *NormalizerCustom {
    normalizer.charFilter = append(normalizer.charFilter, charFilter...)
    
    return normalizer
}

func (normalizer *NormalizerCustom) AddFilter(filter ...FilterName) *NormalizerCustom {
    normalizer.filter = append(normalizer.filter, filter...)
    
    return normalizer
}

func (normalizer *NormalizerCustom) Type() NormalizerType {
    return NormalizerTypeCustom
}

func (normalizer *NormalizerCustom) Name() NormalizerName {
    return normalizer.name
}

func (normalizer *NormalizerCustom) Source() (interface{}, error) {
    source := map[string]interface{}{
        "type": normalizer.Type(),
    }

    if len(normalizer.charFilter) > 0 {
        source["char_filter"] = normalizer.charFilter
    }

    if len(normalizer.filter) > 0 {
        source["filter"] = normalizer.filter
    }

    return source, nil
}
