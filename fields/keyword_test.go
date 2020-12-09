package fields

import (
    "github.com/diez37/elastic-migrate/settings"
    "reflect"
    "testing"
)

func TestKeyword_Fields(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        fields []*Field
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "set",
            fields: fields{},
            args: args{fields: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want: &Keyword{
                fields: []*Field{
                    {name: "name", field: &Keyword{}},
                },
            },
        },
        {
            name: "add",
            fields: fields{fields: []*Field{
                {name: "3gram", field: &Text{}},
            }},
            args: args{fields: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want: &Keyword{
                fields: []*Field{
                    {name: "3gram", field: &Text{}},
                    {name: "name", field: &Keyword{}},
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.Fields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Fields() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_GetType(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    tests := []struct {
        name   string
        fields fields
        want   Type
    }{
        {
            name:   "type keyword",
            fields: fields{},
            want:   TypeKeyword,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_Meta(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        metas []*Meta
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &Keyword{},
        },
        {
            name:   "set meta",
            fields: fields{},
            args: args{metas: []*Meta{
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
            want: &Keyword{meta: []*Meta{
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
        },
        {
            name: "add meta",
            fields: fields{meta: []*Meta{
                {
                    name:  "author",
                    value: "diez37",
                },
                {
                    name:  "year",
                    value: 2020,
                },
            }},
            args: args{metas: []*Meta{
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
            want: &Keyword{meta: []*Meta{
                {
                    name:  "author",
                    value: "diez37",
                },
                {
                    name:  "year",
                    value: 2020,
                },
                {
                    name:  "version",
                    value: "0.0.1",
                },
                {
                    name:  "boost",
                    value: 3.5,
                },
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.Meta(tt.args.metas...); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Meta() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetBoost(t *testing.T) {
    initBoost := 9.0
    setBoost := 5.4

    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        boost float64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "set boost",
            fields: fields{},
            args:   args{boost: setBoost},
            want:   &Keyword{boost: &setBoost},
        },
        {
            name:   "change boost",
            fields: fields{boost: &initBoost},
            args:   args{boost: setBoost},
            want:   &Keyword{boost: &setBoost},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetBoost() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetDocValues(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        docValues bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{docValues: true},
            want:   &Keyword{docValues: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{docValues: false},
            want:   &Keyword{docValues: &testFalse},
        },
        {
            name:   "change",
            fields: fields{docValues: &testTrue},
            args:   args{docValues: false},
            want:   &Keyword{docValues: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetDocValues(tt.args.docValues); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetDocValues() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetEagerGlobalOrdinals(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        eagerGlobalOrdinals bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{eagerGlobalOrdinals: true},
            want:   &Keyword{eagerGlobalOrdinals: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{eagerGlobalOrdinals: false},
            want:   &Keyword{eagerGlobalOrdinals: &testFalse},
        },
        {
            name:   "change",
            fields: fields{eagerGlobalOrdinals: &testTrue},
            args:   args{eagerGlobalOrdinals: false},
            want:   &Keyword{eagerGlobalOrdinals: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetEagerGlobalOrdinals(tt.args.eagerGlobalOrdinals); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetEagerGlobalOrdinals() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetIgnoreAbove(t *testing.T) {
    initValue := 9
    setValue := 5

    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        ignoreAbove int
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "set boost",
            fields: fields{},
            args:   args{ignoreAbove: setValue},
            want:   &Keyword{ignoreAbove: &setValue},
        },
        {
            name:   "change boost",
            fields: fields{ignoreAbove: &initValue},
            args:   args{ignoreAbove: setValue},
            want:   &Keyword{ignoreAbove: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetIgnoreAbove(tt.args.ignoreAbove); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIgnoreAbove() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetIndex(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        index bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{index: true},
            want:   &Keyword{index: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{index: false},
            want:   &Keyword{index: &testFalse},
        },
        {
            name:   "change",
            fields: fields{index: &testTrue},
            args:   args{index: false},
            want:   &Keyword{index: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndex() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetIndexOptions(t *testing.T) {
    initValue := IndexOptionOffsets
    setValue := IndexOptionDocs

    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        indexOptions IndexOption
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{indexOptions: IndexOptionDocs},
            want:   &Keyword{indexOptions: &setValue},
        },
        {
            name:   "change",
            fields: fields{indexOptions: &initValue},
            args:   args{indexOptions: IndexOptionDocs},
            want:   &Keyword{indexOptions: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetIndexOptions(tt.args.indexOptions); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetIndexOptions() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetNormalizer(t *testing.T) {
    setValue := settings.NormalizerNameLowercase

    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        normalizer settings.NormalizerName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{normalizer: settings.NormalizerNameLowercase},
            want:   &Keyword{normalizer: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetNormalizer(tt.args.normalizer); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNormalizer() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetNorms(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        norms bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{norms: true},
            want:   &Keyword{norms: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{norms: false},
            want:   &Keyword{norms: &testFalse},
        },
        {
            name:   "change",
            fields: fields{norms: &testTrue},
            args:   args{norms: false},
            want:   &Keyword{norms: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetNorms(tt.args.norms); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNorms() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetNullValue(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        nullValue interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "empty",
            fields: fields{},
            args:   args{},
            want:   &Keyword{nullValue: nil},
        },
        {
            name:   "int",
            fields: fields{},
            args:   args{nullValue: 1},
            want:   &Keyword{nullValue: 1},
        },
        {
            name:   "float",
            fields: fields{},
            args:   args{nullValue: 4.5},
            want:   &Keyword{nullValue: 4.5},
        },
        {
            name:   "string",
            fields: fields{},
            args:   args{nullValue: "test"},
            want:   &Keyword{nullValue: "test"},
        },
        {
            name:   "bool",
            fields: fields{},
            args:   args{nullValue: true},
            want:   &Keyword{nullValue: true},
        },
        {
            name:   "change",
            fields: fields{nullValue: "test"},
            args:   args{nullValue: true},
            want:   &Keyword{nullValue: true},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetNullValue(tt.args.nullValue); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetNullValue() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetSimilarity(t *testing.T) {
    initValue := settings.SimilarityName("test1")
    setValue := settings.SimilarityName("test2")

    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        similarity settings.SimilarityName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "set",
            fields: fields{},
            args:   args{similarity: setValue},
            want:   &Keyword{similarity: &setValue},
        },
        {
            name:   "change",
            fields: fields{similarity: &initValue},
            args:   args{similarity: setValue},
            want:   &Keyword{similarity: &setValue},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetSimilarity(tt.args.similarity); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetSimilarity() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetSplitQueriesOnWhitespace(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        splitQueriesOnWhitespace bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{splitQueriesOnWhitespace: true},
            want:   &Keyword{splitQueriesOnWhitespace: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{splitQueriesOnWhitespace: false},
            want:   &Keyword{splitQueriesOnWhitespace: &testFalse},
        },
        {
            name:   "change",
            fields: fields{splitQueriesOnWhitespace: &testTrue},
            args:   args{splitQueriesOnWhitespace: false},
            want:   &Keyword{splitQueriesOnWhitespace: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetSplitQueriesOnWhitespace(tt.args.splitQueriesOnWhitespace); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetSplitQueriesOnWhitespace() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_SetStore(t *testing.T) {
    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
    }
    type args struct {
        store bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   *Keyword
    }{
        {
            name:   "true",
            fields: fields{},
            args:   args{store: true},
            want:   &Keyword{store: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{store: false},
            want:   &Keyword{store: &testFalse},
        },
        {
            name:   "change",
            fields: fields{store: &testTrue},
            args:   args{store: false},
            want:   &Keyword{store: &testFalse},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
            }
            if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SetStore() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyword_Source(t *testing.T) {
    boost := 2.5
    ignoreAbove := 2
    indexOptions := IndexOptionFreqs
    similarity := settings.SimilarityName("test1")
    normalizer := settings.NormalizerNameLowercase

    type fields struct {
        docValues                *bool
        store                    *bool
        index                    *bool
        eagerGlobalOrdinals      *bool
        norms                    *bool
        splitQueriesOnWhitespace *bool
        boost                    *float64
        ignoreAbove              *int
        indexOptions             *IndexOption
        similarity               *settings.SimilarityName
        normalizer               *settings.NormalizerName
        nullValue                interface{}
        fields                   []*Field
        meta                     []*Meta
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
                "type": TypeKeyword,
            },
            wantErr: false,
        },
        {
            name:   "docValues",
            fields: fields{docValues: &testTrue},
            want: map[string]interface{}{
                "type":       TypeKeyword,
                "doc_values": true,
            },
        },
        {
            name:   "store",
            fields: fields{store: &testTrue},
            want: map[string]interface{}{
                "type":  TypeKeyword,
                "store": true,
            },
        },
        {
            name:   "index",
            fields: fields{index: &testTrue},
            want: map[string]interface{}{
                "type":  TypeKeyword,
                "index": true,
            },
        },
        {
            name:   "eagerGlobalOrdinals",
            fields: fields{eagerGlobalOrdinals: &testTrue},
            want: map[string]interface{}{
                "type":                  TypeKeyword,
                "eager_global_ordinals": true,
            },
        },
        {
            name:   "norms",
            fields: fields{norms: &testTrue},
            want: map[string]interface{}{
                "type":  TypeKeyword,
                "norms": true,
            },
        },
        {
            name:   "splitQueriesOnWhitespace",
            fields: fields{splitQueriesOnWhitespace: &testTrue},
            want: map[string]interface{}{
                "type":                        TypeKeyword,
                "split_queries_on_whitespace": true,
            },
        },
        {
            name:   "boost",
            fields: fields{boost: &boost},
            want: map[string]interface{}{
                "type":  TypeKeyword,
                "boost": boost,
            },
        },
        {
            name:   "ignoreAbove",
            fields: fields{ignoreAbove: &ignoreAbove},
            want: map[string]interface{}{
                "type":         TypeKeyword,
                "ignore_above": ignoreAbove,
            },
        },
        {
            name:   "indexOptions",
            fields: fields{indexOptions: &indexOptions},
            want: map[string]interface{}{
                "type":          TypeKeyword,
                "index_options": indexOptions,
            },
        },
        {
            name:   "similarity",
            fields: fields{similarity: &similarity},
            want: map[string]interface{}{
                "type":       TypeKeyword,
                "similarity": settings.SimilarityName("test1"),
            },
        },
        {
            name:   "normalizer",
            fields: fields{normalizer: &normalizer},
            want: map[string]interface{}{
                "type":       TypeKeyword,
                "normalizer": normalizer,
            },
        },
        {
            name:   "nullValue",
            fields: fields{nullValue: 2.7},
            want: map[string]interface{}{
                "type":       TypeKeyword,
                "null_value": 2.7,
            },
        },
        {
            name: "meta",
            fields: fields{meta: []*Meta{
                {
                    name:  "author",
                    value: "diez37",
                },
                {
                    name:  "version",
                    value: "0.0.1",
                },
            }},
            want: map[string]interface{}{
                "type": TypeKeyword,
                "meta": map[string]interface{}{
                    "author":  "diez37",
                    "version": "0.0.1",
                },
            },
        },
        {
            name: "fields",
            fields: fields{fields: []*Field{
                {name: "name", field: &Keyword{}},
            }},
            want: map[string]interface{}{
                "type": TypeKeyword,
                "fields": map[string]interface{}{
                    "name": map[string]interface{}{
                        "type": TypeKeyword,
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &Keyword{
                docValues:                tt.fields.docValues,
                store:                    tt.fields.store,
                index:                    tt.fields.index,
                eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
                norms:                    tt.fields.norms,
                splitQueriesOnWhitespace: tt.fields.splitQueriesOnWhitespace,
                boost:                    tt.fields.boost,
                ignoreAbove:              tt.fields.ignoreAbove,
                indexOptions:             tt.fields.indexOptions,
                similarity:               tt.fields.similarity,
                normalizer:               tt.fields.normalizer,
                nullValue:                tt.fields.nullValue,
                fields:                   tt.fields.fields,
                meta:                     tt.fields.meta,
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

func TestNewKeyword(t *testing.T) {
    tests := []struct {
        name string
        want *Keyword
    }{
        {
            name: "new",
            want: &Keyword{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewKeyword(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewKeyword() = %v, want %v", got, tt.want)
            }
        })
    }
}
