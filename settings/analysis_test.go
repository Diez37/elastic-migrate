package settings

import (
	"reflect"
	"testing"
)

func TestAnalysis_AddAnalyzer(t *testing.T) {
	type fields struct {
		analyzer   []Analyzer
		normalizer []Normalizer
		filter     []Filter
		charFilter []CharFilter
		tokenizer  []Tokenizer
	}
	type args struct {
		analyzer []Analyzer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Analysis
	}{
		{
			name: "set_one",
			args: args{analyzer: []Analyzer{&AnalyzerStandard{}}},
			want: &Analysis{analyzer: []Analyzer{&AnalyzerStandard{}}},
		},
		{
			name: "set_two",
			args: args{analyzer: []Analyzer{&AnalyzerStandard{}, &AnalyzerStop{}}},
			want: &Analysis{analyzer: []Analyzer{&AnalyzerStandard{}, &AnalyzerStop{}}},
		},
		{
			name: "add",
			fields: fields{analyzer: []Analyzer{&AnalyzerStandard{}}},
			args: args{analyzer: []Analyzer{&AnalyzerStop{}}},
			want: &Analysis{analyzer: []Analyzer{&AnalyzerStandard{}, &AnalyzerStop{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &Analysis{
				analyzer:   tt.fields.analyzer,
				normalizer: tt.fields.normalizer,
				filter:     tt.fields.filter,
				charFilter: tt.fields.charFilter,
				tokenizer:  tt.fields.tokenizer,
			}
			if got := analysis.AddAnalyzer(tt.args.analyzer...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalysis_AddCharFilter(t *testing.T) {
	type fields struct {
		analyzer   []Analyzer
		normalizer []Normalizer
		filter     []Filter
		charFilter []CharFilter
		tokenizer  []Tokenizer
	}
	type args struct {
		charFilter []CharFilter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Analysis
	}{
		{
			name: "set_one",
			args: args{charFilter: []CharFilter{&CharFilterHtmlStrip{}}},
			want: &Analysis{charFilter: []CharFilter{&CharFilterHtmlStrip{}}},
		},
		{
			name: "set_two",
			args: args{charFilter: []CharFilter{&CharFilterHtmlStrip{}, &CharFilterMapping{}}},
			want: &Analysis{charFilter: []CharFilter{&CharFilterHtmlStrip{}, &CharFilterMapping{}}},
		},
		{
			name: "add",
			fields: fields{charFilter: []CharFilter{&CharFilterHtmlStrip{}}},
			args: args{charFilter: []CharFilter{&CharFilterMapping{}}},
			want: &Analysis{charFilter: []CharFilter{&CharFilterHtmlStrip{}, &CharFilterMapping{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &Analysis{
				analyzer:   tt.fields.analyzer,
				normalizer: tt.fields.normalizer,
				filter:     tt.fields.filter,
				charFilter: tt.fields.charFilter,
				tokenizer:  tt.fields.tokenizer,
			}
			if got := analysis.AddCharFilter(tt.args.charFilter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCharFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalysis_AddFilter(t *testing.T) {
	type fields struct {
		analyzer   []Analyzer
		normalizer []Normalizer
		filter     []Filter
		charFilter []CharFilter
		tokenizer  []Tokenizer
	}
	type args struct {
		filter []Filter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Analysis
	}{
		{
			name: "set_one",
			args: args{filter: []Filter{&FilterAsciiFolding{}}},
			want: &Analysis{filter: []Filter{&FilterAsciiFolding{}}},
		},
		{
			name: "set_two",
			args: args{filter: []Filter{&FilterAsciiFolding{}, &FilterCjkBigRam{}}},
			want: &Analysis{filter: []Filter{&FilterAsciiFolding{}, &FilterCjkBigRam{}}},
		},
		{
			name: "add",
			fields: fields{filter: []Filter{&FilterAsciiFolding{}}},
			args: args{filter: []Filter{&FilterCjkBigRam{}}},
			want: &Analysis{filter: []Filter{&FilterAsciiFolding{}, &FilterCjkBigRam{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &Analysis{
				analyzer:   tt.fields.analyzer,
				normalizer: tt.fields.normalizer,
				filter:     tt.fields.filter,
				charFilter: tt.fields.charFilter,
				tokenizer:  tt.fields.tokenizer,
			}
			if got := analysis.AddFilter(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalysis_AddNormalizer(t *testing.T) {
	type fields struct {
		analyzer   []Analyzer
		normalizer []Normalizer
		filter     []Filter
		charFilter []CharFilter
		tokenizer  []Tokenizer
	}
	type args struct {
		normalizer []Normalizer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Analysis
	}{
		{
			name: "set_one",
			args: args{normalizer: []Normalizer{&NormalizerCustom{}}},
			want: &Analysis{normalizer: []Normalizer{&NormalizerCustom{}}},
		},
		{
			name: "set_two",
			args: args{normalizer: []Normalizer{&NormalizerCustom{name: "test1"}, &NormalizerCustom{name: "test2"}}},
			want: &Analysis{normalizer: []Normalizer{&NormalizerCustom{name: "test1"}, &NormalizerCustom{name: "test2"}}},
		},
		{
			name: "add",
			fields: fields{normalizer: []Normalizer{&NormalizerCustom{name: "test1"}}},
			args: args{normalizer: []Normalizer{&NormalizerCustom{name: "test2"}}},
			want: &Analysis{normalizer: []Normalizer{&NormalizerCustom{name: "test1"}, &NormalizerCustom{name: "test2"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &Analysis{
				analyzer:   tt.fields.analyzer,
				normalizer: tt.fields.normalizer,
				filter:     tt.fields.filter,
				charFilter: tt.fields.charFilter,
				tokenizer:  tt.fields.tokenizer,
			}
			if got := analysis.AddNormalizer(tt.args.normalizer...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNormalizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalysis_AddTokenizer(t *testing.T) {
	type fields struct {
		analyzer   []Analyzer
		normalizer []Normalizer
		filter     []Filter
		charFilter []CharFilter
		tokenizer  []Tokenizer
	}
	type args struct {
		tokenizer []Tokenizer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Analysis
	}{
		{
			name: "set_one",
			args: args{tokenizer: []Tokenizer{&TokenizerCharGroup{}}},
			want: &Analysis{tokenizer: []Tokenizer{&TokenizerCharGroup{}}},
		},
		{
			name: "set_two",
			args: args{tokenizer: []Tokenizer{&TokenizerCharGroup{}, &TokenizerClassic{}}},
			want: &Analysis{tokenizer: []Tokenizer{&TokenizerCharGroup{}, &TokenizerClassic{}}},
		},
		{
			name: "add",
			fields: fields{tokenizer: []Tokenizer{&TokenizerCharGroup{}}},
			args: args{tokenizer: []Tokenizer{&TokenizerClassic{}}},
			want: &Analysis{tokenizer: []Tokenizer{&TokenizerCharGroup{}, &TokenizerClassic{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &Analysis{
				analyzer:   tt.fields.analyzer,
				normalizer: tt.fields.normalizer,
				filter:     tt.fields.filter,
				charFilter: tt.fields.charFilter,
				tokenizer:  tt.fields.tokenizer,
			}
			if got := analysis.AddTokenizer(tt.args.tokenizer...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTokenizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalysis_Source(t *testing.T) {
	type fields struct {
		analyzer   []Analyzer
		normalizer []Normalizer
		filter     []Filter
		charFilter []CharFilter
		tokenizer  []Tokenizer
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "analyzer",
			fields: fields{analyzer: []Analyzer{&AnalyzerStandard{name: "test"}}},
			want: map[string]interface{}{
				"analyzer": map[AnalyzerName]interface{}{
					"test": map[string]interface{}{
						"type": AnalyzerTypeStandard,
					},
				},
			},
		},
		{
			name: "normalizer",
			fields: fields{normalizer: []Normalizer{&NormalizerCustom{name: "test"}}},
			want: map[string]interface{}{
				"normalizer": map[NormalizerName]interface{}{
					"test": map[string]interface{}{
						"type": NormalizerTypeCustom,
					},
				},
			},
		},
		{
			name: "filter",
			fields: fields{filter: []Filter{&FilterUnique{name: "test"}}},
			want: map[string]interface{}{
				"filter": map[FilterName]interface{}{
					"test": map[string]interface{}{
						"type": FilterTypeUnique,
					},
				},
			},
		},
		{
			name: "char_filter",
			fields: fields{charFilter: []CharFilter{&CharFilterHtmlStrip{name: "test"}}},
			want: map[string]interface{}{
				"char_filter": map[CharFilterName]interface{}{
					"test": map[string]interface{}{
						"type": CharFilterTypeHtmlStrip,
					},
				},
			},
		},
		{
			name: "tokenizer",
			fields: fields{tokenizer: []Tokenizer{&TokenizerClassic{name: "test"}}},
			want: map[string]interface{}{
				"tokenizer": map[TokenizerName]interface{}{
					"test": map[string]interface{}{
						"type": TokenizerTypeClassic,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &Analysis{
				analyzer:   tt.fields.analyzer,
				normalizer: tt.fields.normalizer,
				filter:     tt.fields.filter,
				charFilter: tt.fields.charFilter,
				tokenizer:  tt.fields.tokenizer,
			}
			got, err := analysis.Source()
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

func TestNewAnalysis(t *testing.T) {
	tests := []struct {
		name string
		want *Analysis
	}{
		{
			want: &Analysis{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalysis(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalysis() = %v, want %v", got, tt.want)
			}
		})
	}
}
