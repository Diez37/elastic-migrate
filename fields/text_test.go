package fields

import (
	"github.com/diez37/elastic-migrate/settings"
	"reflect"
	"testing"
)

func TestNewText(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Text
	}{
		{
			args: args{name: "test"},
			want: &Text{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewText(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_Fields(t *testing.T) {
	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args: args{fields: []*Field{
				{name: "name", field: &Keyword{}},
			}},
			want: &Text{
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
			want: &Text{
				fields: []*Field{
					{name: "3gram", field: &Text{}},
					{name: "name", field: &Keyword{}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.Fields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_GetType(t *testing.T) {
	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "type text",
			fields: fields{},
			want:   TypeText,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_Meta(t *testing.T) {
	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "empty",
			fields: fields{},
			args:   args{},
			want:   &Text{},
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
			want: &Text{meta: []*Meta{
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
			want: &Text{meta: []*Meta{
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
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.Meta(tt.args.metas...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Meta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetAnalyzer(t *testing.T) {
	initAnalyzerName := settings.AnalyzerNameStandard
	changeAnalyzerName := settings.AnalyzerNameSimple

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		analyzer settings.AnalyzerName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{analyzer: initAnalyzerName},
			want:   &Text{analyzer: &initAnalyzerName},
		},
		{
			name:   "change",
			fields: fields{analyzer: &changeAnalyzerName},
			args:   args{analyzer: initAnalyzerName},
			want:   &Text{analyzer: &initAnalyzerName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetAnalyzer(tt.args.analyzer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetBoost(t *testing.T) {
	initBoost := 9.0
	setBoost := 5.4

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "set boost",
			fields: fields{},
			args:   args{boost: setBoost},
			want:   &Text{boost: &setBoost},
		},
		{
			name:   "change boost",
			fields: fields{boost: &initBoost},
			args:   args{boost: setBoost},
			want:   &Text{boost: &setBoost},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetBoost(tt.args.boost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetEagerGlobalOrdinals(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{eagerGlobalOrdinals: true},
			want:   &Text{eagerGlobalOrdinals: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{eagerGlobalOrdinals: false},
			want:   &Text{eagerGlobalOrdinals: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{eagerGlobalOrdinals: &boolTrue},
			args:   args{eagerGlobalOrdinals: false},
			want:   &Text{eagerGlobalOrdinals: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetEagerGlobalOrdinals(tt.args.eagerGlobalOrdinals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEagerGlobalOrdinals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetFielddata(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		fielddata bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{fielddata: true},
			want:   &Text{fieldData: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{fielddata: false},
			want:   &Text{fieldData: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{fielddata: &boolTrue},
			args:   args{fielddata: false},
			want:   &Text{fieldData: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetFieldData(tt.args.fielddata); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFieldData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetFielddataFrequencyFilter(t *testing.T) {
	min := 1.5

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		fielddataFrequencyFilter *FieldDataFrequencyFilter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{fielddataFrequencyFilter: &FieldDataFrequencyFilter{}},
			want:   &Text{fieldDataFrequencyFilter: &FieldDataFrequencyFilter{}},
		},
		{
			name:   "change",
			fields: fields{fielddataFrequencyFilter: &FieldDataFrequencyFilter{}},
			args:   args{fielddataFrequencyFilter: &FieldDataFrequencyFilter{min: &min}},
			want:   &Text{fieldDataFrequencyFilter: &FieldDataFrequencyFilter{min: &min}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetFieldDataFrequencyFilter(tt.args.fielddataFrequencyFilter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFieldDataFrequencyFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetIndex(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{index: true},
			want:   &Text{index: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{index: false},
			want:   &Text{index: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{index: &boolTrue},
			args:   args{index: false},
			want:   &Text{index: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetIndexOptions(t *testing.T) {
	initValue := IndexOptionOffsets
	setValue := IndexOptionDocs

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{indexOptions: IndexOptionDocs},
			want:   &Text{indexOptions: &setValue},
		},
		{
			name:   "change",
			fields: fields{indexOptions: &initValue},
			args:   args{indexOptions: IndexOptionDocs},
			want:   &Text{indexOptions: &setValue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetIndexOptions(tt.args.indexOptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndexOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetIndexPhrases(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		indexPhrases bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{indexPhrases: true},
			want:   &Text{indexPhrases: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{indexPhrases: false},
			want:   &Text{indexPhrases: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{indexPhrases: &boolTrue},
			args:   args{indexPhrases: false},
			want:   &Text{indexPhrases: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetIndexPhrases(tt.args.indexPhrases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndexPhrases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetIndexPrefixes(t *testing.T) {
	min := 1

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		indexPrefixes *IndexPrefixes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{indexPrefixes: &IndexPrefixes{}},
			want:   &Text{indexPrefixes: &IndexPrefixes{}},
		},
		{
			name:   "change",
			fields: fields{indexPrefixes: &IndexPrefixes{}},
			args:   args{indexPrefixes: &IndexPrefixes{minChars: &min}},
			want:   &Text{indexPrefixes: &IndexPrefixes{minChars: &min}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetIndexPrefixes(tt.args.indexPrefixes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndexPrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetNorms(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{norms: true},
			want:   &Text{norms: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{norms: false},
			want:   &Text{norms: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{norms: &boolTrue},
			args:   args{norms: false},
			want:   &Text{norms: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetNorms(tt.args.norms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNorms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetPositionIncrementGap(t *testing.T) {
	initValue := 9
	setValue := 5

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		positionIncrementGap int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set boost",
			fields: fields{},
			args:   args{positionIncrementGap: setValue},
			want:   &Text{positionIncrementGap: &setValue},
		},
		{
			name:   "change boost",
			fields: fields{positionIncrementGap: &initValue},
			args:   args{positionIncrementGap: setValue},
			want:   &Text{positionIncrementGap: &setValue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetPositionIncrementGap(tt.args.positionIncrementGap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPositionIncrementGap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetSearchAnalyzer(t *testing.T) {
	initAnalyzerName := settings.AnalyzerNameStandard
	changeAnalyzerName := settings.AnalyzerNameSimple

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		searchAnalyzer settings.AnalyzerName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{searchAnalyzer: initAnalyzerName},
			want:   &Text{searchAnalyzer: &initAnalyzerName},
		},
		{
			name:   "change",
			fields: fields{searchAnalyzer: &changeAnalyzerName},
			args:   args{searchAnalyzer: initAnalyzerName},
			want:   &Text{searchAnalyzer: &initAnalyzerName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetSearchAnalyzer(tt.args.searchAnalyzer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSearchAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetSearchQuoteAnalyzer(t *testing.T) {
	initAnalyzerName := settings.AnalyzerNameStandard
	changeAnalyzerName := settings.AnalyzerNameSimple

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		searchQuoteAnalyzer settings.AnalyzerName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{searchQuoteAnalyzer: initAnalyzerName},
			want:   &Text{searchQuoteAnalyzer: &initAnalyzerName},
		},
		{
			name:   "change",
			fields: fields{searchQuoteAnalyzer: &changeAnalyzerName},
			args:   args{searchQuoteAnalyzer: initAnalyzerName},
			want:   &Text{searchQuoteAnalyzer: &initAnalyzerName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetSearchQuoteAnalyzer(tt.args.searchQuoteAnalyzer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSearchQuoteAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetSimilarity(t *testing.T) {
	initValue := settings.SimilarityName("test1")
	setValue := settings.SimilarityName("test2")

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{similarity: setValue},
			want:   &Text{similarity: &setValue},
		},
		{
			name:   "change",
			fields: fields{similarity: &initValue},
			args:   args{similarity: setValue},
			want:   &Text{similarity: &setValue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetSimilarity(tt.args.similarity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetStore(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
		want   *Text
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{store: true},
			want:   &Text{store: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{store: false},
			want:   &Text{store: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{store: &boolTrue},
			args:   args{store: false},
			want:   &Text{store: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetStore(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_SetTermVector(t *testing.T) {
	initValue := TermVectorYes
	setValue := TermVectorWithPositions
	termVectorNo := TermVectorNo
	termVectorYes := TermVectorYes
	termVectorWithPositions := TermVectorWithPositions
	termVectorWithOffsets := TermVectorWithOffsets
	termVectorWithPositionsOffsets := TermVectorWithPositionsOffsets
	termVectorWithPositionsPayloads := TermVectorWithPositionsPayloads
	termVectorWithPositionsOffsetsPayloads := TermVectorWithPositionsOffsetsPayloads

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
		fields                   []*Field
		meta                     []*Meta
	}
	type args struct {
		termVector TermVector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Text
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{termVector: setValue},
			want:   &Text{termVector: &setValue},
		},
		{
			name:   "change",
			fields: fields{termVector: &initValue},
			args:   args{termVector: setValue},
			want:   &Text{termVector: &setValue},
		},
		{
			name: "term_vector_no",
			args: args{termVector: TermVectorNo},
			want: &Text{termVector: &termVectorNo},
		},
		{
			name: "term_vector_yes",
			args: args{termVector: TermVectorYes},
			want: &Text{termVector: &termVectorYes},
		},
		{
			name: "term_vector_with_positions",
			args: args{termVector: TermVectorWithPositions},
			want: &Text{termVector: &termVectorWithPositions},
		},
		{
			name: "term_vector_with_offsets",
			args: args{termVector: TermVectorWithOffsets},
			want: &Text{termVector: &termVectorWithOffsets},
		},
		{
			name: "term_vector_with_positions_offsets",
			args: args{termVector: TermVectorWithPositionsOffsets},
			want: &Text{termVector: &termVectorWithPositionsOffsets},
		},
		{
			name: "term_vector_with_positions_payloads",
			args: args{termVector: TermVectorWithPositionsPayloads},
			want: &Text{termVector: &termVectorWithPositionsPayloads},
		},
		{
			name: "term_vector_with_positions_offsets_payloads",
			args: args{termVector: TermVectorWithPositionsOffsetsPayloads},
			want: &Text{termVector: &termVectorWithPositionsOffsetsPayloads},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
				fields:                   tt.fields.fields,
				meta:                     tt.fields.meta,
			}
			if got := field.SetTermVector(tt.args.termVector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTermVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_Source(t *testing.T) {
	boost := 2.5
	min := 1.5
	initAnalyzerName := settings.AnalyzerNameStandard
	initIndexOptionValue := IndexOptionOffsets
	initSimilarityOptionValue := settings.SimilarityName("test1")
	initTermVectorValue := TermVectorYes
	indexPrefixesMin := 1
	initPositionIncrementGap := 9
	boolTrue := true

	type fields struct {
		fielddata                *bool
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
		fielddataFrequencyFilter *FieldDataFrequencyFilter
		indexOptions             *IndexOption
		indexPrefixes            *IndexPrefixes
		similarity               *settings.SimilarityName
		termVector               *TermVector
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
				"type": TypeText,
			},
			wantErr: false,
		},
		{
			name:   "analyzer",
			fields: fields{analyzer: &initAnalyzerName},
			want: map[string]interface{}{
				"type":     TypeText,
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
				"type":            TypeText,
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
				"type":                  TypeText,
				"search_quote_analyzer": initAnalyzerName,
				"search_analyzer":       initAnalyzerName,
				"analyzer":              initAnalyzerName,
			},
			wantErr: false,
		},
		{
			name:   "boost",
			fields: fields{boost: &boost},
			want: map[string]interface{}{
				"type":  TypeText,
				"boost": boost,
			},
		},
		{
			name:   "eager global ordinals",
			fields: fields{eagerGlobalOrdinals: &boolTrue},
			want: map[string]interface{}{
				"type":                  TypeText,
				"eager_global_ordinals": true,
			},
		},
		{
			name:   "fieldData",
			fields: fields{fielddata: &boolTrue},
			want: map[string]interface{}{
				"type":      TypeText,
				"fieldData": true,
			},
		},
		{
			name:   "fieldData frequency filter",
			fields: fields{fielddataFrequencyFilter: &FieldDataFrequencyFilter{min: &min}},
			want: map[string]interface{}{
				"type": TypeText,
				"fielddata_frequency_filter": map[string]interface{}{
					"min": min,
				},
			},
		},
		{
			name: "fields",
			fields: fields{fields: []*Field{
				{name: "name", field: &Keyword{}},
			}},
			want: map[string]interface{}{
				"type": TypeText,
				"fields": map[FieldName]interface{}{
					"name": map[string]interface{}{
						"type": TypeKeyword,
					},
				},
			},
		},
		{
			name:   "index",
			fields: fields{index: &boolTrue},
			want: map[string]interface{}{
				"type":  TypeText,
				"index": true,
			},
			wantErr: false,
		},
		{
			name:   "indexOptions",
			fields: fields{indexOptions: &initIndexOptionValue},
			want: map[string]interface{}{
				"type":          TypeText,
				"index_options": initIndexOptionValue,
			},
			wantErr: false,
		},
		{
			name:   "indexPrefixes",
			fields: fields{indexPrefixes: &IndexPrefixes{minChars: &indexPrefixesMin}},
			want: map[string]interface{}{
				"type": TypeText,
				"index_prefixes": map[string]interface{}{
					"min_chars": indexPrefixesMin,
				},
			},
			wantErr: false,
		},
		{
			name:   "indexPhrases",
			fields: fields{indexPhrases: &boolTrue},
			want: map[string]interface{}{
				"type":          TypeText,
				"index_phrases": true,
			},
			wantErr: false,
		},
		{
			name:   "norms",
			fields: fields{norms: &boolTrue},
			want: map[string]interface{}{
				"type":  TypeText,
				"norms": true,
			},
			wantErr: false,
		},
		{
			name:   "positionIncrementGap",
			fields: fields{positionIncrementGap: &initPositionIncrementGap},
			want: map[string]interface{}{
				"type":                   TypeText,
				"position_increment_gap": initPositionIncrementGap,
			},
			wantErr: false,
		},
		{
			name:   "store",
			fields: fields{store: &boolTrue},
			want: map[string]interface{}{
				"type":  TypeText,
				"store": true,
			},
			wantErr: false,
		},
		{
			name:   "similarity",
			fields: fields{similarity: &initSimilarityOptionValue},
			want: map[string]interface{}{
				"type":       TypeText,
				"similarity": settings.SimilarityName("test1"),
			},
			wantErr: false,
		},
		{
			name:   "termVector",
			fields: fields{termVector: &initTermVectorValue},
			want: map[string]interface{}{
				"type":        TypeText,
				"term_vector": initTermVectorValue,
			},
			wantErr: false,
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
				"type": TypeText,
				"meta": map[string]interface{}{
					"author":  "diez37",
					"version": "0.0.1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Text{
				fieldData:                tt.fields.fielddata,
				eagerGlobalOrdinals:      tt.fields.eagerGlobalOrdinals,
				indexPhrases:             tt.fields.indexPhrases,
				index:                    tt.fields.index,
				norms:                    tt.fields.norms,
				store:                    tt.fields.store,
				positionIncrementGap:     tt.fields.positionIncrementGap,
				boost:                    tt.fields.boost,
				analyzer:                 tt.fields.analyzer,
				searchAnalyzer:           tt.fields.searchAnalyzer,
				searchQuoteAnalyzer:      tt.fields.searchQuoteAnalyzer,
				fieldDataFrequencyFilter: tt.fields.fielddataFrequencyFilter,
				indexOptions:             tt.fields.indexOptions,
				indexPrefixes:            tt.fields.indexPrefixes,
				similarity:               tt.fields.similarity,
				termVector:               tt.fields.termVector,
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
