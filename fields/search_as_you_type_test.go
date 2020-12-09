package fields

import (
    "github.com/diez37/elastic-migrate/settings"
    "reflect"
    "testing"
)

func TestNewSearchAsYouType(t *testing.T) {
    tests := []struct {
        name string
        want *SearchAsYouType
    }{
        {
            name: "new",
            want: &SearchAsYouType{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewSearchAsYouType(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewSearchAsYouType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_GetType(t *testing.T) {
    type fields struct {
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
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type search as you type",
            fields: fields{},
            want:   TypeSearchAsYouType,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetAnalyzer(t *testing.T) {
    initAnalyzerName := settings.AnalyzerNameStandard
    changeAnalyzerName := settings.AnalyzerNameSimple

    type fields struct {
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
    type args struct {
        analyzer settings.AnalyzerName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{analyzer: initAnalyzerName},
            want:   &SearchAsYouType{analyzer: &initAnalyzerName},
        },
        {
            name:   "change",
            fields: fields{analyzer: &changeAnalyzerName},
            args:   args{analyzer: initAnalyzerName},
            want:   &SearchAsYouType{analyzer: &initAnalyzerName},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetAnalyzer(tt.args.analyzer); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetAnalyzer() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetIndex(t *testing.T) {
    type fields struct {
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
    type args struct {
        index bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{index: true},
            want:   &SearchAsYouType{index: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{index: false},
            want:   &SearchAsYouType{index: &testFalse},
        },
        {
            name:   "change",
            fields: fields{index: &testTrue},
            args:   args{index: false},
            want:   &SearchAsYouType{index: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndex() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetIndexOptions(t *testing.T) {
    initValue := IndexOptionOffsets
    setValue := IndexOptionDocs

    type fields struct {
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
    type args struct {
        indexOptions IndexOption
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{indexOptions: IndexOptionDocs},
            want:   &SearchAsYouType{indexOptions: &setValue},
        },
        {
            name:   "change",
            fields: fields{indexOptions: &initValue},
            args:   args{indexOptions: IndexOptionDocs},
            want:   &SearchAsYouType{indexOptions: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetIndexOptions(tt.args.indexOptions); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndexOptions() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetMaxShingleSize(t *testing.T) {
    initMaxShingleSize := 2
    setMaxShingleSize := 4

    type fields struct {
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
    type args struct {
        maxShingleSize int
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{maxShingleSize: initMaxShingleSize},
            want:   &SearchAsYouType{maxShingleSize: &initMaxShingleSize},
        },
        {
            name:   "change",
            fields: fields{maxShingleSize: &initMaxShingleSize},
            args:   args{maxShingleSize: setMaxShingleSize},
            want:   &SearchAsYouType{maxShingleSize: &setMaxShingleSize},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetMaxShingleSize(tt.args.maxShingleSize); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetMaxShingleSize() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetNorms(t *testing.T) {
    type fields struct {
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
    type args struct {
        norms bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{norms: true},
            want:   &SearchAsYouType{norms: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{norms: false},
            want:   &SearchAsYouType{norms: &testFalse},
        },
        {
            name:   "change",
            fields: fields{norms: &testTrue},
            args:   args{norms: false},
            want:   &SearchAsYouType{norms: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetNorms(tt.args.norms); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNorms() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetSearchAnalyzer(t *testing.T) {
    initAnalyzerName := settings.AnalyzerNameStandard
    changeAnalyzerName := settings.AnalyzerNameSimple

    type fields struct {
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
    type args struct {
        searchAnalyzer settings.AnalyzerName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{searchAnalyzer: initAnalyzerName},
            want:   &SearchAsYouType{searchAnalyzer: &initAnalyzerName},
        },
        {
            name:   "change",
            fields: fields{searchAnalyzer: &changeAnalyzerName},
            args:   args{searchAnalyzer: initAnalyzerName},
            want:   &SearchAsYouType{searchAnalyzer: &initAnalyzerName},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetSearchAnalyzer(tt.args.searchAnalyzer); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetSearchAnalyzer() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetSearchQuoteAnalyzer(t *testing.T) {
    initAnalyzerName := settings.AnalyzerNameStandard
    changeAnalyzerName := settings.AnalyzerNameSimple

    type fields struct {
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
    type args struct {
        searchQuoteAnalyzer settings.AnalyzerName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{searchQuoteAnalyzer: initAnalyzerName},
            want:   &SearchAsYouType{searchQuoteAnalyzer: &initAnalyzerName},
        },
        {
            name:   "change",
            fields: fields{searchQuoteAnalyzer: &changeAnalyzerName},
            args:   args{searchQuoteAnalyzer: initAnalyzerName},
            want:   &SearchAsYouType{searchQuoteAnalyzer: &initAnalyzerName},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetSearchQuoteAnalyzer(tt.args.searchQuoteAnalyzer); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetSearchQuoteAnalyzer() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetSimilarity(t *testing.T) {
    initValue := settings.SimilarityName("test1")
    setValue := settings.SimilarityName("test2")

    type fields struct {
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
    type args struct {
        similarity settings.SimilarityName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{similarity: setValue},
            want:   &SearchAsYouType{similarity: &setValue},
        },
        {
            name:   "change",
            fields: fields{similarity: &initValue},
            args:   args{similarity: setValue},
            want:   &SearchAsYouType{similarity: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetSimilarity(tt.args.similarity); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetSimilarity() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetStore(t *testing.T) {
    type fields struct {
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
    type args struct {
        store bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{store: true},
            want:   &SearchAsYouType{store: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{store: false},
            want:   &SearchAsYouType{store: &testFalse},
        },
        {
            name:   "change",
            fields: fields{store: &testTrue},
            args:   args{store: false},
            want:   &SearchAsYouType{store: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetStore() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_SetTermVector(t *testing.T) {
    initValue := TermVectorYes
    setValue := TermVectorWithPositions

    type fields struct {
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
    type args struct {
        termVector TermVector
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *SearchAsYouType
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{termVector: setValue},
            want:   &SearchAsYouType{termVector: &setValue},
        },
        {
            name:   "change",
            fields: fields{termVector: &initValue},
            args:   args{termVector: setValue},
            want:   &SearchAsYouType{termVector: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            if got := field.SetTermVector(tt.args.termVector); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetTermVector() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSearchAsYouType_Source(t *testing.T) {
    initMaxShingleSize := 4
    initAnalyzerName := settings.AnalyzerNameStandard
    initIndexOptionValue := IndexOptionOffsets
    initSimilarityOptionValue := settings.SimilarityName("test1")
    initTermVectorValue := TermVectorYes

    type fields struct {
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
    tests := []struct {
        name    string
        fields  fields
        want    interface{}
        wantErr bool
    }{
        {
            name:   "empty",
            fields: fields{},
            want: map[string]interface{}{
                "type": TypeSearchAsYouType,
            },
            wantErr: false,
        },
        {
            name:   "maxShingleSize",
            fields: fields{maxShingleSize: &initMaxShingleSize},
            want: map[string]interface{}{
                "type":             TypeSearchAsYouType,
                "max_shingle_size": initMaxShingleSize,
            },
            wantErr: false,
        },
        {
            name:   "analyzer",
            fields: fields{analyzer: &initAnalyzerName},
            want: map[string]interface{}{
                "type":     TypeSearchAsYouType,
                "analyzer": initAnalyzerName,
            },
            wantErr: false,
        },
        {
            name: "searchAnalyzer",
            fields: fields{
                searchAnalyzer: &initAnalyzerName,
                analyzer:       &initAnalyzerName,
            },
            want: map[string]interface{}{
                "type":            TypeSearchAsYouType,
                "search_analyzer": initAnalyzerName,
                "analyzer":        initAnalyzerName,
            },
            wantErr: false,
        },
        {
            name: "searchQuoteAnalyzer",
            fields: fields{
                searchQuoteAnalyzer: &initAnalyzerName,
                searchAnalyzer:      &initAnalyzerName,
                analyzer:            &initAnalyzerName,
            },
            want: map[string]interface{}{
                "type":                  TypeSearchAsYouType,
                "search_quote_analyzer": initAnalyzerName,
                "search_analyzer":       initAnalyzerName,
                "analyzer":              initAnalyzerName,
            },
            wantErr: false,
        },
        {
            name:   "index",
            fields: fields{index: &testTrue},
            want: map[string]interface{}{
                "type":  TypeSearchAsYouType,
                "index": true,
            },
            wantErr: false,
        },
        {
            name:   "indexOptions",
            fields: fields{indexOptions: &initIndexOptionValue},
            want: map[string]interface{}{
                "type":          TypeSearchAsYouType,
                "index_options": initIndexOptionValue,
            },
            wantErr: false,
        },
        {
            name:   "norms",
            fields: fields{norms: &testTrue},
            want: map[string]interface{}{
                "type":  TypeSearchAsYouType,
                "norms": true,
            },
            wantErr: false,
        },
        {
            name:   "store",
            fields: fields{store: &testTrue},
            want: map[string]interface{}{
                "type":  TypeSearchAsYouType,
                "store": true,
            },
            wantErr: false,
        },
        {
            name:   "similarity",
            fields: fields{similarity: &initSimilarityOptionValue},
            want: map[string]interface{}{
                "type":       TypeSearchAsYouType,
                "similarity": settings.SimilarityName("test1"),
            },
            wantErr: false,
        },
        {
            name:   "termVector",
            fields: fields{termVector: &initTermVectorValue},
            want: map[string]interface{}{
                "type":        TypeSearchAsYouType,
                "term_vector": initTermVectorValue,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &SearchAsYouType{
                index:               tt.fields.index,
                norms:               tt.fields.norms,
                store:               tt.fields.store,
                maxShingleSize:      tt.fields.maxShingleSize,
                analyzer:            tt.fields.analyzer,
                searchAnalyzer:      tt.fields.searchAnalyzer,
                searchQuoteAnalyzer: tt.fields.searchQuoteAnalyzer,
                indexOptions:        tt.fields.indexOptions,
                similarity:          tt.fields.similarity,
                termVector:          tt.fields.termVector,
            }
            got, err := field.Source()
            if (err != nil) != tt.wantErr {
                t.Errorf("Source() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Source() got = %v, want %v", got, tt.want)
            }
        })
    }
}
