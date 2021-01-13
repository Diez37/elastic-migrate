package settings

import (
	"github.com/olivere/elastic/v7"
	"reflect"
	"testing"
)

func TestFilterAsciiFolding_Name(t *testing.T) {
	type fields struct {
		name             FilterName
		preserveOriginal *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterAsciiFolding{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterAsciiFolding_SetPreserveOriginal(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             FilterName
		preserveOriginal *bool
	}
	type args struct {
		preserveOriginal bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterAsciiFolding
	}{
		{
			name: "false",
			args: args{preserveOriginal: false},
			want: &FilterAsciiFolding{preserveOriginal: &boolFalse},
		},
		{
			name: "true",
			args: args{preserveOriginal: true},
			want: &FilterAsciiFolding{preserveOriginal: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{preserveOriginal: &boolFalse},
			args:   args{preserveOriginal: true},
			want:   &FilterAsciiFolding{preserveOriginal: &boolTrue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterAsciiFolding{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.SetPreserveOriginal(tt.args.preserveOriginal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPreserveOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterAsciiFolding_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		name             FilterName
		preserveOriginal *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeAsciiFolding,
			},
		},
		{
			name:   "preserve_original",
			fields: fields{preserveOriginal: &boolTrue},
			want: map[string]interface{}{
				"type":              FilterTypeAsciiFolding,
				"preserve_original": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterAsciiFolding{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			got, err := filter.Source()
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

func TestFilterAsciiFolding_Type(t *testing.T) {
	type fields struct {
		name             FilterName
		preserveOriginal *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeAsciiFolding,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterAsciiFolding{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCjkBigRam_AddIgnoredScripts(t *testing.T) {
	type fields struct {
		name           FilterName
		outputUnigrams *bool
		ignoredScripts []FilterBigramIgnoredScripts
	}
	type args struct {
		ignoredScripts []FilterBigramIgnoredScripts
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCjkBigRam
	}{
		{
			name: "han",
			args: args{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsHan}},
			want: &FilterCjkBigRam{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsHan}},
		},
		{
			name: "hangul",
			args: args{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsHangul}},
			want: &FilterCjkBigRam{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsHangul}},
		},
		{
			name: "hiragana",
			args: args{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsHiragana}},
			want: &FilterCjkBigRam{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsHiragana}},
		},
		{
			name: "katakana",
			args: args{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsKatakana}},
			want: &FilterCjkBigRam{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsKatakana}},
		},
		{
			name: "many",
			args: args{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsKatakana, BigramIgnoredScriptsHan}},
			want: &FilterCjkBigRam{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsKatakana, BigramIgnoredScriptsHan}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCjkBigRam{
				name:           tt.fields.name,
				outputUnigrams: tt.fields.outputUnigrams,
				ignoredScripts: tt.fields.ignoredScripts,
			}
			if got := filter.AddIgnoredScripts(tt.args.ignoredScripts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddIgnoredScripts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCjkBigRam_Name(t *testing.T) {
	type fields struct {
		name           FilterName
		outputUnigrams *bool
		ignoredScripts []FilterBigramIgnoredScripts
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCjkBigRam{
				name:           tt.fields.name,
				outputUnigrams: tt.fields.outputUnigrams,
				ignoredScripts: tt.fields.ignoredScripts,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCjkBigRam_SetOutputUnigrams(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name           FilterName
		outputUnigrams *bool
		ignoredScripts []FilterBigramIgnoredScripts
	}
	type args struct {
		outputUnigrams bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCjkBigRam
	}{
		{
			name: "false",
			args: args{outputUnigrams: false},
			want: &FilterCjkBigRam{outputUnigrams: &boolFalse},
		},
		{
			name: "true",
			args: args{outputUnigrams: true},
			want: &FilterCjkBigRam{outputUnigrams: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{outputUnigrams: &boolFalse},
			args:   args{outputUnigrams: true},
			want:   &FilterCjkBigRam{outputUnigrams: &boolTrue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCjkBigRam{
				name:           tt.fields.name,
				outputUnigrams: tt.fields.outputUnigrams,
				ignoredScripts: tt.fields.ignoredScripts,
			}
			if got := filter.SetOutputUnigrams(tt.args.outputUnigrams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOutputUnigrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCjkBigRam_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		name           FilterName
		outputUnigrams *bool
		ignoredScripts []FilterBigramIgnoredScripts
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeCjkBigram,
			},
		},
		{
			name:   "output_unigrams",
			fields: fields{outputUnigrams: &boolTrue},
			want: map[string]interface{}{
				"type":            FilterTypeCjkBigram,
				"output_unigrams": true,
			},
		},
		{
			name:   "ignored_scripts",
			fields: fields{ignoredScripts: []FilterBigramIgnoredScripts{BigramIgnoredScriptsKatakana, BigramIgnoredScriptsHan}},
			want: map[string]interface{}{
				"type":            FilterTypeCjkBigram,
				"ignored_scripts": []FilterBigramIgnoredScripts{BigramIgnoredScriptsKatakana, BigramIgnoredScriptsHan},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCjkBigRam{
				name:           tt.fields.name,
				outputUnigrams: tt.fields.outputUnigrams,
				ignoredScripts: tt.fields.ignoredScripts,
			}
			got, err := filter.Source()
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

func TestFilterCjkBigRam_Type(t *testing.T) {
	type fields struct {
		name           FilterName
		outputUnigrams *bool
		ignoredScripts []FilterBigramIgnoredScripts
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeCjkBigram,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCjkBigRam{
				name:           tt.fields.name,
				outputUnigrams: tt.fields.outputUnigrams,
				ignoredScripts: tt.fields.ignoredScripts,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCommonGrams_AddCommonWords(t *testing.T) {
	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	type args struct {
		commonWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCommonGrams
	}{
		{
			name: "set",
			args: args{commonWords: []string{"test"}},
			want: &FilterCommonGrams{commonWords: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{commonWords: []string{"test"}},
			args:   args{commonWords: []string{"test1"}},
			want:   &FilterCommonGrams{commonWords: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			if got := filter.AddCommonWords(tt.args.commonWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCommonWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCommonGrams_Name(t *testing.T) {
	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCommonGrams_SetCommonWordsPath(t *testing.T) {
	commonWordsPathSet := "/tmp/commonWords.txt"
	commonWordsPathChange := "/tmp/commonWords.v1.txt"

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	type args struct {
		commonWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCommonGrams
	}{
		{
			name: "set",
			args: args{commonWordsPath: "/tmp/commonWords.txt"},
			want: &FilterCommonGrams{commonWordsPath: &commonWordsPathSet},
		},
		{
			name:   "change",
			fields: fields{commonWordsPath: &commonWordsPathSet},
			args:   args{commonWordsPath: "/tmp/commonWords.v1.txt"},
			want:   &FilterCommonGrams{commonWordsPath: &commonWordsPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			if got := filter.SetCommonWordsPath(tt.args.commonWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCommonWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCommonGrams_SetIgnoreCase(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	type args struct {
		ignoreCase bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCommonGrams
	}{
		{
			name: "false",
			args: args{ignoreCase: false},
			want: &FilterCommonGrams{ignoreCase: &boolFalse},
		},
		{
			name: "true",
			args: args{ignoreCase: true},
			want: &FilterCommonGrams{ignoreCase: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{ignoreCase: &boolFalse},
			args:   args{ignoreCase: true},
			want:   &FilterCommonGrams{ignoreCase: &boolTrue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			if got := filter.SetIgnoreCase(tt.args.ignoreCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCommonGrams_SetQueryMode(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	type args struct {
		queryMode bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCommonGrams
	}{
		{
			name: "false",
			args: args{queryMode: false},
			want: &FilterCommonGrams{queryMode: &boolFalse},
		},
		{
			name: "true",
			args: args{queryMode: true},
			want: &FilterCommonGrams{queryMode: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{queryMode: &boolFalse},
			args:   args{queryMode: true},
			want:   &FilterCommonGrams{queryMode: &boolTrue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			if got := filter.SetQueryMode(tt.args.queryMode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetQueryMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCommonGrams_Source(t *testing.T) {
	commonWordsPath := "/tmp/commonWords.txt"
	boolTrue := true

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeCommonGrams,
			},
		},
		{
			name:   "ignore_case",
			fields: fields{ignoreCase: &boolTrue},
			want: map[string]interface{}{
				"type":        FilterTypeCommonGrams,
				"ignore_case": true,
			},
		},
		{
			name:   "query_mode",
			fields: fields{queryMode: &boolTrue},
			want: map[string]interface{}{
				"type":       FilterTypeCommonGrams,
				"query_mode": true,
			},
		},
		{
			name:   "common_words_path",
			fields: fields{commonWordsPath: &commonWordsPath},
			want: map[string]interface{}{
				"type":              FilterTypeCommonGrams,
				"common_words_path": "/tmp/commonWords.txt",
			},
		},
		{
			name:   "common_words",
			fields: fields{commonWords: []string{"test", "test1"}},
			want: map[string]interface{}{
				"type":         FilterTypeCommonGrams,
				"common_words": []string{"test", "test1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			got, err := filter.Source()
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

func TestFilterCommonGrams_Type(t *testing.T) {
	type fields struct {
		name            FilterName
		ignoreCase      *bool
		queryMode       *bool
		commonWordsPath *string
		commonWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeCommonGrams,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCommonGrams{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				queryMode:       tt.fields.queryMode,
				commonWordsPath: tt.fields.commonWordsPath,
				commonWords:     tt.fields.commonWords,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCondition_AddFilter(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
		filter []FilterName
	}
	type args struct {
		filters []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCondition
	}{
		{
			name: "filter_name_arabic",
			args: args{filters: []FilterName{FilterNameArabic}},
			want: &FilterCondition{filter: []FilterName{FilterNameArabic}},
		},
		{
			name: "filter_name_asciifolding",
			args: args{filters: []FilterName{FilterNameAsciifolding}},
			want: &FilterCondition{filter: []FilterName{FilterNameAsciifolding}},
		},
		{
			name: "filter_name_bengali",
			args: args{filters: []FilterName{FilterNameBengali}},
			want: &FilterCondition{filter: []FilterName{FilterNameBengali}},
		},
		{
			name: "filter_name_cjk_width",
			args: args{filters: []FilterName{FilterNameCjkWidth}},
			want: &FilterCondition{filter: []FilterName{FilterNameCjkWidth}},
		},
		{
			name: "filter_name_decimal_digit",
			args: args{filters: []FilterName{FilterNameDecimalDigit}},
			want: &FilterCondition{filter: []FilterName{FilterNameDecimalDigit}},
		},
		{
			name: "filter_name_delimited_payload",
			args: args{filters: []FilterName{FilterNameDelimitedPayload}},
			want: &FilterCondition{filter: []FilterName{FilterNameDelimitedPayload}},
		},
		{
			name: "filter_name_elision",
			args: args{filters: []FilterName{FilterNameElision}},
			want: &FilterCondition{filter: []FilterName{FilterNameElision}},
		},
		{
			name: "filter_name_german",
			args: args{filters: []FilterName{FilterNameGerman}},
			want: &FilterCondition{filter: []FilterName{FilterNameGerman}},
		},
		{
			name: "filter_name_hindi",
			args: args{filters: []FilterName{FilterNameHindi}},
			want: &FilterCondition{filter: []FilterName{FilterNameHindi}},
		},
		{
			name: "filter_name_indic",
			args: args{filters: []FilterName{FilterNameIndic}},
			want: &FilterCondition{filter: []FilterName{FilterNameIndic}},
		},
		{
			name: "filter_name_lowercase",
			args: args{filters: []FilterName{FilterNameLowercase}},
			want: &FilterCondition{filter: []FilterName{FilterNameLowercase}},
		},
		{
			name: "filter_name_persian",
			args: args{filters: []FilterName{FilterNamePersian}},
			want: &FilterCondition{filter: []FilterName{FilterNamePersian}},
		},
		{
			name: "filter_name_scandinavian_folding",
			args: args{filters: []FilterName{FilterNameScandinavianFolding}},
			want: &FilterCondition{filter: []FilterName{FilterNameScandinavianFolding}},
		},
		{
			name: "filter_name_serbian",
			args: args{filters: []FilterName{FilterNameSerbian}},
			want: &FilterCondition{filter: []FilterName{FilterNameSerbian}},
		},
		{
			name: "filter_name_sorani",
			args: args{filters: []FilterName{FilterNameSorani}},
			want: &FilterCondition{filter: []FilterName{FilterNameSorani}},
		},
		{
			name: "filter_name_uppercase",
			args: args{filters: []FilterName{FilterNameUppercase}},
			want: &FilterCondition{filter: []FilterName{FilterNameUppercase}},
		},
		{
			name: "filter_name_apostrophe",
			args: args{filters: []FilterName{FilterNameApostrophe}},
			want: &FilterCondition{filter: []FilterName{FilterNameApostrophe}},
		},
		{
			name: "filter_name_ascii_folding",
			args: args{filters: []FilterName{FilterNameAsciiFolding}},
			want: &FilterCondition{filter: []FilterName{FilterNameAsciiFolding}},
		},
		{
			name: "filter_name_cjk_bigram",
			args: args{filters: []FilterName{FilterNameCjkBigram}},
			want: &FilterCondition{filter: []FilterName{FilterNameCjkBigram}},
		},
		{
			name: "filter_name_classic",
			args: args{filters: []FilterName{FilterNameClassic}},
			want: &FilterCondition{filter: []FilterName{FilterNameClassic}},
		},
		{
			name: "filter_name_edge_ngram",
			args: args{filters: []FilterName{FilterNameEdgeNgram}},
			want: &FilterCondition{filter: []FilterName{FilterNameEdgeNgram}},
		},
		{
			name: "filter_name_fingerprint",
			args: args{filters: []FilterName{FilterNameFingerprint}},
			want: &FilterCondition{filter: []FilterName{FilterNameFingerprint}},
		},
		{
			name: "filter_name_flatten_graph",
			args: args{filters: []FilterName{FilterNameFlattenGraph}},
			want: &FilterCondition{filter: []FilterName{FilterNameFlattenGraph}},
		},
		{
			name: "filter_name_keyword_repeat",
			args: args{filters: []FilterName{FilterNameKeywordRepeat}},
			want: &FilterCondition{filter: []FilterName{FilterNameKeywordRepeat}},
		},
		{
			name: "filter_name_kstem",
			args: args{filters: []FilterName{FilterNameKstem}},
			want: &FilterCondition{filter: []FilterName{FilterNameKstem}},
		},
		{
			name: "filter_name_length",
			args: args{filters: []FilterName{FilterNameLength}},
			want: &FilterCondition{filter: []FilterName{FilterNameLength}},
		},
		{
			name: "filter_name_limit",
			args: args{filters: []FilterName{FilterNameLimit}},
			want: &FilterCondition{filter: []FilterName{FilterNameLimit}},
		},
		{
			name: "filter_name_ngram",
			args: args{filters: []FilterName{FilterNameNgram}},
			want: &FilterCondition{filter: []FilterName{FilterNameNgram}},
		},
		{
			name: "filter_name_porter_stem",
			args: args{filters: []FilterName{FilterNamePorterStem}},
			want: &FilterCondition{filter: []FilterName{FilterNamePorterStem}},
		},
		{
			name: "filter_name_remove_duplicates",
			args: args{filters: []FilterName{FilterNameRemoveDuplicates}},
			want: &FilterCondition{filter: []FilterName{FilterNameRemoveDuplicates}},
		},
		{
			name: "filter_name_reverse",
			args: args{filters: []FilterName{FilterNameReverse}},
			want: &FilterCondition{filter: []FilterName{FilterNameReverse}},
		},
		{
			name: "filter_name_shingle",
			args: args{filters: []FilterName{FilterNameShingle}},
			want: &FilterCondition{filter: []FilterName{FilterNameShingle}},
		},
		{
			name: "filter_name_stemmer",
			args: args{filters: []FilterName{FilterNameStemmer}},
			want: &FilterCondition{filter: []FilterName{FilterNameStemmer}},
		},
		{
			name: "filter_name_stop",
			args: args{filters: []FilterName{FilterNameStop}},
			want: &FilterCondition{filter: []FilterName{FilterNameStop}},
		},
		{
			name: "filter_name_trim",
			args: args{filters: []FilterName{FilterNameTrim}},
			want: &FilterCondition{filter: []FilterName{FilterNameTrim}},
		},
		{
			name: "filter_name_truncate",
			args: args{filters: []FilterName{FilterNameTruncate}},
			want: &FilterCondition{filter: []FilterName{FilterNameTruncate}},
		},
		{
			name: "filter_name_unique",
			args: args{filters: []FilterName{FilterNameUnique}},
			want: &FilterCondition{filter: []FilterName{FilterNameUnique}},
		},
		{
			name: "filter_name_word_delimiter",
			args: args{filters: []FilterName{FilterNameWordDelimiter}},
			want: &FilterCondition{filter: []FilterName{FilterNameWordDelimiter}},
		},
		{
			name: "filter_name_word_delimiter_graph",
			args: args{filters: []FilterName{FilterNameWordDelimiterGraph}},
			want: &FilterCondition{filter: []FilterName{FilterNameWordDelimiterGraph}},
		},
		{
			name: "many",
			args: args{filters: []FilterName{FilterNameWordDelimiterGraph, FilterNameWordDelimiter}},
			want: &FilterCondition{filter: []FilterName{FilterNameWordDelimiterGraph, FilterNameWordDelimiter}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCondition{
				name:   tt.fields.name,
				script: tt.fields.script,
				filter: tt.fields.filter,
			}
			if got := filter.AddFilter(tt.args.filters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCondition_Name(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
		filter []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCondition{
				name:   tt.fields.name,
				script: tt.fields.script,
				filter: tt.fields.filter,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCondition_SetScript(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
		filter []FilterName
	}
	type args struct {
		script *elastic.Script
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterCondition
	}{
		{
			name: "set",
			args: args{script: elastic.NewScript("test")},
			want: &FilterCondition{script: elastic.NewScript("test")},
		},
		{
			name:   "change",
			fields: fields{script: elastic.NewScript("test1")},
			args:   args{script: elastic.NewScript("test")},
			want:   &FilterCondition{script: elastic.NewScript("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCondition{
				name:   tt.fields.name,
				script: tt.fields.script,
				filter: tt.fields.filter,
			}
			if got := filter.SetScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCondition_Source(t *testing.T) {
	scriptSrc, _ := elastic.NewScript("test").Source()

	type fields struct {
		name   FilterName
		script *elastic.Script
		filter []FilterName
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeCondition,
			},
		},
		{
			name:   "script",
			fields: fields{script: elastic.NewScript("test")},
			want: map[string]interface{}{
				"type":   FilterTypeCondition,
				"script": scriptSrc,
			},
		},
		{
			name:   "filter",
			fields: fields{filter: []FilterName{FilterNameArabic, FilterNameClassic}},
			want: map[string]interface{}{
				"type":   FilterTypeCondition,
				"filter": []FilterName{FilterNameArabic, FilterNameClassic},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCondition{
				name:   tt.fields.name,
				script: tt.fields.script,
				filter: tt.fields.filter,
			}
			got, err := filter.Source()
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

func TestFilterCondition_Type(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
		filter []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeCondition,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterCondition{
				name:   tt.fields.name,
				script: tt.fields.script,
				filter: tt.fields.filter,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDelimitedPayload_Name(t *testing.T) {
	type fields struct {
		name      FilterName
		delimiter *string
		encoding  *FilterDelimitedPayloadEncoding
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDelimitedPayload{
				name:      tt.fields.name,
				delimiter: tt.fields.delimiter,
				encoding:  tt.fields.encoding,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDelimitedPayload_SetDelimiter(t *testing.T) {
	delimiterSet := "|"
	delimiterChange := "!"

	type fields struct {
		name      FilterName
		delimiter *string
		encoding  *FilterDelimitedPayloadEncoding
	}
	type args struct {
		delimiter string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDelimitedPayload
	}{
		{
			args: args{delimiter: "|"},
			want: &FilterDelimitedPayload{delimiter: &delimiterSet},
		},
		{
			args:   args{delimiter: "!"},
			fields: fields{delimiter: &delimiterSet},
			want:   &FilterDelimitedPayload{delimiter: &delimiterChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDelimitedPayload{
				name:      tt.fields.name,
				delimiter: tt.fields.delimiter,
				encoding:  tt.fields.encoding,
			}
			if got := filter.SetDelimiter(tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDelimitedPayload_SetEncoding(t *testing.T) {
	delimitedPayloadEncodingFloat := DelimitedPayloadEncodingFloat
	delimitedPayloadEncodingInteger := DelimitedPayloadEncodingInteger
	delimitedPayloadEncodingIdentity := DelimitedPayloadEncodingIdentity

	type fields struct {
		name      FilterName
		delimiter *string
		encoding  *FilterDelimitedPayloadEncoding
	}
	type args struct {
		encoding FilterDelimitedPayloadEncoding
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDelimitedPayload
	}{
		{
			name: "delimited_payload_encoding_float",
			args: args{encoding: DelimitedPayloadEncodingFloat},
			want: &FilterDelimitedPayload{encoding: &delimitedPayloadEncodingFloat},
		},
		{
			name: "delimited_payload_encoding_integer",
			args: args{encoding: DelimitedPayloadEncodingInteger},
			want: &FilterDelimitedPayload{encoding: &delimitedPayloadEncodingInteger},
		},
		{
			name: "delimited_payload_encoding_identity",
			args: args{encoding: DelimitedPayloadEncodingIdentity},
			want: &FilterDelimitedPayload{encoding: &delimitedPayloadEncodingIdentity},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDelimitedPayload{
				name:      tt.fields.name,
				delimiter: tt.fields.delimiter,
				encoding:  tt.fields.encoding,
			}
			if got := filter.SetEncoding(tt.args.encoding); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDelimitedPayload_Source(t *testing.T) {
	delimiterSet := "|"
	delimitedPayloadEncodingFloat := DelimitedPayloadEncodingFloat

	type fields struct {
		name      FilterName
		delimiter *string
		encoding  *FilterDelimitedPayloadEncoding
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeDelimitedPayload,
			},
		},
		{
			name:   "delimiter",
			fields: fields{delimiter: &delimiterSet},
			want: map[string]interface{}{
				"type":      FilterTypeDelimitedPayload,
				"delimiter": "|",
			},
		},
		{
			name:   "encoding",
			fields: fields{encoding: &delimitedPayloadEncodingFloat},
			want: map[string]interface{}{
				"type":     FilterTypeDelimitedPayload,
				"encoding": DelimitedPayloadEncodingFloat,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDelimitedPayload{
				name:      tt.fields.name,
				delimiter: tt.fields.delimiter,
				encoding:  tt.fields.encoding,
			}
			got, err := filter.Source()
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

func TestFilterDelimitedPayload_Type(t *testing.T) {
	type fields struct {
		name      FilterName
		delimiter *string
		encoding  *FilterDelimitedPayloadEncoding
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeDelimitedPayload,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDelimitedPayload{
				name:      tt.fields.name,
				delimiter: tt.fields.delimiter,
				encoding:  tt.fields.encoding,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_AddWordList(t *testing.T) {
	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	type args struct {
		wordList []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDictionaryDecompounder
	}{
		{
			name: "set",
			args: args{wordList: []string{"test"}},
			want: &FilterDictionaryDecompounder{wordList: []string{"test"}},
		},
		{
			name:   "add",
			fields: fields{wordList: []string{"test"}},
			args:   args{wordList: []string{"test1"}},
			want:   &FilterDictionaryDecompounder{wordList: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.AddWordList(tt.args.wordList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWordList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_Name(t *testing.T) {
	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_SetMaxSubwordSize(t *testing.T) {
	maxSubwordSizeSet := uint32(1)
	maxSubwordSizeChange := uint32(2)

	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	type args struct {
		maxSubwordSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDictionaryDecompounder
	}{
		{
			name: "set",
			args: args{maxSubwordSize: 1},
			want: &FilterDictionaryDecompounder{maxSubwordSize: &maxSubwordSizeSet},
		},
		{
			name:   "change",
			fields: fields{maxSubwordSize: &maxSubwordSizeSet},
			args:   args{maxSubwordSize: 2},
			want:   &FilterDictionaryDecompounder{maxSubwordSize: &maxSubwordSizeChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.SetMaxSubwordSize(tt.args.maxSubwordSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxSubwordSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_SetMinSubwordSize(t *testing.T) {
	minSubwordSizeSet := uint32(1)
	minSubwordSizeChange := uint32(2)

	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	type args struct {
		minSubwordSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDictionaryDecompounder
	}{
		{
			name: "set",
			args: args{minSubwordSize: 1},
			want: &FilterDictionaryDecompounder{minSubwordSize: &minSubwordSizeSet},
		},
		{
			name:   "change",
			fields: fields{minSubwordSize: &minSubwordSizeSet},
			args:   args{minSubwordSize: 2},
			want:   &FilterDictionaryDecompounder{minSubwordSize: &minSubwordSizeChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.SetMinSubwordSize(tt.args.minSubwordSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinSubwordSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_SetMinWordSize(t *testing.T) {
	minWordSizeSet := uint32(1)
	minWordSizeChange := uint32(2)

	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	type args struct {
		minWordSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDictionaryDecompounder
	}{
		{
			name: "set",
			args: args{minWordSize: 1},
			want: &FilterDictionaryDecompounder{minWordSize: &minWordSizeSet},
		},
		{
			name:   "change",
			fields: fields{minWordSize: &minWordSizeSet},
			args:   args{minWordSize: 2},
			want:   &FilterDictionaryDecompounder{minWordSize: &minWordSizeChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.SetMinWordSize(tt.args.minWordSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinWordSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_SetOnlyLongestMatch(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	type args struct {
		onlyLongestMatch bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDictionaryDecompounder
	}{
		{
			name: "set",
			args: args{onlyLongestMatch: true},
			want: &FilterDictionaryDecompounder{onlyLongestMatch: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{onlyLongestMatch: &boolTrue},
			args:   args{onlyLongestMatch: false},
			want:   &FilterDictionaryDecompounder{onlyLongestMatch: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.SetOnlyLongestMatch(tt.args.onlyLongestMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlyLongestMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_SetWordListPath(t *testing.T) {
	wordListPathSet := "/temp/wordList.txt"
	wordListPathChange := "/temp/wordList.v1.txt"

	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	type args struct {
		wordListPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterDictionaryDecompounder
	}{
		{
			name: "set",
			args: args{wordListPath: "/temp/wordList.txt"},
			want: &FilterDictionaryDecompounder{wordListPath: &wordListPathSet},
		},
		{
			name:   "change",
			fields: fields{wordListPath: &wordListPathSet},
			args:   args{wordListPath: "/temp/wordList.v1.txt"},
			want:   &FilterDictionaryDecompounder{wordListPath: &wordListPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.SetWordListPath(tt.args.wordListPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWordListPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDictionaryDecompounder_Source(t *testing.T) {
	wordListPath := "/temp/wordList.txt"
	maxSubwordSize := uint32(1)
	minSubwordSize := uint32(1)
	minWordSize := uint32(1)
	boolTrue := true

	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeDictionaryDecompounder,
			},
		},
		{
			name:   "word_list_path",
			fields: fields{wordListPath: &wordListPath},
			want: map[string]interface{}{
				"type":           FilterTypeDictionaryDecompounder,
				"word_list_path": "/temp/wordList.txt",
			},
		},
		{
			name:   "max_subword_size",
			fields: fields{maxSubwordSize: &maxSubwordSize},
			want: map[string]interface{}{
				"type":             FilterTypeDictionaryDecompounder,
				"max_subword_size": maxSubwordSize,
			},
		},
		{
			name:   "min_subword_size",
			fields: fields{minSubwordSize: &minSubwordSize},
			want: map[string]interface{}{
				"type":             FilterTypeDictionaryDecompounder,
				"min_subword_size": minSubwordSize,
			},
		},
		{
			name:   "min_word_size",
			fields: fields{minWordSize: &minWordSize},
			want: map[string]interface{}{
				"type":          FilterTypeDictionaryDecompounder,
				"min_word_size": minWordSize,
			},
		},
		{
			name:   "only_longest_match",
			fields: fields{onlyLongestMatch: &boolTrue},
			want: map[string]interface{}{
				"type":               FilterTypeDictionaryDecompounder,
				"only_longest_match": true,
			},
		},
		{
			name:   "word_list",
			fields: fields{wordList: []string{"test"}},
			want: map[string]interface{}{
				"type":      FilterTypeDictionaryDecompounder,
				"word_list": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			got, err := filter.Source()
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

func TestFilterDictionaryDecompounder_Type(t *testing.T) {
	type fields struct {
		name             FilterName
		wordListPath     *string
		maxSubwordSize   *uint32
		minSubwordSize   *uint32
		minWordSize      *uint32
		onlyLongestMatch *bool
		wordList         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeDictionaryDecompounder,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterDictionaryDecompounder{
				name:             tt.fields.name,
				wordListPath:     tt.fields.wordListPath,
				maxSubwordSize:   tt.fields.maxSubwordSize,
				minSubwordSize:   tt.fields.minSubwordSize,
				minWordSize:      tt.fields.minWordSize,
				onlyLongestMatch: tt.fields.onlyLongestMatch,
				wordList:         tt.fields.wordList,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEdgeNgram_Name(t *testing.T) {
	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEdgeNgram_SetMaxGram(t *testing.T) {
	maxGramSet := uint32(1)
	maxGramChange := uint32(2)

	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	type args struct {
		maxGram uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterEdgeNgram
	}{
		{
			name: "set",
			args: args{maxGram: 1},
			want: &FilterEdgeNgram{maxGram: &maxGramSet},
		},
		{
			name:   "change",
			fields: fields{maxGram: &maxGramSet},
			args:   args{maxGram: 2},
			want:   &FilterEdgeNgram{maxGram: &maxGramChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			if got := filter.SetMaxGram(tt.args.maxGram); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEdgeNgram_SetMinGram(t *testing.T) {
	minGramSet := uint32(1)
	minGramChange := uint32(2)

	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	type args struct {
		minGram uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterEdgeNgram
	}{
		{
			name: "set",
			args: args{minGram: 1},
			want: &FilterEdgeNgram{minGram: &minGramSet},
		},
		{
			name:   "change",
			fields: fields{minGram: &minGramSet},
			args:   args{minGram: 2},
			want:   &FilterEdgeNgram{minGram: &minGramChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			if got := filter.SetMinGram(tt.args.minGram); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEdgeNgram_SetPreserveOriginal(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	type args struct {
		preserveOriginal bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterEdgeNgram
	}{
		{
			name: "set",
			args: args{preserveOriginal: true},
			want: &FilterEdgeNgram{preserveOriginal: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{preserveOriginal: &boolTrue},
			args:   args{preserveOriginal: false},
			want:   &FilterEdgeNgram{preserveOriginal: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			if got := filter.SetPreserveOriginal(tt.args.preserveOriginal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPreserveOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEdgeNgram_SetSide(t *testing.T) {
	edgeNgramSideFront := EdgeNgramSideFront
	edgeNgramSideBack := EdgeNgramSideBack

	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	type args struct {
		side FilterEdgeNgramSide
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterEdgeNgram
	}{
		{
			name: "edge_ngram_side_front",
			args: args{side: EdgeNgramSideFront},
			want: &FilterEdgeNgram{side: &edgeNgramSideFront},
		},
		{
			name: "edge_ngram_side_back",
			args: args{side: EdgeNgramSideBack},
			want: &FilterEdgeNgram{side: &edgeNgramSideBack},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			if got := filter.SetSide(tt.args.side); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterEdgeNgram_Source(t *testing.T) {
	maxGram := uint32(1)
	minGram := uint32(1)
	boolTrue := true
	edgeNgramSideFront := EdgeNgramSideFront

	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeEdgeNgram,
			},
		},
		{
			name:   "max_gram",
			fields: fields{maxGram: &maxGram},
			want: map[string]interface{}{
				"type":     FilterTypeEdgeNgram,
				"max_gram": uint32(1),
			},
		},
		{
			name:   "min_gram",
			fields: fields{minGram: &minGram},
			want: map[string]interface{}{
				"type":     FilterTypeEdgeNgram,
				"min_gram": uint32(1),
			},
		},
		{
			name:   "preserve_original",
			fields: fields{preserveOriginal: &boolTrue},
			want: map[string]interface{}{
				"type":              FilterTypeEdgeNgram,
				"preserve_original": true,
			},
		},
		{
			name:   "side",
			fields: fields{side: &edgeNgramSideFront},
			want: map[string]interface{}{
				"type": FilterTypeEdgeNgram,
				"side": EdgeNgramSideFront,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			got, err := filter.Source()
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

func TestFilterEdgeNgram_Type(t *testing.T) {
	type fields struct {
		name             FilterName
		maxGram          *uint32
		minGram          *uint32
		preserveOriginal *bool
		side             *FilterEdgeNgramSide
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeEdgeNgram,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterEdgeNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
				side:             tt.fields.side,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterElision_AddArticles(t *testing.T) {
	type fields struct {
		name         FilterName
		articlesPath *string
		articlesCase *bool
		articles     []string
	}
	type args struct {
		articles []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterElision
	}{
		{
			name: "set",
			args: args{articles: []string{"test"}},
			want: &FilterElision{articles: []string{"test"}},
		},
		{
			name:   "add",
			fields: fields{articles: []string{"test"}},
			args:   args{articles: []string{"test1"}},
			want:   &FilterElision{articles: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterElision{
				name:         tt.fields.name,
				articlesPath: tt.fields.articlesPath,
				articlesCase: tt.fields.articlesCase,
				articles:     tt.fields.articles,
			}
			if got := filter.AddArticles(tt.args.articles...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterElision_Name(t *testing.T) {
	type fields struct {
		name         FilterName
		articlesPath *string
		articlesCase *bool
		articles     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterElision{
				name:         tt.fields.name,
				articlesPath: tt.fields.articlesPath,
				articlesCase: tt.fields.articlesCase,
				articles:     tt.fields.articles,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterElision_SetArticlesCase(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name         FilterName
		articlesPath *string
		articlesCase *bool
		articles     []string
	}
	type args struct {
		articlesCase bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterElision
	}{
		{
			name: "set",
			args: args{articlesCase: true},
			want: &FilterElision{articlesCase: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{articlesCase: &boolTrue},
			args:   args{articlesCase: false},
			want:   &FilterElision{articlesCase: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterElision{
				name:         tt.fields.name,
				articlesPath: tt.fields.articlesPath,
				articlesCase: tt.fields.articlesCase,
				articles:     tt.fields.articles,
			}
			if got := filter.SetArticlesCase(tt.args.articlesCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetArticlesCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterElision_SetArticlesPath(t *testing.T) {
	pathSet := "/temp/test.txt"
	pathChange := "/temp/test.v1.txt"

	type fields struct {
		name         FilterName
		articlesPath *string
		articlesCase *bool
		articles     []string
	}
	type args struct {
		articlesPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterElision
	}{
		{
			name: "set",
			args: args{articlesPath: "/temp/test.txt"},
			want: &FilterElision{articlesPath: &pathSet},
		},
		{
			name:   "change",
			fields: fields{articlesPath: &pathSet},
			args:   args{articlesPath: "/temp/test.v1.txt"},
			want:   &FilterElision{articlesPath: &pathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterElision{
				name:         tt.fields.name,
				articlesPath: tt.fields.articlesPath,
				articlesCase: tt.fields.articlesCase,
				articles:     tt.fields.articles,
			}
			if got := filter.SetArticlesPath(tt.args.articlesPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetArticlesPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterElision_Source(t *testing.T) {
	boolTrue := true
	path := "/temp/test.txt"

	type fields struct {
		name         FilterName
		articlesPath *string
		articlesCase *bool
		articles     []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeElision,
			},
		},
		{
			name:   "articles_path",
			fields: fields{articlesPath: &path},
			want: map[string]interface{}{
				"type":          FilterTypeElision,
				"articles_path": "/temp/test.txt",
			},
		},
		{
			name:   "articles_case",
			fields: fields{articlesCase: &boolTrue},
			want: map[string]interface{}{
				"type":          FilterTypeElision,
				"articles_case": true,
			},
		},
		{
			name:   "articles",
			fields: fields{articles: []string{"test"}},
			want: map[string]interface{}{
				"type":     FilterTypeElision,
				"articles": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterElision{
				name:         tt.fields.name,
				articlesPath: tt.fields.articlesPath,
				articlesCase: tt.fields.articlesCase,
				articles:     tt.fields.articles,
			}
			got, err := filter.Source()
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

func TestFilterElision_Type(t *testing.T) {
	type fields struct {
		name         FilterName
		articlesPath *string
		articlesCase *bool
		articles     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeElision,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterElision{
				name:         tt.fields.name,
				articlesPath: tt.fields.articlesPath,
				articlesCase: tt.fields.articlesCase,
				articles:     tt.fields.articles,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterFingerprint_Name(t *testing.T) {
	type fields struct {
		name          FilterName
		maxOutputSize *uint32
		separator     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterFingerprint_SetMaxOutputSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name          FilterName
		maxOutputSize *uint32
		separator     *string
	}
	type args struct {
		maxOutputSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterFingerprint
	}{
		{
			name: "set",
			args: args{maxOutputSize: 1},
			want: &FilterFingerprint{maxOutputSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{maxOutputSize: &numberSet},
			args:   args{maxOutputSize: 2},
			want:   &FilterFingerprint{maxOutputSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
			}
			if got := filter.SetMaxOutputSize(tt.args.maxOutputSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxOutputSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterFingerprint_SetSeparator(t *testing.T) {
	separatorSet := "|"
	separatorChange := "!"

	type fields struct {
		name          FilterName
		maxOutputSize *uint32
		separator     *string
	}
	type args struct {
		separator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterFingerprint
	}{
		{
			name: "set",
			args: args{separator: "|"},
			want: &FilterFingerprint{separator: &separatorSet},
		},
		{
			name:   "change",
			fields: fields{separator: &separatorSet},
			args:   args{separator: "!"},
			want:   &FilterFingerprint{separator: &separatorChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
			}
			if got := filter.SetSeparator(tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterFingerprint_Source(t *testing.T) {
	number := uint32(1)
	separator := "|"

	type fields struct {
		name          FilterName
		maxOutputSize *uint32
		separator     *string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeFingerprint,
			},
		},
		{
			name:   "max_output_size",
			fields: fields{maxOutputSize: &number},
			want: map[string]interface{}{
				"type":            FilterTypeFingerprint,
				"max_output_size": uint32(1),
			},
		},
		{
			name:   "separator",
			fields: fields{separator: &separator},
			want: map[string]interface{}{
				"type":      FilterTypeFingerprint,
				"separator": "|",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
			}
			got, err := filter.Source()
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

func TestFilterFingerprint_Type(t *testing.T) {
	type fields struct {
		name          FilterName
		maxOutputSize *uint32
		separator     *string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeFingerprint,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_AddWordList(t *testing.T) {
	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		wordList []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{wordList: []string{"test"}},
			want: &FilterHyphenationDecompounder{wordList: []string{"test"}},
		},
		{
			name:   "add",
			fields: fields{wordList: []string{"test"}},
			args:   args{wordList: []string{"test1"}},
			want:   &FilterHyphenationDecompounder{wordList: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.AddWordList(tt.args.wordList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWordList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_Name(t *testing.T) {
	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_SetHyphenationPatternsPath(t *testing.T) {
	pathSet := "/temp/test.txt"
	pathChange := "/temp/test.v1.txt"

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		hyphenationPatternsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{hyphenationPatternsPath: "/temp/test.txt"},
			want: &FilterHyphenationDecompounder{hyphenationPatternsPath: &pathSet},
		},
		{
			name:   "change",
			fields: fields{hyphenationPatternsPath: &pathSet},
			args:   args{hyphenationPatternsPath: "/temp/test.v1.txt"},
			want:   &FilterHyphenationDecompounder{hyphenationPatternsPath: &pathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.SetHyphenationPatternsPath(tt.args.hyphenationPatternsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHyphenationPatternsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_SetMaxSubwordSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		maxSubwordSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{maxSubwordSize: 1},
			want: &FilterHyphenationDecompounder{maxSubwordSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{maxSubwordSize: &numberSet},
			args:   args{maxSubwordSize: 2},
			want:   &FilterHyphenationDecompounder{maxSubwordSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.SetMaxSubwordSize(tt.args.maxSubwordSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxSubwordSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_SetMinSubwordSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		minSubwordSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{minSubwordSize: 1},
			want: &FilterHyphenationDecompounder{minSubwordSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{minSubwordSize: &numberSet},
			args:   args{minSubwordSize: 2},
			want:   &FilterHyphenationDecompounder{minSubwordSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.SetMinSubwordSize(tt.args.minSubwordSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinSubwordSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_SetMinWordSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		minWordSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{minWordSize: 1},
			want: &FilterHyphenationDecompounder{minWordSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{minWordSize: &numberSet},
			args:   args{minWordSize: 2},
			want:   &FilterHyphenationDecompounder{minWordSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.SetMinWordSize(tt.args.minWordSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinWordSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_SetOnlyLongestMatch(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		onlyLongestMatch bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{onlyLongestMatch: true},
			want: &FilterHyphenationDecompounder{onlyLongestMatch: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{onlyLongestMatch: &boolTrue},
			args:   args{onlyLongestMatch: false},
			want:   &FilterHyphenationDecompounder{onlyLongestMatch: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.SetOnlyLongestMatch(tt.args.onlyLongestMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlyLongestMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_SetWordListPath(t *testing.T) {
	pathSet := "/temp/test.txt"
	pathChange := "/temp/test.v1.txt"

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	type args struct {
		wordListPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterHyphenationDecompounder
	}{
		{
			name: "set",
			args: args{wordListPath: "/temp/test.txt"},
			want: &FilterHyphenationDecompounder{wordListPath: &pathSet},
		},
		{
			name:   "change",
			fields: fields{wordListPath: &pathSet},
			args:   args{wordListPath: "/temp/test.v1.txt"},
			want:   &FilterHyphenationDecompounder{wordListPath: &pathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.SetWordListPath(tt.args.wordListPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWordListPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterHyphenationDecompounder_Source(t *testing.T) {
	number := uint32(1)
	boolTrue := true
	path := "/temp/test.txt"

	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeHyphenationDecompounder,
			},
		},
		{
			name:   "hyphenation_patterns_path",
			fields: fields{hyphenationPatternsPath: &path},
			want: map[string]interface{}{
				"type":                      FilterTypeHyphenationDecompounder,
				"hyphenation_patterns_path": "/temp/test.txt",
			},
		},
		{
			name:   "word_list_path",
			fields: fields{wordListPath: &path},
			want: map[string]interface{}{
				"type":           FilterTypeHyphenationDecompounder,
				"word_list_path": "/temp/test.txt",
			},
		},
		{
			name:   "max_subword_size",
			fields: fields{maxSubwordSize: &number},
			want: map[string]interface{}{
				"type":             FilterTypeHyphenationDecompounder,
				"max_subword_size": uint32(1),
			},
		},
		{
			name:   "min_subword_size",
			fields: fields{minSubwordSize: &number},
			want: map[string]interface{}{
				"type":             FilterTypeHyphenationDecompounder,
				"min_subword_size": uint32(1),
			},
		},
		{
			name:   "min_word_size",
			fields: fields{minWordSize: &number},
			want: map[string]interface{}{
				"type":          FilterTypeHyphenationDecompounder,
				"min_word_size": uint32(1),
			},
		},
		{
			name:   "only_longest_match",
			fields: fields{onlyLongestMatch: &boolTrue},
			want: map[string]interface{}{
				"type":               FilterTypeHyphenationDecompounder,
				"only_longest_match": true,
			},
		},
		{
			name:   "word_list",
			fields: fields{wordList: []string{"test"}},
			want: map[string]interface{}{
				"type":      FilterTypeHyphenationDecompounder,
				"word_list": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			got, err := filter.Source()
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

func TestFilterHyphenationDecompounder_Type(t *testing.T) {
	type fields struct {
		name                    FilterName
		hyphenationPatternsPath *string
		wordListPath            *string
		maxSubwordSize          *uint32
		minSubwordSize          *uint32
		minWordSize             *uint32
		onlyLongestMatch        *bool
		wordList                []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeHyphenationDecompounder,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterHyphenationDecompounder{
				name:                    tt.fields.name,
				hyphenationPatternsPath: tt.fields.hyphenationPatternsPath,
				wordListPath:            tt.fields.wordListPath,
				maxSubwordSize:          tt.fields.maxSubwordSize,
				minSubwordSize:          tt.fields.minSubwordSize,
				minWordSize:             tt.fields.minWordSize,
				onlyLongestMatch:        tt.fields.onlyLongestMatch,
				wordList:                tt.fields.wordList,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeepTypes_AddTypes(t *testing.T) {
	type fields struct {
		name  FilterName
		mode  *FilterKeepTypesMode
		types []string
	}
	type args struct {
		types []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeepTypes
	}{
		{
			name: "set",
			args: args{types: []string{"test"}},
			want: &FilterKeepTypes{types: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{types: []string{"test"}},
			args:   args{types: []string{"test1"}},
			want:   &FilterKeepTypes{types: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeepTypes{
				name:  tt.fields.name,
				mode:  tt.fields.mode,
				types: tt.fields.types,
			}
			if got := filter.AddTypes(tt.args.types...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeepTypes_Name(t *testing.T) {
	type fields struct {
		name  FilterName
		mode  *FilterKeepTypesMode
		types []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeepTypes{
				name:  tt.fields.name,
				mode:  tt.fields.mode,
				types: tt.fields.types,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeepTypes_SetMode(t *testing.T) {
	keepTypesModeInclude := KeepTypesModeInclude
	keepTypesModeExclude := KeepTypesModeExclude

	type fields struct {
		name  FilterName
		mode  *FilterKeepTypesMode
		types []string
	}
	type args struct {
		mode FilterKeepTypesMode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeepTypes
	}{
		{
			name: "keep_types_mode_include",
			args: args{mode: KeepTypesModeInclude},
			want: &FilterKeepTypes{mode: &keepTypesModeInclude},
		},
		{
			name: "keep_types_mode_exclude",
			args: args{mode: KeepTypesModeExclude},
			want: &FilterKeepTypes{mode: &keepTypesModeExclude},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeepTypes{
				name:  tt.fields.name,
				mode:  tt.fields.mode,
				types: tt.fields.types,
			}
			if got := filter.SetMode(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeepTypes_Source(t *testing.T) {
	keepTypesModeInclude := KeepTypesModeInclude

	type fields struct {
		name  FilterName
		mode  *FilterKeepTypesMode
		types []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeKeepTypes,
			},
		},
		{
			name:   "mode",
			fields: fields{mode: &keepTypesModeInclude},
			want: map[string]interface{}{
				"type": FilterTypeKeepTypes,
				"mode": KeepTypesModeInclude,
			},
		},
		{
			name:   "types",
			fields: fields{types: []string{"test"}},
			want: map[string]interface{}{
				"type":  FilterTypeKeepTypes,
				"types": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeepTypes{
				name:  tt.fields.name,
				mode:  tt.fields.mode,
				types: tt.fields.types,
			}
			got, err := filter.Source()
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

func TestFilterKeepTypes_Type(t *testing.T) {
	type fields struct {
		name  FilterName
		mode  *FilterKeepTypesMode
		types []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeKeepTypes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeepTypes{
				name:  tt.fields.name,
				mode:  tt.fields.mode,
				types: tt.fields.types,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeep_AddKeepWords(t *testing.T) {
	type fields struct {
		name          FilterName
		keepWordsCase *bool
		keepWordsPath *string
		keepWords     []string
	}
	type args struct {
		keepWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeep
	}{
		{
			name: "set",
			args: args{keepWords: []string{"test"}},
			want: &FilterKeep{keepWords: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{keepWords: []string{"test"}},
			args:   args{keepWords: []string{"test1"}},
			want:   &FilterKeep{keepWords: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeep{
				name:          tt.fields.name,
				keepWordsCase: tt.fields.keepWordsCase,
				keepWordsPath: tt.fields.keepWordsPath,
				keepWords:     tt.fields.keepWords,
			}
			if got := filter.AddKeepWords(tt.args.keepWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddKeepWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeep_Name(t *testing.T) {
	type fields struct {
		name          FilterName
		keepWordsCase *bool
		keepWordsPath *string
		keepWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeep{
				name:          tt.fields.name,
				keepWordsCase: tt.fields.keepWordsCase,
				keepWordsPath: tt.fields.keepWordsPath,
				keepWords:     tt.fields.keepWords,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeep_SetKeepWordsCase(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name          FilterName
		keepWordsCase *bool
		keepWordsPath *string
		keepWords     []string
	}
	type args struct {
		keepWordsCase bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeep
	}{
		{
			name: "set",
			args: args{keepWordsCase: true},
			want: &FilterKeep{keepWordsCase: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{keepWordsCase: &boolTrue},
			args:   args{keepWordsCase: false},
			want:   &FilterKeep{keepWordsCase: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeep{
				name:          tt.fields.name,
				keepWordsCase: tt.fields.keepWordsCase,
				keepWordsPath: tt.fields.keepWordsPath,
				keepWords:     tt.fields.keepWords,
			}
			if got := filter.SetKeepWordsCase(tt.args.keepWordsCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetKeepWordsCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeep_SetKeepWordsPath(t *testing.T) {
	pathSet := "/temp/test.txt"
	pathChange := "/temp/test.v1.txt"

	type fields struct {
		name          FilterName
		keepWordsCase *bool
		keepWordsPath *string
		keepWords     []string
	}
	type args struct {
		keepWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeep
	}{
		{
			name: "set",
			args: args{keepWordsPath: "/temp/test.txt"},
			want: &FilterKeep{keepWordsPath: &pathSet},
		},
		{
			name:   "change",
			fields: fields{keepWordsPath: &pathSet},
			args:   args{keepWordsPath: "/temp/test.v1.txt"},
			want:   &FilterKeep{keepWordsPath: &pathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeep{
				name:          tt.fields.name,
				keepWordsCase: tt.fields.keepWordsCase,
				keepWordsPath: tt.fields.keepWordsPath,
				keepWords:     tt.fields.keepWords,
			}
			if got := filter.SetKeepWordsPath(tt.args.keepWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetKeepWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeep_Source(t *testing.T) {
	boolTrue := true
	path := "/temp/test.txt"

	type fields struct {
		name          FilterName
		keepWordsCase *bool
		keepWordsPath *string
		keepWords     []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeKeep,
			},
		},
		{
			name:   "keep_words_case",
			fields: fields{keepWordsCase: &boolTrue},
			want: map[string]interface{}{
				"type":            FilterTypeKeep,
				"keep_words_case": true,
			},
		},
		{
			name:   "keep_words_path",
			fields: fields{keepWordsPath: &path},
			want: map[string]interface{}{
				"type":            FilterTypeKeep,
				"keep_words_path": "/temp/test.txt",
			},
		},
		{
			name:   "keep_words",
			fields: fields{keepWords: []string{"test"}},
			want: map[string]interface{}{
				"type":       FilterTypeKeep,
				"keep_words": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeep{
				name:          tt.fields.name,
				keepWordsCase: tt.fields.keepWordsCase,
				keepWordsPath: tt.fields.keepWordsPath,
				keepWords:     tt.fields.keepWords,
			}
			got, err := filter.Source()
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

func TestFilterKeep_Type(t *testing.T) {
	type fields struct {
		name          FilterName
		keepWordsCase *bool
		keepWordsPath *string
		keepWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeKeep,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeep{
				name:          tt.fields.name,
				keepWordsCase: tt.fields.keepWordsCase,
				keepWordsPath: tt.fields.keepWordsPath,
				keepWords:     tt.fields.keepWords,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeywordMarker_AddKeywords(t *testing.T) {
	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	type args struct {
		keywords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeywordMarker
	}{
		{
			name: "set",
			args: args{keywords: []string{"test"}},
			want: &FilterKeywordMarker{keywords: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{keywords: []string{"test"}},
			args:   args{keywords: []string{"test1"}},
			want:   &FilterKeywordMarker{keywords: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			if got := filter.AddKeywords(tt.args.keywords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeywordMarker_Name(t *testing.T) {
	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeywordMarker_SetIgnoreCase(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	type args struct {
		ignoreCase bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeywordMarker
	}{
		{
			name: "set",
			args: args{ignoreCase: true},
			want: &FilterKeywordMarker{ignoreCase: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{ignoreCase: &boolTrue},
			args:   args{ignoreCase: false},
			want:   &FilterKeywordMarker{ignoreCase: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			if got := filter.SetIgnoreCase(tt.args.ignoreCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeywordMarker_SetKeywordsPath(t *testing.T) {
	pathSet := "/temp/test.txt"
	pathChange := "/temp/test.v1.txt"

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	type args struct {
		keywordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeywordMarker
	}{
		{
			name: "set",
			args: args{keywordsPath: "/temp/test.txt"},
			want: &FilterKeywordMarker{keywordsPath: &pathSet},
		},
		{
			name:   "change",
			fields: fields{keywordsPath: &pathSet},
			args:   args{keywordsPath: "/temp/test.v1.txt"},
			want:   &FilterKeywordMarker{keywordsPath: &pathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			if got := filter.SetKeywordsPath(tt.args.keywordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetKeywordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeywordMarker_SetKeywordsPattern(t *testing.T) {
	patternSet := "test"
	patternChange := "test1"

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	type args struct {
		keywordsPattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterKeywordMarker
	}{
		{
			name: "set",
			args: args{keywordsPattern: "test"},
			want: &FilterKeywordMarker{keywordsPattern: &patternSet},
		},
		{
			name:   "change",
			fields: fields{keywordsPattern: &patternSet},
			args:   args{keywordsPattern: "test1"},
			want:   &FilterKeywordMarker{keywordsPattern: &patternChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			if got := filter.SetKeywordsPattern(tt.args.keywordsPattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetKeywordsPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterKeywordMarker_Source(t *testing.T) {
	boolTrue := true
	path := "/temp/test.txt"
	pattern := "test"

	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeKeywordMarker,
			},
		},
		{
			name:   "ignore_case",
			fields: fields{ignoreCase: &boolTrue},
			want: map[string]interface{}{
				"type":        FilterTypeKeywordMarker,
				"ignore_case": true,
			},
		},
		{
			name:   "keywords_path",
			fields: fields{keywordsPath: &path},
			want: map[string]interface{}{
				"type":          FilterTypeKeywordMarker,
				"keywords_path": "/temp/test.txt",
			},
		},
		{
			name:   "keywords_pattern",
			fields: fields{keywordsPattern: &pattern},
			want: map[string]interface{}{
				"type":             FilterTypeKeywordMarker,
				"keywords_pattern": "test",
			},
		},
		{
			name:   "keywords",
			fields: fields{keywords: []string{"test"}},
			want: map[string]interface{}{
				"type":     FilterTypeKeywordMarker,
				"keywords": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			got, err := filter.Source()
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

func TestFilterKeywordMarker_Type(t *testing.T) {
	type fields struct {
		name            FilterName
		ignoreCase      *bool
		keywordsPath    *string
		keywordsPattern *string
		keywords        []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeKeywordMarker,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterKeywordMarker{
				name:            tt.fields.name,
				ignoreCase:      tt.fields.ignoreCase,
				keywordsPath:    tt.fields.keywordsPath,
				keywordsPattern: tt.fields.keywordsPattern,
				keywords:        tt.fields.keywords,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLength_Name(t *testing.T) {
	type fields struct {
		name FilterName
		min  *uint32
		max  *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLength{
				name: tt.fields.name,
				min:  tt.fields.min,
				max:  tt.fields.max,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLength_SetMax(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name FilterName
		min  *uint32
		max  *uint32
	}
	type args struct {
		max uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterLength
	}{
		{
			name: "set",
			args: args{max: 1},
			want: &FilterLength{max: &numberSet},
		},
		{
			name:   "change",
			fields: fields{max: &numberSet},
			args:   args{max: 2},
			want:   &FilterLength{max: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLength{
				name: tt.fields.name,
				min:  tt.fields.min,
				max:  tt.fields.max,
			}
			if got := filter.SetMax(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLength_SetMin(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name FilterName
		min  *uint32
		max  *uint32
	}
	type args struct {
		min uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterLength
	}{
		{
			name: "set",
			args: args{min: 1},
			want: &FilterLength{min: &numberSet},
		},
		{
			name:   "change",
			fields: fields{min: &numberSet},
			args:   args{min: 2},
			want:   &FilterLength{min: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLength{
				name: tt.fields.name,
				min:  tt.fields.min,
				max:  tt.fields.max,
			}
			if got := filter.SetMin(tt.args.min); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLength_Source(t *testing.T) {
	number := uint32(1)

	type fields struct {
		name FilterName
		min  *uint32
		max  *uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeLength,
			},
		},
		{
			name:   "min",
			fields: fields{min: &number},
			want: map[string]interface{}{
				"type": FilterTypeLength,
				"min":  uint32(1),
			},
		},
		{
			name:   "max",
			fields: fields{max: &number},
			want: map[string]interface{}{
				"type": FilterTypeLength,
				"max":  uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLength{
				name: tt.fields.name,
				min:  tt.fields.min,
				max:  tt.fields.max,
			}
			got, err := filter.Source()
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

func TestFilterLength_Type(t *testing.T) {
	type fields struct {
		name FilterName
		min  *uint32
		max  *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLength{
				name: tt.fields.name,
				min:  tt.fields.min,
				max:  tt.fields.max,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLimit_Name(t *testing.T) {
	type fields struct {
		name             FilterName
		maxTokenCount    *uint32
		consumeAllTokens *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLimit{
				name:             tt.fields.name,
				maxTokenCount:    tt.fields.maxTokenCount,
				consumeAllTokens: tt.fields.consumeAllTokens,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLimit_SetConsumeAllTokens(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             FilterName
		maxTokenCount    *uint32
		consumeAllTokens *bool
	}
	type args struct {
		consumeAllTokens bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterLimit
	}{
		{
			name: "set",
			args: args{consumeAllTokens: true},
			want: &FilterLimit{consumeAllTokens: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{consumeAllTokens: &boolTrue},
			args:   args{consumeAllTokens: false},
			want:   &FilterLimit{consumeAllTokens: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLimit{
				name:             tt.fields.name,
				maxTokenCount:    tt.fields.maxTokenCount,
				consumeAllTokens: tt.fields.consumeAllTokens,
			}
			if got := filter.SetConsumeAllTokens(tt.args.consumeAllTokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConsumeAllTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLimit_SetMaxTokenCount(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name             FilterName
		maxTokenCount    *uint32
		consumeAllTokens *bool
	}
	type args struct {
		maxTokenCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterLimit
	}{
		{
			name: "set",
			args: args{maxTokenCount: 1},
			want: &FilterLimit{maxTokenCount: &numberSet},
		},
		{
			name:   "change",
			fields: fields{maxTokenCount: &numberSet},
			args:   args{maxTokenCount: 2},
			want:   &FilterLimit{maxTokenCount: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLimit{
				name:             tt.fields.name,
				maxTokenCount:    tt.fields.maxTokenCount,
				consumeAllTokens: tt.fields.consumeAllTokens,
			}
			if got := filter.SetMaxTokenCount(tt.args.maxTokenCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxTokenCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLimit_Source(t *testing.T) {
	boolTrue := true
	number := uint32(1)

	type fields struct {
		name             FilterName
		maxTokenCount    *uint32
		consumeAllTokens *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeLimit,
			},
		},
		{
			name:   "max_token_count",
			fields: fields{maxTokenCount: &number},
			want: map[string]interface{}{
				"type":            FilterTypeLimit,
				"max_token_count": uint32(1),
			},
		},
		{
			name:   "consume_all_tokens",
			fields: fields{consumeAllTokens: &boolTrue},
			want: map[string]interface{}{
				"type":               FilterTypeLimit,
				"consume_all_tokens": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLimit{
				name:             tt.fields.name,
				maxTokenCount:    tt.fields.maxTokenCount,
				consumeAllTokens: tt.fields.consumeAllTokens,
			}
			got, err := filter.Source()
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

func TestFilterLimit_Type(t *testing.T) {
	type fields struct {
		name             FilterName
		maxTokenCount    *uint32
		consumeAllTokens *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeLimit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLimit{
				name:             tt.fields.name,
				maxTokenCount:    tt.fields.maxTokenCount,
				consumeAllTokens: tt.fields.consumeAllTokens,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLowercase_Name(t *testing.T) {
	type fields struct {
		name     FilterName
		language *FilterLowercaseLanguageType
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLowercase{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLowercase_SetLanguage(t *testing.T) {
	lowercaseLanguageTypeGreek := LowercaseLanguageTypeGreek
	lowercaseLanguageTypeIrish := LowercaseLanguageTypeIrish
	lowercaseLanguageTypeTurkish := LowercaseLanguageTypeTurkish

	type fields struct {
		name     FilterName
		language *FilterLowercaseLanguageType
	}
	type args struct {
		language FilterLowercaseLanguageType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterLowercase
	}{
		{
			name: "lowercase_language_type_greek",
			args: args{language: LowercaseLanguageTypeGreek},
			want: &FilterLowercase{language: &lowercaseLanguageTypeGreek},
		},
		{
			name: "lowercase_language_type_irish",
			args: args{language: LowercaseLanguageTypeIrish},
			want: &FilterLowercase{language: &lowercaseLanguageTypeIrish},
		},
		{
			name: "lowercase_language_type_turkish",
			args: args{language: LowercaseLanguageTypeTurkish},
			want: &FilterLowercase{language: &lowercaseLanguageTypeTurkish},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLowercase{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.SetLanguage(tt.args.language); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLowercase_Source(t *testing.T) {
	lowercaseLanguageTypeGreek := LowercaseLanguageTypeGreek

	type fields struct {
		name     FilterName
		language *FilterLowercaseLanguageType
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeLowercase,
			},
		},
		{
			name:   "language",
			fields: fields{language: &lowercaseLanguageTypeGreek},
			want: map[string]interface{}{
				"type":     FilterTypeLowercase,
				"language": LowercaseLanguageTypeGreek,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLowercase{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			got, err := filter.Source()
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

func TestFilterLowercase_Type(t *testing.T) {
	type fields struct {
		name     FilterName
		language *FilterLowercaseLanguageType
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeLowercase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterLowercase{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMinHash_Name(t *testing.T) {
	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMinHash_SetBucketCount(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	type args struct {
		bucketCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterMinHash
	}{
		{
			name: "set",
			args: args{bucketCount: 1},
			want: &FilterMinHash{bucketCount: &numberSet},
		},
		{
			name:   "change",
			fields: fields{bucketCount: &numberSet},
			args:   args{bucketCount: 2},
			want:   &FilterMinHash{bucketCount: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			if got := filter.SetBucketCount(tt.args.bucketCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBucketCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMinHash_SetHashCount(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	type args struct {
		hashCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterMinHash
	}{
		{
			name: "set",
			args: args{hashCount: 1},
			want: &FilterMinHash{hashCount: &numberSet},
		},
		{
			name:   "change",
			fields: fields{hashCount: &numberSet},
			args:   args{hashCount: 2},
			want:   &FilterMinHash{hashCount: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			if got := filter.SetHashCount(tt.args.hashCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHashCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMinHash_SetHashSetSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	type args struct {
		hashSetSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterMinHash
	}{
		{
			name: "set",
			args: args{hashSetSize: 1},
			want: &FilterMinHash{hashSetSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{hashSetSize: &numberSet},
			args:   args{hashSetSize: 2},
			want:   &FilterMinHash{hashSetSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			if got := filter.SetHashSetSize(tt.args.hashSetSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHashSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMinHash_SetWithRotation(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	type args struct {
		withRotation bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterMinHash
	}{
		{
			name: "set",
			args: args{withRotation: true},
			want: &FilterMinHash{withRotation: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{withRotation: &boolTrue},
			args:   args{withRotation: false},
			want:   &FilterMinHash{withRotation: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			if got := filter.SetWithRotation(tt.args.withRotation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWithRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMinHash_Source(t *testing.T) {
	number := uint32(1)
	boolTrue := true

	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeMinHash,
			},
		},
		{
			name:   "bucket_count",
			fields: fields{bucketCount: &number},
			want: map[string]interface{}{
				"type":         FilterTypeMinHash,
				"bucket_count": uint32(1),
			},
		},
		{
			name:   "hash_count",
			fields: fields{hashCount: &number},
			want: map[string]interface{}{
				"type":       FilterTypeMinHash,
				"hash_count": uint32(1),
			},
		},
		{
			name:   "hash_set_size",
			fields: fields{hashSetSize: &number},
			want: map[string]interface{}{
				"type":          FilterTypeMinHash,
				"hash_set_size": uint32(1),
			},
		},
		{
			name:   "with_rotation",
			fields: fields{withRotation: &boolTrue},
			want: map[string]interface{}{
				"type":          FilterTypeMinHash,
				"with_rotation": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			got, err := filter.Source()
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

func TestFilterMinHash_Type(t *testing.T) {
	type fields struct {
		name         FilterName
		bucketCount  *uint32
		hashCount    *uint32
		hashSetSize  *uint32
		withRotation *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeMinHash,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMinHash{
				name:         tt.fields.name,
				bucketCount:  tt.fields.bucketCount,
				hashCount:    tt.fields.hashCount,
				hashSetSize:  tt.fields.hashSetSize,
				withRotation: tt.fields.withRotation,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMultiplexer_AddFilters(t *testing.T) {
	type fields struct {
		name    FilterName
		filters []FilterName
	}
	type args struct {
		filters []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterMultiplexer
	}{
		{
			name: "set",
			args: args{filters: []FilterName{FilterNameArabic}},
			want: &FilterMultiplexer{filters: []FilterName{FilterNameArabic}},
		},
		{
			name:   "add",
			fields: fields{filters: []FilterName{FilterNameAsciifolding}},
			args:   args{filters: []FilterName{FilterNameArabic}},
			want:   &FilterMultiplexer{filters: []FilterName{FilterNameAsciifolding, FilterNameArabic}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMultiplexer{
				name:    tt.fields.name,
				filters: tt.fields.filters,
			}
			if got := filter.AddFilters(tt.args.filters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMultiplexer_Name(t *testing.T) {
	type fields struct {
		name    FilterName
		filters []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMultiplexer{
				name:    tt.fields.name,
				filters: tt.fields.filters,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMultiplexer_Source(t *testing.T) {
	type fields struct {
		name    FilterName
		filters []FilterName
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeMultiplexer,
			},
		},
		{
			name:   "filters",
			fields: fields{filters: []FilterName{FilterNameArabic}},
			want: map[string]interface{}{
				"type":    FilterTypeMultiplexer,
				"filters": []FilterName{FilterNameArabic},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMultiplexer{
				name:    tt.fields.name,
				filters: tt.fields.filters,
			}
			got, err := filter.Source()
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

func TestFilterMultiplexer_Type(t *testing.T) {
	type fields struct {
		name    FilterName
		filters []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeMultiplexer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterMultiplexer{
				name:    tt.fields.name,
				filters: tt.fields.filters,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNgram_Name(t *testing.T) {
	type fields struct {
		name             FilterName
		maxGram          *uint8
		minGram          *uint8
		preserveOriginal *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNgram_SetMaxGram(t *testing.T) {
	numberSet := uint8(1)
	numberChange := uint8(2)

	type fields struct {
		name             FilterName
		maxGram          *uint8
		minGram          *uint8
		preserveOriginal *bool
	}
	type args struct {
		maxGram uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterNgram
	}{
		{
			name: "set",
			args: args{maxGram: 1},
			want: &FilterNgram{maxGram: &numberSet},
		},
		{
			name:   "change",
			fields: fields{maxGram: &numberSet},
			args:   args{maxGram: 2},
			want:   &FilterNgram{maxGram: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.SetMaxGram(tt.args.maxGram); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNgram_SetMinGram(t *testing.T) {
	numberSet := uint8(1)
	numberChange := uint8(2)

	type fields struct {
		name             FilterName
		maxGram          *uint8
		minGram          *uint8
		preserveOriginal *bool
	}
	type args struct {
		minGram uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterNgram
	}{
		{
			name: "set",
			args: args{minGram: 1},
			want: &FilterNgram{minGram: &numberSet},
		},
		{
			name:   "change",
			fields: fields{minGram: &numberSet},
			args:   args{minGram: 2},
			want:   &FilterNgram{minGram: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.SetMinGram(tt.args.minGram); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNgram_SetPreserveOriginal(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             FilterName
		maxGram          *uint8
		minGram          *uint8
		preserveOriginal *bool
	}
	type args struct {
		preserveOriginal bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterNgram
	}{
		{
			name: "set",
			args: args{preserveOriginal: true},
			want: &FilterNgram{preserveOriginal: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{preserveOriginal: &boolTrue},
			args:   args{preserveOriginal: false},
			want:   &FilterNgram{preserveOriginal: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.SetPreserveOriginal(tt.args.preserveOriginal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPreserveOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNgram_Source(t *testing.T) {
	boolTrue := true
	number := uint8(1)

	type fields struct {
		name             FilterName
		maxGram          *uint8
		minGram          *uint8
		preserveOriginal *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeNgram,
			},
		},
		{
			name:   "max_gram",
			fields: fields{maxGram: &number},
			want: map[string]interface{}{
				"type":     FilterTypeNgram,
				"max_gram": uint8(1),
			},
		},
		{
			name:   "min_gram",
			fields: fields{minGram: &number},
			want: map[string]interface{}{
				"type":     FilterTypeNgram,
				"min_gram": uint8(1),
			},
		},
		{
			name:   "preserve_original",
			fields: fields{preserveOriginal: &boolTrue},
			want: map[string]interface{}{
				"type":              FilterTypeNgram,
				"preserve_original": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			got, err := filter.Source()
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

func TestFilterNgram_Type(t *testing.T) {
	type fields struct {
		name             FilterName
		maxGram          *uint8
		minGram          *uint8
		preserveOriginal *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeNgram,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterNgram{
				name:             tt.fields.name,
				maxGram:          tt.fields.maxGram,
				minGram:          tt.fields.minGram,
				preserveOriginal: tt.fields.preserveOriginal,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternCapture_AddPatterns(t *testing.T) {
	type fields struct {
		name             FilterName
		preserveOriginal *bool
		patterns         []string
	}
	type args struct {
		patterns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPatternCapture
	}{
		{
			name: "set",
			args: args{patterns: []string{"test"}},
			want: &FilterPatternCapture{patterns: []string{"test"}},
		},
		{
			name:   "add",
			fields: fields{patterns: []string{"test"}},
			args:   args{patterns: []string{"test1"}},
			want:   &FilterPatternCapture{patterns: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternCapture{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
				patterns:         tt.fields.patterns,
			}
			if got := filter.AddPatterns(tt.args.patterns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternCapture_Name(t *testing.T) {
	type fields struct {
		name             FilterName
		preserveOriginal *bool
		patterns         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternCapture{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
				patterns:         tt.fields.patterns,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternCapture_SetPreserveOriginal(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name             FilterName
		preserveOriginal *bool
		patterns         []string
	}
	type args struct {
		preserveOriginal bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPatternCapture
	}{
		{
			name: "set",
			args: args{preserveOriginal: true},
			want: &FilterPatternCapture{preserveOriginal: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{preserveOriginal: &boolTrue},
			args:   args{preserveOriginal: false},
			want:   &FilterPatternCapture{preserveOriginal: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternCapture{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
				patterns:         tt.fields.patterns,
			}
			if got := filter.SetPreserveOriginal(tt.args.preserveOriginal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPreserveOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternCapture_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		name             FilterName
		preserveOriginal *bool
		patterns         []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypePatternCapture,
			},
		},
		{
			name:   "preserve_original",
			fields: fields{preserveOriginal: &boolTrue},
			want: map[string]interface{}{
				"type":              FilterTypePatternCapture,
				"preserve_original": true,
			},
		},
		{
			name:   "patterns",
			fields: fields{patterns: []string{"test"}},
			want: map[string]interface{}{
				"type":     FilterTypePatternCapture,
				"patterns": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternCapture{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
				patterns:         tt.fields.patterns,
			}
			got, err := filter.Source()
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

func TestFilterPatternCapture_Type(t *testing.T) {
	type fields struct {
		name             FilterName
		preserveOriginal *bool
		patterns         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypePatternCapture,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternCapture{
				name:             tt.fields.name,
				preserveOriginal: tt.fields.preserveOriginal,
				patterns:         tt.fields.patterns,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternReplace_Name(t *testing.T) {
	type fields struct {
		name        FilterName
		all         *bool
		pattern     *string
		replacement *string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternReplace{
				name:        tt.fields.name,
				all:         tt.fields.all,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternReplace_SetAll(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name        FilterName
		all         *bool
		pattern     *string
		replacement *string
	}
	type args struct {
		all bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPatternReplace
	}{
		{
			name: "set",
			args: args{all: true},
			want: &FilterPatternReplace{all: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{all: &boolTrue},
			args:   args{all: false},
			want:   &FilterPatternReplace{all: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternReplace{
				name:        tt.fields.name,
				all:         tt.fields.all,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
			}
			if got := filter.SetAll(tt.args.all); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternReplace_SetPattern(t *testing.T) {
	patternSet := "test"
	patternChange := "test1"

	type fields struct {
		name        FilterName
		all         *bool
		pattern     *string
		replacement *string
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPatternReplace
	}{
		{
			name: "set",
			args: args{pattern: patternSet},
			want: &FilterPatternReplace{pattern: &patternSet},
		},
		{
			name:   "change",
			fields: fields{pattern: &patternSet},
			args:   args{pattern: patternChange},
			want:   &FilterPatternReplace{pattern: &patternChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternReplace{
				name:        tt.fields.name,
				all:         tt.fields.all,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
			}
			if got := filter.SetPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternReplace_SetReplacement(t *testing.T) {
	patternSet := "test"
	patternChange := "test1"

	type fields struct {
		name        FilterName
		all         *bool
		pattern     *string
		replacement *string
	}
	type args struct {
		replacement string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPatternReplace
	}{
		{
			name: "set",
			args: args{replacement: patternSet},
			want: &FilterPatternReplace{replacement: &patternSet},
		},
		{
			name:   "change",
			fields: fields{replacement: &patternSet},
			args:   args{replacement: patternChange},
			want:   &FilterPatternReplace{replacement: &patternChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternReplace{
				name:        tt.fields.name,
				all:         tt.fields.all,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
			}
			if got := filter.SetReplacement(tt.args.replacement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPatternReplace_Source(t *testing.T) {
	boolTrue := true
	testString := "test"

	type fields struct {
		name        FilterName
		all         *bool
		pattern     *string
		replacement *string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypePatternReplace,
			},
		},
		{
			name:   "all",
			fields: fields{all: &boolTrue},
			want: map[string]interface{}{
				"type": FilterTypePatternReplace,
				"all":  true,
			},
		},
		{
			name:   "pattern",
			fields: fields{pattern: &testString},
			want: map[string]interface{}{
				"type":    FilterTypePatternReplace,
				"pattern": "test",
			},
		},
		{
			name:   "replacement",
			fields: fields{replacement: &testString},
			want: map[string]interface{}{
				"type":        FilterTypePatternReplace,
				"replacement": "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternReplace{
				name:        tt.fields.name,
				all:         tt.fields.all,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
			}
			got, err := filter.Source()
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

func TestFilterPatternReplace_Type(t *testing.T) {
	type fields struct {
		name        FilterName
		all         *bool
		pattern     *string
		replacement *string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypePatternReplace,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPatternReplace{
				name:        tt.fields.name,
				all:         tt.fields.all,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_Name(t *testing.T) {
	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_SetEncoder(t *testing.T) {
	phoneticEncoderMetaphone := PhoneticEncoderMetaphone
	phoneticEncoderDoubleMetaphone := PhoneticEncoderDoubleMetaphone
	phoneticEncoderSoundex := PhoneticEncoderSoundex
	phoneticEncoderRefinedSoundex := PhoneticEncoderRefinedSoundex
	phoneticEncoderCaverphone1 := PhoneticEncoderCaverphone1
	phoneticEncoderCaverphone2 := PhoneticEncoderCaverphone2
	phoneticEncoderCologne := PhoneticEncoderCologne
	phoneticEncoderNysiis := PhoneticEncoderNysiis
	phoneticEncoderKoelnerphonetik := PhoneticEncoderKoelnerphonetik
	phoneticEncoderHaasephonetik := PhoneticEncoderHaasephonetik
	phoneticEncoderBeiderMorse := PhoneticEncoderBeiderMorse
	phoneticEncoderDaitchMokotoff := PhoneticEncoderDaitchMokotoff

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	type args struct {
		encoder FilterPhoneticEncoder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPhonetic
	}{
		{
			name: "phonetic_encoder_metaphone",
			args: args{encoder: PhoneticEncoderMetaphone},
			want: &FilterPhonetic{encoder: &phoneticEncoderMetaphone},
		},
		{
			name: "phonetic_encoder_double_metaphone",
			args: args{encoder: PhoneticEncoderDoubleMetaphone},
			want: &FilterPhonetic{encoder: &phoneticEncoderDoubleMetaphone},
		},
		{
			name: "phonetic_encoder_soundex",
			args: args{encoder: PhoneticEncoderSoundex},
			want: &FilterPhonetic{encoder: &phoneticEncoderSoundex},
		},
		{
			name: "phonetic_encoder_refined_soundex",
			args: args{encoder: PhoneticEncoderRefinedSoundex},
			want: &FilterPhonetic{encoder: &phoneticEncoderRefinedSoundex},
		},
		{
			name: "phonetic_encoder_caverphone1",
			args: args{encoder: PhoneticEncoderCaverphone1},
			want: &FilterPhonetic{encoder: &phoneticEncoderCaverphone1},
		},
		{
			name: "phonetic_encoder_caverphone2",
			args: args{encoder: PhoneticEncoderCaverphone2},
			want: &FilterPhonetic{encoder: &phoneticEncoderCaverphone2},
		},
		{
			name: "phonetic_encoder_cologne",
			args: args{encoder: PhoneticEncoderCologne},
			want: &FilterPhonetic{encoder: &phoneticEncoderCologne},
		},
		{
			name: "phonetic_encoder_nysiis",
			args: args{encoder: PhoneticEncoderNysiis},
			want: &FilterPhonetic{encoder: &phoneticEncoderNysiis},
		},
		{
			name: "phonetic_encoder_koelnerphonetik",
			args: args{encoder: PhoneticEncoderKoelnerphonetik},
			want: &FilterPhonetic{encoder: &phoneticEncoderKoelnerphonetik},
		},
		{
			name: "phonetic_encoder_haasephonetik",
			args: args{encoder: PhoneticEncoderHaasephonetik},
			want: &FilterPhonetic{encoder: &phoneticEncoderHaasephonetik},
		},
		{
			name: "phonetic_encoder_beider_morse",
			args: args{encoder: PhoneticEncoderBeiderMorse},
			want: &FilterPhonetic{encoder: &phoneticEncoderBeiderMorse},
		},
		{
			name: "phonetic_encoder_daitch_mokotoff",
			args: args{encoder: PhoneticEncoderDaitchMokotoff},
			want: &FilterPhonetic{encoder: &phoneticEncoderDaitchMokotoff},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.SetEncoder(tt.args.encoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_SetLanguageSet(t *testing.T) {
	phoneticLanguageSetAny := PhoneticLanguageSetAny
	phoneticLanguageSetCommon := PhoneticLanguageSetCommon
	phoneticLanguageSetCyrillic := PhoneticLanguageSetCyrillic
	phoneticLanguageSetEnglish := PhoneticLanguageSetEnglish
	phoneticLanguageSetFrench := PhoneticLanguageSetFrench
	phoneticLanguageSetGerman := PhoneticLanguageSetGerman
	phoneticLanguageSetHebrew := PhoneticLanguageSetHebrew
	phoneticLanguageSetHungarian := PhoneticLanguageSetHungarian
	phoneticLanguageSetPolish := PhoneticLanguageSetPolish
	phoneticLanguageSetRomanian := PhoneticLanguageSetRomanian
	phoneticLanguageSetRussian := PhoneticLanguageSetRussian
	phoneticLanguageSetSpanish := PhoneticLanguageSetSpanish

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	type args struct {
		languageSet FilterPhoneticLanguageSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPhonetic
	}{
		{
			name: "phonetic_language_set_any",
			args: args{languageSet: PhoneticLanguageSetAny},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetAny},
		},
		{
			name: "phonetic_language_set_common",
			args: args{languageSet: PhoneticLanguageSetCommon},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetCommon},
		},
		{
			name: "phonetic_language_set_cyrillic",
			args: args{languageSet: PhoneticLanguageSetCyrillic},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetCyrillic},
		},
		{
			name: "phonetic_language_set_english",
			args: args{languageSet: PhoneticLanguageSetEnglish},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetEnglish},
		},
		{
			name: "phonetic_language_set_french",
			args: args{languageSet: PhoneticLanguageSetFrench},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetFrench},
		},
		{
			name: "phonetic_language_set_german",
			args: args{languageSet: PhoneticLanguageSetGerman},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetGerman},
		},
		{
			name: "phonetic_language_set_hebrew",
			args: args{languageSet: PhoneticLanguageSetHebrew},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetHebrew},
		},
		{
			name: "phonetic_language_set_hungarian",
			args: args{languageSet: PhoneticLanguageSetHungarian},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetHungarian},
		},
		{
			name: "phonetic_language_set_polish",
			args: args{languageSet: PhoneticLanguageSetPolish},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetPolish},
		},
		{
			name: "phonetic_language_set_romanian",
			args: args{languageSet: PhoneticLanguageSetRomanian},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetRomanian},
		},
		{
			name: "phonetic_language_set_russian",
			args: args{languageSet: PhoneticLanguageSetRussian},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetRussian},
		},
		{
			name: "phonetic_language_set_spanish",
			args: args{languageSet: PhoneticLanguageSetSpanish},
			want: &FilterPhonetic{languageSet: &phoneticLanguageSetSpanish},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.SetLanguageSet(tt.args.languageSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLanguageSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_SetMaxCodeLen(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	type args struct {
		maxCodeLen uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPhonetic
	}{
		{
			name: "set",
			args: args{maxCodeLen: 1},
			want: &FilterPhonetic{maxCodeLen: &numberSet},
		},
		{
			name:   "change",
			fields: fields{maxCodeLen: &numberSet},
			args:   args{maxCodeLen: 2},
			want:   &FilterPhonetic{maxCodeLen: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.SetMaxCodeLen(tt.args.maxCodeLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxCodeLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_SetNameType(t *testing.T) {
	phoneticNameTypeAshkenazi := PhoneticNameTypeAshkenazi
	phoneticNameTypeSephardic := PhoneticNameTypeSephardic
	phoneticNameTypeGeneric := PhoneticNameTypeGeneric

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	type args struct {
		nameType FilterPhoneticNameType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPhonetic
	}{
		{
			name: "phonetic_name_type_ashkenazi",
			args: args{nameType: PhoneticNameTypeAshkenazi},
			want: &FilterPhonetic{nameType: &phoneticNameTypeAshkenazi},
		},
		{
			name: "phonetic_name_type_sephardic",
			args: args{nameType: PhoneticNameTypeSephardic},
			want: &FilterPhonetic{nameType: &phoneticNameTypeSephardic},
		},
		{
			name: "phonetic_name_type_generic",
			args: args{nameType: PhoneticNameTypeGeneric},
			want: &FilterPhonetic{nameType: &phoneticNameTypeGeneric},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.SetNameType(tt.args.nameType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNameType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_SetReplace(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	type args struct {
		replace bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPhonetic
	}{
		{
			name: "set",
			args: args{replace: true},
			want: &FilterPhonetic{replace: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{replace: &boolTrue},
			args:   args{replace: false},
			want:   &FilterPhonetic{replace: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.SetReplace(tt.args.replace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_SetRuleType(t *testing.T) {
	phoneticRuleTypeExact := PhoneticRuleTypeExact
	phoneticRuleTypeApprox := PhoneticRuleTypeApprox

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	type args struct {
		ruleType FilterPhoneticRuleType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPhonetic
	}{
		{
			name: "phonetic_rule_type_exact",
			args: args{ruleType: PhoneticRuleTypeExact},
			want: &FilterPhonetic{ruleType: &phoneticRuleTypeExact},
		},
		{
			name: "phonetic_rule_type_approx",
			args: args{ruleType: PhoneticRuleTypeApprox},
			want: &FilterPhonetic{ruleType: &phoneticRuleTypeApprox},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.SetRuleType(tt.args.ruleType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRuleType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPhonetic_Source(t *testing.T) {
	phoneticEncoderMetaphone := PhoneticEncoderMetaphone
	phoneticRuleTypeExact := PhoneticRuleTypeExact
	phoneticNameTypeAshkenazi := PhoneticNameTypeAshkenazi
	phoneticLanguageSetAny := PhoneticLanguageSetAny
	boolTrue := true
	number := uint32(1)

	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypePhonetic,
			},
		},
		{
			name:   "encoder",
			fields: fields{encoder: &phoneticEncoderMetaphone},
			want: map[string]interface{}{
				"type":    FilterTypePhonetic,
				"encoder": PhoneticEncoderMetaphone,
			},
		},
		{
			name:   "rule_type",
			fields: fields{ruleType: &phoneticRuleTypeExact},
			want: map[string]interface{}{
				"type":      FilterTypePhonetic,
				"rule_type": PhoneticRuleTypeExact,
			},
		},
		{
			name:   "name_type",
			fields: fields{nameType: &phoneticNameTypeAshkenazi},
			want: map[string]interface{}{
				"type":      FilterTypePhonetic,
				"name_type": PhoneticNameTypeAshkenazi,
			},
		},
		{
			name:   "language_set",
			fields: fields{languageSet: &phoneticLanguageSetAny},
			want: map[string]interface{}{
				"type":         FilterTypePhonetic,
				"language_set": PhoneticLanguageSetAny,
			},
		},
		{
			name:   "replace",
			fields: fields{replace: &boolTrue},
			want: map[string]interface{}{
				"type":    FilterTypePhonetic,
				"replace": true,
			},
		},
		{
			name:   "max_code_len",
			fields: fields{maxCodeLen: &number},
			want: map[string]interface{}{
				"type":         FilterTypePhonetic,
				"max_code_len": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			got, err := filter.Source()
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

func TestFilterPhonetic_Type(t *testing.T) {
	type fields struct {
		name        FilterName
		encoder     *FilterPhoneticEncoder
		ruleType    *FilterPhoneticRuleType
		nameType    *FilterPhoneticNameType
		languageSet *FilterPhoneticLanguageSet
		replace     *bool
		maxCodeLen  *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypePhonetic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPhonetic{
				name:        tt.fields.name,
				encoder:     tt.fields.encoder,
				ruleType:    tt.fields.ruleType,
				nameType:    tt.fields.nameType,
				languageSet: tt.fields.languageSet,
				replace:     tt.fields.replace,
				maxCodeLen:  tt.fields.maxCodeLen,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPredicateTokenFilter_Name(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPredicateTokenFilter{
				name:   tt.fields.name,
				script: tt.fields.script,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPredicateTokenFilter_SetScript(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
	}
	type args struct {
		script *elastic.Script
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterPredicateTokenFilter
	}{
		{
			name: "set",
			args: args{script: elastic.NewScript("test")},
			want: &FilterPredicateTokenFilter{script: elastic.NewScript("test")},
		},
		{
			name:   "change",
			fields: fields{script: elastic.NewScript("test")},
			args:   args{script: elastic.NewScript("test1")},
			want:   &FilterPredicateTokenFilter{script: elastic.NewScript("test1")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPredicateTokenFilter{
				name:   tt.fields.name,
				script: tt.fields.script,
			}
			if got := filter.SetScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPredicateTokenFilter_Source(t *testing.T) {
	scriptSrc, _ := elastic.NewScript("test").Source()

	type fields struct {
		name   FilterName
		script *elastic.Script
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypePredicateTokenFilter,
			},
		},
		{
			name:   "script",
			fields: fields{script: elastic.NewScript("test")},
			want: map[string]interface{}{
				"type":   FilterTypePredicateTokenFilter,
				"script": scriptSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPredicateTokenFilter{
				name:   tt.fields.name,
				script: tt.fields.script,
			}
			got, err := filter.Source()
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

func TestFilterPredicateTokenFilter_Type(t *testing.T) {
	type fields struct {
		name   FilterName
		script *elastic.Script
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypePredicateTokenFilter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterPredicateTokenFilter{
				name:   tt.fields.name,
				script: tt.fields.script,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_Name(t *testing.T) {
	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_SetFillerToken(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	type args struct {
		fillerToken string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterShingle
	}{
		{
			name: "set",
			args: args{fillerToken: "test"},
			want: &FilterShingle{fillerToken: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{fillerToken: &testStringSet},
			args:   args{fillerToken: "test1"},
			want:   &FilterShingle{fillerToken: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.SetFillerToken(tt.args.fillerToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFillerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_SetMaxShingleSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	type args struct {
		maxShingleSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterShingle
	}{
		{
			name: "set",
			args: args{maxShingleSize: 1},
			want: &FilterShingle{maxShingleSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{maxShingleSize: &numberSet},
			args:   args{maxShingleSize: 2},
			want:   &FilterShingle{maxShingleSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.SetMaxShingleSize(tt.args.maxShingleSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxShingleSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_SetMinShingleSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	type args struct {
		minShingleSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterShingle
	}{
		{
			name: "set",
			args: args{minShingleSize: 1},
			want: &FilterShingle{minShingleSize: &numberSet},
		},
		{
			name:   "change",
			fields: fields{minShingleSize: &numberSet},
			args:   args{minShingleSize: 2},
			want:   &FilterShingle{minShingleSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.SetMinShingleSize(tt.args.minShingleSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinShingleSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_SetOutputUnigrams(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	type args struct {
		outputUnigrams bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterShingle
	}{
		{
			name: "set",
			args: args{outputUnigrams: true},
			want: &FilterShingle{outputUnigrams: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{outputUnigrams: &boolTrue},
			args:   args{outputUnigrams: false},
			want:   &FilterShingle{outputUnigrams: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.SetOutputUnigrams(tt.args.outputUnigrams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOutputUnigrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_SetOutputUnigramsIfNoShingles(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	type args struct {
		outputUnigramsIfNoShingles bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterShingle
	}{
		{
			name: "set",
			args: args{outputUnigramsIfNoShingles: true},
			want: &FilterShingle{outputUnigramsIfNoShingles: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{outputUnigramsIfNoShingles: &boolTrue},
			args:   args{outputUnigramsIfNoShingles: false},
			want:   &FilterShingle{outputUnigramsIfNoShingles: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.SetOutputUnigramsIfNoShingles(tt.args.outputUnigramsIfNoShingles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOutputUnigramsIfNoShingles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_SetTokenSeparator(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	type args struct {
		tokenSeparator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterShingle
	}{
		{
			name: "set",
			args: args{tokenSeparator: "test"},
			want: &FilterShingle{tokenSeparator: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{tokenSeparator: &testStringSet},
			args:   args{tokenSeparator: "test1"},
			want:   &FilterShingle{tokenSeparator: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.SetTokenSeparator(tt.args.tokenSeparator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTokenSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterShingle_Source(t *testing.T) {
	number := uint32(1)
	boolTrue := true
	testString := "test"

	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeShingle,
			},
		},
		{
			name:   "max_shingle_size",
			fields: fields{maxShingleSize: &number},
			want: map[string]interface{}{
				"type":             FilterTypeShingle,
				"max_shingle_size": uint32(1),
			},
		},
		{
			name:   "min_shingle_size",
			fields: fields{minShingleSize: &number},
			want: map[string]interface{}{
				"type":             FilterTypeShingle,
				"min_shingle_size": uint32(1),
			},
		},
		{
			name:   "output_unigrams",
			fields: fields{outputUnigrams: &boolTrue},
			want: map[string]interface{}{
				"type":            FilterTypeShingle,
				"output_unigrams": true,
			},
		},
		{
			name:   "output_unigrams_if_no_shingles",
			fields: fields{outputUnigramsIfNoShingles: &boolTrue},
			want: map[string]interface{}{
				"type":                           FilterTypeShingle,
				"output_unigrams_if_no_shingles": true,
			},
		},
		{
			name:   "token_separator",
			fields: fields{tokenSeparator: &testString},
			want: map[string]interface{}{
				"type":            FilterTypeShingle,
				"token_separator": "test",
			},
		},
		{
			name:   "filler_token",
			fields: fields{fillerToken: &testString},
			want: map[string]interface{}{
				"type":         FilterTypeShingle,
				"filler_token": "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			got, err := filter.Source()
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

func TestFilterShingle_Type(t *testing.T) {
	type fields struct {
		name                       FilterName
		maxShingleSize             *uint32
		minShingleSize             *uint32
		outputUnigrams             *bool
		outputUnigramsIfNoShingles *bool
		tokenSeparator             *string
		fillerToken                *string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeShingle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterShingle{
				name:                       tt.fields.name,
				maxShingleSize:             tt.fields.maxShingleSize,
				minShingleSize:             tt.fields.minShingleSize,
				outputUnigrams:             tt.fields.outputUnigrams,
				outputUnigramsIfNoShingles: tt.fields.outputUnigramsIfNoShingles,
				tokenSeparator:             tt.fields.tokenSeparator,
				fillerToken:                tt.fields.fillerToken,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSnowball_Name(t *testing.T) {
	type fields struct {
		name     FilterName
		language *FilterSnowballLanguage
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSnowball{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSnowball_SetLanguage(t *testing.T) {
	snowballLanguageArabic := SnowballLanguageArabic
	snowballLanguageArmenian := SnowballLanguageArmenian
	snowballLanguageBasque := SnowballLanguageBasque
	snowballLanguageCatalan := SnowballLanguageCatalan
	snowballLanguageDanish := SnowballLanguageDanish
	snowballLanguageDutch := SnowballLanguageDutch
	snowballLanguageEnglish := SnowballLanguageEnglish
	snowballLanguageEstonian := SnowballLanguageEstonian
	snowballLanguageFinnish := SnowballLanguageFinnish
	snowballLanguageFrench := SnowballLanguageFrench
	snowballLanguageGerman := SnowballLanguageGerman
	snowballLanguageGerman2 := SnowballLanguageGerman2
	snowballLanguageHungarian := SnowballLanguageHungarian
	snowballLanguageItalian := SnowballLanguageItalian
	snowballLanguageIrish := SnowballLanguageIrish
	snowballLanguageKp := SnowballLanguageKp
	snowballLanguageLithuanian := SnowballLanguageLithuanian
	snowballLanguageLovins := SnowballLanguageLovins
	snowballLanguageNorwegian := SnowballLanguageNorwegian
	snowballLanguagePorter := SnowballLanguagePorter
	snowballLanguagePortuguese := SnowballLanguagePortuguese
	snowballLanguageRomanian := SnowballLanguageRomanian
	snowballLanguageRussian := SnowballLanguageRussian
	snowballLanguageSpanish := SnowballLanguageSpanish
	snowballLanguageSwedish := SnowballLanguageSwedish
	snowballLanguageTurkish := SnowballLanguageTurkish

	type fields struct {
		name     FilterName
		language *FilterSnowballLanguage
	}
	type args struct {
		language FilterSnowballLanguage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterSnowball
	}{
		{
			name: "snowball_language_arabic",
			args: args{language: SnowballLanguageArabic},
			want: &FilterSnowball{language: &snowballLanguageArabic},
		},
		{
			name: "snowball_language_armenian",
			args: args{language: SnowballLanguageArmenian},
			want: &FilterSnowball{language: &snowballLanguageArmenian},
		},
		{
			name: "snowball_language_basque",
			args: args{language: SnowballLanguageBasque},
			want: &FilterSnowball{language: &snowballLanguageBasque},
		},
		{
			name: "snowball_language_catalan",
			args: args{language: SnowballLanguageCatalan},
			want: &FilterSnowball{language: &snowballLanguageCatalan},
		},
		{
			name: "snowball_language_danish",
			args: args{language: SnowballLanguageDanish},
			want: &FilterSnowball{language: &snowballLanguageDanish},
		},
		{
			name: "snowball_language_dutch",
			args: args{language: SnowballLanguageDutch},
			want: &FilterSnowball{language: &snowballLanguageDutch},
		},
		{
			name: "snowball_language_english",
			args: args{language: SnowballLanguageEnglish},
			want: &FilterSnowball{language: &snowballLanguageEnglish},
		},
		{
			name: "snowball_language_estonian",
			args: args{language: SnowballLanguageEstonian},
			want: &FilterSnowball{language: &snowballLanguageEstonian},
		},
		{
			name: "snowball_language_finnish",
			args: args{language: SnowballLanguageFinnish},
			want: &FilterSnowball{language: &snowballLanguageFinnish},
		},
		{
			name: "snowball_language_french",
			args: args{language: SnowballLanguageFrench},
			want: &FilterSnowball{language: &snowballLanguageFrench},
		},
		{
			name: "snowball_language_german",
			args: args{language: SnowballLanguageGerman},
			want: &FilterSnowball{language: &snowballLanguageGerman},
		},
		{
			name: "snowball_language_german2",
			args: args{language: SnowballLanguageGerman2},
			want: &FilterSnowball{language: &snowballLanguageGerman2},
		},
		{
			name: "snowball_language_hungarian",
			args: args{language: SnowballLanguageHungarian},
			want: &FilterSnowball{language: &snowballLanguageHungarian},
		},
		{
			name: "snowball_language_italian",
			args: args{language: SnowballLanguageItalian},
			want: &FilterSnowball{language: &snowballLanguageItalian},
		},
		{
			name: "snowball_language_irish",
			args: args{language: SnowballLanguageIrish},
			want: &FilterSnowball{language: &snowballLanguageIrish},
		},
		{
			name: "snowball_language_kp",
			args: args{language: SnowballLanguageKp},
			want: &FilterSnowball{language: &snowballLanguageKp},
		},
		{
			name: "snowball_language_lithuanian",
			args: args{language: SnowballLanguageLithuanian},
			want: &FilterSnowball{language: &snowballLanguageLithuanian},
		},
		{
			name: "snowball_language_lovins",
			args: args{language: SnowballLanguageLovins},
			want: &FilterSnowball{language: &snowballLanguageLovins},
		},
		{
			name: "snowball_language_norwegian",
			args: args{language: SnowballLanguageNorwegian},
			want: &FilterSnowball{language: &snowballLanguageNorwegian},
		},
		{
			name: "snowball_language_porter",
			args: args{language: SnowballLanguagePorter},
			want: &FilterSnowball{language: &snowballLanguagePorter},
		},
		{
			name: "snowball_language_portuguese",
			args: args{language: SnowballLanguagePortuguese},
			want: &FilterSnowball{language: &snowballLanguagePortuguese},
		},
		{
			name: "snowball_language_romanian",
			args: args{language: SnowballLanguageRomanian},
			want: &FilterSnowball{language: &snowballLanguageRomanian},
		},
		{
			name: "snowball_language_russian",
			args: args{language: SnowballLanguageRussian},
			want: &FilterSnowball{language: &snowballLanguageRussian},
		},
		{
			name: "snowball_language_spanish",
			args: args{language: SnowballLanguageSpanish},
			want: &FilterSnowball{language: &snowballLanguageSpanish},
		},
		{
			name: "snowball_language_swedish",
			args: args{language: SnowballLanguageSwedish},
			want: &FilterSnowball{language: &snowballLanguageSwedish},
		},
		{
			name: "snowball_language_turkish",
			args: args{language: SnowballLanguageTurkish},
			want: &FilterSnowball{language: &snowballLanguageTurkish},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSnowball{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.SetLanguage(tt.args.language); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSnowball_Source(t *testing.T) {
	snowballLanguageArabic := SnowballLanguageArabic

	type fields struct {
		name     FilterName
		language *FilterSnowballLanguage
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeSnowball,
			},
		},
		{
			name:   "language",
			fields: fields{language: &snowballLanguageArabic},
			want: map[string]interface{}{
				"type":     FilterTypeSnowball,
				"language": SnowballLanguageArabic,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSnowball{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			got, err := filter.Source()
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

func TestFilterSnowball_Type(t *testing.T) {
	type fields struct {
		name     FilterName
		language *FilterSnowballLanguage
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeSnowball,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSnowball{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmerOverride_AddRules(t *testing.T) {
	type fields struct {
		name      FilterName
		rulesPath *string
		rules     []*StemmerOverrideRule
	}
	type args struct {
		rules []*StemmerOverrideRule
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStemmerOverride
	}{
		{
			name: "set",
			args: args{rules: []*StemmerOverrideRule{{}}},
			want: &FilterStemmerOverride{rules: []*StemmerOverrideRule{{}}},
		},
		{
			name:   "add",
			fields: fields{rules: []*StemmerOverrideRule{{stemmer: "|"}}},
			args:   args{rules: []*StemmerOverrideRule{{}}},
			want:   &FilterStemmerOverride{rules: []*StemmerOverrideRule{{stemmer: "|"}, {}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmerOverride{
				name:      tt.fields.name,
				rulesPath: tt.fields.rulesPath,
				rules:     tt.fields.rules,
			}
			if got := filter.AddRules(tt.args.rules...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmerOverride_Name(t *testing.T) {
	type fields struct {
		name      FilterName
		rulesPath *string
		rules     []*StemmerOverrideRule
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmerOverride{
				name:      tt.fields.name,
				rulesPath: tt.fields.rulesPath,
				rules:     tt.fields.rules,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmerOverride_SetRulesPath(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		name      FilterName
		rulesPath *string
		rules     []*StemmerOverrideRule
	}
	type args struct {
		rulesPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStemmerOverride
	}{
		{
			name: "set",
			args: args{rulesPath: "test"},
			want: &FilterStemmerOverride{rulesPath: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{rulesPath: &testStringSet},
			args:   args{rulesPath: "test1"},
			want:   &FilterStemmerOverride{rulesPath: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmerOverride{
				name:      tt.fields.name,
				rulesPath: tt.fields.rulesPath,
				rules:     tt.fields.rules,
			}
			if got := filter.SetRulesPath(tt.args.rulesPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRulesPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmerOverride_Source(t *testing.T) {
	testString := "test"

	type fields struct {
		name      FilterName
		rulesPath *string
		rules     []*StemmerOverrideRule
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeStemmerOverride,
			},
		},
		{
			name:   "rules_path",
			fields: fields{rulesPath: &testString},
			want: map[string]interface{}{
				"type":       FilterTypeStemmerOverride,
				"rules_path": "test",
			},
		},
		{
			name:   "rules",
			fields: fields{rules: []*StemmerOverrideRule{{stemmer: "|", worlds: []string{"test", "test1"}}}},
			want: map[string]interface{}{
				"type":  FilterTypeStemmerOverride,
				"rules": []string{"test,test1 => |"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmerOverride{
				name:      tt.fields.name,
				rulesPath: tt.fields.rulesPath,
				rules:     tt.fields.rules,
			}
			got, err := filter.Source()
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

func TestFilterStemmerOverride_Type(t *testing.T) {
	type fields struct {
		name      FilterName
		rulesPath *string
		rules     []*StemmerOverrideRule
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeStemmerOverride,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmerOverride{
				name:      tt.fields.name,
				rulesPath: tt.fields.rulesPath,
				rules:     tt.fields.rules,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmer_Name(t *testing.T) {
	type fields struct {
		name     FilterName
		language *FilterStemmerLanguage
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmer{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmer_SetLanguage(t *testing.T) {
	stemmerLanguageArabic := StemmerLanguageArabic
	stemmerLanguageArmenian := StemmerLanguageArmenian
	stemmerLanguageBasque := StemmerLanguageBasque
	stemmerLanguageBengali := StemmerLanguageBengali
	stemmerLanguageBrazilian := StemmerLanguageBrazilian
	stemmerLanguageBulgarian := StemmerLanguageBulgarian
	stemmerLanguageCatalan := StemmerLanguageCatalan
	stemmerLanguageCzech := StemmerLanguageCzech
	stemmerLanguageDanish := StemmerLanguageDanish
	stemmerLanguageDutch := StemmerLanguageDutch
	stemmerLanguageDutchKp := StemmerLanguageDutchKp
	stemmerLanguageEnglish := StemmerLanguageEnglish
	stemmerLanguageLightEnglish := StemmerLanguageLightEnglish
	stemmerLanguageLovins := StemmerLanguageLovins
	stemmerLanguageMinimalEnglish := StemmerLanguageMinimalEnglish
	stemmerLanguagePorter2 := StemmerLanguagePorter2
	stemmerLanguagePossessiveEnglish := StemmerLanguagePossessiveEnglish
	stemmerLanguageEstonian := StemmerLanguageEstonian
	stemmerLanguageFinnish := StemmerLanguageFinnish
	stemmerLanguageLightFinnish := StemmerLanguageLightFinnish
	stemmerLanguageLightFrench := StemmerLanguageLightFrench
	stemmerLanguageFrench := StemmerLanguageFrench
	stemmerLanguageMinimalFrench := StemmerLanguageMinimalFrench
	stemmerLanguageGalician := StemmerLanguageGalician
	stemmerLanguageMinimalGalician := StemmerLanguageMinimalGalician
	stemmerLanguageLightGerman := StemmerLanguageLightGerman
	stemmerLanguageGerman := StemmerLanguageGerman
	stemmerLanguageGerman2 := StemmerLanguageGerman2
	stemmerLanguageMinimalGerman := StemmerLanguageMinimalGerman
	stemmerLanguageGreek := StemmerLanguageGreek
	stemmerLanguageHindi := StemmerLanguageHindi
	stemmerLanguageHungarian := StemmerLanguageHungarian
	stemmerLanguageLightHungarian := StemmerLanguageLightHungarian
	stemmerLanguageIndonesian := StemmerLanguageIndonesian
	stemmerLanguageIrish := StemmerLanguageIrish
	stemmerLanguageLightItalian := StemmerLanguageLightItalian
	stemmerLanguageItalian := StemmerLanguageItalian
	stemmerLanguageSorani := StemmerLanguageSorani
	stemmerLanguageLatvian := StemmerLanguageLatvian
	stemmerLanguageLithuanian := StemmerLanguageLithuanian
	stemmerLanguageNorwegian := StemmerLanguageNorwegian
	stemmerLanguageLightNorwegian := StemmerLanguageLightNorwegian
	stemmerLanguageMinimalNorwegian := StemmerLanguageMinimalNorwegian
	stemmerLanguageLightNynorsk := StemmerLanguageLightNynorsk
	stemmerLanguageMinimalNynorsk := StemmerLanguageMinimalNynorsk
	stemmerLanguageLightPortuguese := StemmerLanguageLightPortuguese
	stemmerLanguageMinimalPortuguese := StemmerLanguageMinimalPortuguese
	stemmerLanguagePortuguese := StemmerLanguagePortuguese
	stemmerLanguagePortugueseRslp := StemmerLanguagePortugueseRslp
	stemmerLanguageRomanian := StemmerLanguageRomanian
	stemmerLanguageRussian := StemmerLanguageRussian
	stemmerLanguageLightRussian := StemmerLanguageLightRussian
	stemmerLanguageLightSpanish := StemmerLanguageLightSpanish
	stemmerLanguageSpanish := StemmerLanguageSpanish
	stemmerLanguageSwedish := StemmerLanguageSwedish
	stemmerLanguageLightSwedish := StemmerLanguageLightSwedish
	stemmerLanguageTurkish := StemmerLanguageTurkish

	type fields struct {
		name     FilterName
		language *FilterStemmerLanguage
	}
	type args struct {
		language FilterStemmerLanguage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStemmer
	}{
		{
			name: "stemmer_language_arabic",
			args: args{language: StemmerLanguageArabic},
			want: &FilterStemmer{language: &stemmerLanguageArabic},
		},
		{
			name: "stemmer_language_armenian",
			args: args{language: StemmerLanguageArmenian},
			want: &FilterStemmer{language: &stemmerLanguageArmenian},
		},
		{
			name: "stemmer_language_basque",
			args: args{language: StemmerLanguageBasque},
			want: &FilterStemmer{language: &stemmerLanguageBasque},
		},
		{
			name: "stemmer_language_bengali",
			args: args{language: StemmerLanguageBengali},
			want: &FilterStemmer{language: &stemmerLanguageBengali},
		},
		{
			name: "stemmer_language_brazilian",
			args: args{language: StemmerLanguageBrazilian},
			want: &FilterStemmer{language: &stemmerLanguageBrazilian},
		},
		{
			name: "stemmer_language_bulgarian",
			args: args{language: StemmerLanguageBulgarian},
			want: &FilterStemmer{language: &stemmerLanguageBulgarian},
		},
		{
			name: "stemmer_language_catalan",
			args: args{language: StemmerLanguageCatalan},
			want: &FilterStemmer{language: &stemmerLanguageCatalan},
		},
		{
			name: "stemmer_language_czech",
			args: args{language: StemmerLanguageCzech},
			want: &FilterStemmer{language: &stemmerLanguageCzech},
		},
		{
			name: "stemmer_language_danish",
			args: args{language: StemmerLanguageDanish},
			want: &FilterStemmer{language: &stemmerLanguageDanish},
		},
		{
			name: "stemmer_language_dutch",
			args: args{language: StemmerLanguageDutch},
			want: &FilterStemmer{language: &stemmerLanguageDutch},
		},
		{
			name: "stemmer_language_dutch_kp",
			args: args{language: StemmerLanguageDutchKp},
			want: &FilterStemmer{language: &stemmerLanguageDutchKp},
		},
		{
			name: "stemmer_language_english",
			args: args{language: StemmerLanguageEnglish},
			want: &FilterStemmer{language: &stemmerLanguageEnglish},
		},
		{
			name: "stemmer_language_light_english",
			args: args{language: StemmerLanguageLightEnglish},
			want: &FilterStemmer{language: &stemmerLanguageLightEnglish},
		},
		{
			name: "stemmer_language_lovins",
			args: args{language: StemmerLanguageLovins},
			want: &FilterStemmer{language: &stemmerLanguageLovins},
		},
		{
			name: "stemmer_language_minimal_english",
			args: args{language: StemmerLanguageMinimalEnglish},
			want: &FilterStemmer{language: &stemmerLanguageMinimalEnglish},
		},
		{
			name: "stemmer_language_porter2",
			args: args{language: StemmerLanguagePorter2},
			want: &FilterStemmer{language: &stemmerLanguagePorter2},
		},
		{
			name: "stemmer_language_possessive_english",
			args: args{language: StemmerLanguagePossessiveEnglish},
			want: &FilterStemmer{language: &stemmerLanguagePossessiveEnglish},
		},
		{
			name: "stemmer_language_estonian",
			args: args{language: StemmerLanguageEstonian},
			want: &FilterStemmer{language: &stemmerLanguageEstonian},
		},
		{
			name: "stemmer_language_finnish",
			args: args{language: StemmerLanguageFinnish},
			want: &FilterStemmer{language: &stemmerLanguageFinnish},
		},
		{
			name: "stemmer_language_light_finnish",
			args: args{language: StemmerLanguageLightFinnish},
			want: &FilterStemmer{language: &stemmerLanguageLightFinnish},
		},
		{
			name: "stemmer_language_light_french",
			args: args{language: StemmerLanguageLightFrench},
			want: &FilterStemmer{language: &stemmerLanguageLightFrench},
		},
		{
			name: "stemmer_language_french",
			args: args{language: StemmerLanguageFrench},
			want: &FilterStemmer{language: &stemmerLanguageFrench},
		},
		{
			name: "stemmer_language_minimal_french",
			args: args{language: StemmerLanguageMinimalFrench},
			want: &FilterStemmer{language: &stemmerLanguageMinimalFrench},
		},
		{
			name: "stemmer_language_galician",
			args: args{language: StemmerLanguageGalician},
			want: &FilterStemmer{language: &stemmerLanguageGalician},
		},
		{
			name: "stemmer_language_minimal_galician",
			args: args{language: StemmerLanguageMinimalGalician},
			want: &FilterStemmer{language: &stemmerLanguageMinimalGalician},
		},
		{
			name: "stemmer_language_light_german",
			args: args{language: StemmerLanguageLightGerman},
			want: &FilterStemmer{language: &stemmerLanguageLightGerman},
		},
		{
			name: "stemmer_language_german",
			args: args{language: StemmerLanguageGerman},
			want: &FilterStemmer{language: &stemmerLanguageGerman},
		},
		{
			name: "stemmer_language_german2",
			args: args{language: StemmerLanguageGerman2},
			want: &FilterStemmer{language: &stemmerLanguageGerman2},
		},
		{
			name: "stemmer_language_minimal_german",
			args: args{language: StemmerLanguageMinimalGerman},
			want: &FilterStemmer{language: &stemmerLanguageMinimalGerman},
		},
		{
			name: "stemmer_language_greek",
			args: args{language: StemmerLanguageGreek},
			want: &FilterStemmer{language: &stemmerLanguageGreek},
		},
		{
			name: "stemmer_language_hindi",
			args: args{language: StemmerLanguageHindi},
			want: &FilterStemmer{language: &stemmerLanguageHindi},
		},
		{
			name: "stemmer_language_hungarian",
			args: args{language: StemmerLanguageHungarian},
			want: &FilterStemmer{language: &stemmerLanguageHungarian},
		},
		{
			name: "stemmer_language_light_hungarian",
			args: args{language: StemmerLanguageLightHungarian},
			want: &FilterStemmer{language: &stemmerLanguageLightHungarian},
		},
		{
			name: "stemmer_language_indonesian",
			args: args{language: StemmerLanguageIndonesian},
			want: &FilterStemmer{language: &stemmerLanguageIndonesian},
		},
		{
			name: "stemmer_language_irish",
			args: args{language: StemmerLanguageIrish},
			want: &FilterStemmer{language: &stemmerLanguageIrish},
		},
		{
			name: "stemmer_language_light_italian",
			args: args{language: StemmerLanguageLightItalian},
			want: &FilterStemmer{language: &stemmerLanguageLightItalian},
		},
		{
			name: "stemmer_language_italian",
			args: args{language: StemmerLanguageItalian},
			want: &FilterStemmer{language: &stemmerLanguageItalian},
		},
		{
			name: "stemmer_language_sorani",
			args: args{language: StemmerLanguageSorani},
			want: &FilterStemmer{language: &stemmerLanguageSorani},
		},
		{
			name: "stemmer_language_latvian",
			args: args{language: StemmerLanguageLatvian},
			want: &FilterStemmer{language: &stemmerLanguageLatvian},
		},
		{
			name: "stemmer_language_lithuanian",
			args: args{language: StemmerLanguageLithuanian},
			want: &FilterStemmer{language: &stemmerLanguageLithuanian},
		},
		{
			name: "stemmer_language_norwegian",
			args: args{language: StemmerLanguageNorwegian},
			want: &FilterStemmer{language: &stemmerLanguageNorwegian},
		},
		{
			name: "stemmer_language_light_norwegian",
			args: args{language: StemmerLanguageLightNorwegian},
			want: &FilterStemmer{language: &stemmerLanguageLightNorwegian},
		},
		{
			name: "stemmer_language_minimal_norwegian",
			args: args{language: StemmerLanguageMinimalNorwegian},
			want: &FilterStemmer{language: &stemmerLanguageMinimalNorwegian},
		},
		{
			name: "stemmer_language_light_nynorsk",
			args: args{language: StemmerLanguageLightNynorsk},
			want: &FilterStemmer{language: &stemmerLanguageLightNynorsk},
		},
		{
			name: "stemmer_language_minimal_nynorsk",
			args: args{language: StemmerLanguageMinimalNynorsk},
			want: &FilterStemmer{language: &stemmerLanguageMinimalNynorsk},
		},
		{
			name: "stemmer_language_light_portuguese",
			args: args{language: StemmerLanguageLightPortuguese},
			want: &FilterStemmer{language: &stemmerLanguageLightPortuguese},
		},
		{
			name: "stemmer_language_minimal_portuguese",
			args: args{language: StemmerLanguageMinimalPortuguese},
			want: &FilterStemmer{language: &stemmerLanguageMinimalPortuguese},
		},
		{
			name: "stemmer_language_portuguese",
			args: args{language: StemmerLanguagePortuguese},
			want: &FilterStemmer{language: &stemmerLanguagePortuguese},
		},
		{
			name: "stemmer_language_portuguese_rslp",
			args: args{language: StemmerLanguagePortugueseRslp},
			want: &FilterStemmer{language: &stemmerLanguagePortugueseRslp},
		},
		{
			name: "stemmer_language_romanian",
			args: args{language: StemmerLanguageRomanian},
			want: &FilterStemmer{language: &stemmerLanguageRomanian},
		},
		{
			name: "stemmer_language_russian",
			args: args{language: StemmerLanguageRussian},
			want: &FilterStemmer{language: &stemmerLanguageRussian},
		},
		{
			name: "stemmer_language_light_russian",
			args: args{language: StemmerLanguageLightRussian},
			want: &FilterStemmer{language: &stemmerLanguageLightRussian},
		},
		{
			name: "stemmer_language_light_spanish",
			args: args{language: StemmerLanguageLightSpanish},
			want: &FilterStemmer{language: &stemmerLanguageLightSpanish},
		},
		{
			name: "stemmer_language_spanish",
			args: args{language: StemmerLanguageSpanish},
			want: &FilterStemmer{language: &stemmerLanguageSpanish},
		},
		{
			name: "stemmer_language_swedish",
			args: args{language: StemmerLanguageSwedish},
			want: &FilterStemmer{language: &stemmerLanguageSwedish},
		},
		{
			name: "stemmer_language_light_swedish",
			args: args{language: StemmerLanguageLightSwedish},
			want: &FilterStemmer{language: &stemmerLanguageLightSwedish},
		},
		{
			name: "stemmer_language_turkish",
			args: args{language: StemmerLanguageTurkish},
			want: &FilterStemmer{language: &stemmerLanguageTurkish},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmer{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.SetLanguage(tt.args.language); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStemmer_Source(t *testing.T) {
	stemmerLanguageArabic := StemmerLanguageArabic

	type fields struct {
		name     FilterName
		language *FilterStemmerLanguage
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeStemmer,
			},
		},
		{
			name:   "language",
			fields: fields{language: &stemmerLanguageArabic},
			want: map[string]interface{}{
				"type":     FilterTypeStemmer,
				"language": StemmerLanguageArabic,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmer{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			got, err := filter.Source()
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

func TestFilterStemmer_Type(t *testing.T) {
	type fields struct {
		name     FilterName
		language *FilterStemmerLanguage
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeStemmer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStemmer{
				name:     tt.fields.name,
				language: tt.fields.language,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStop_AddStopWords(t *testing.T) {
	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	type args struct {
		stopWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStop
	}{
		{
			name: "set",
			args: args{stopWords: []string{"test"}},
			want: &FilterStop{stopWords: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{stopWords: []string{"test"}},
			args:   args{stopWords: []string{"test1"}},
			want:   &FilterStop{stopWords: []string{"test", "test1"}},
		},
		{
			name:   "stop_words_arabic",
			args:   args{stopWords: []string{StopWordsArabic}},
			want:   &FilterStop{stopWords: []string{"_arabic_"}},
		},
		{
			name:   "stop_words_armenian",
			args:   args{stopWords: []string{StopWordsArmenian}},
			want:   &FilterStop{stopWords: []string{"_armenian_"}},
		},
		{
			name:   "stop_words_turkish",
			args:   args{stopWords: []string{StopWordsTurkish}},
			want:   &FilterStop{stopWords: []string{"_turkish_"}},
		},
		{
			name:   "stop_words_thai",
			args:   args{stopWords: []string{StopWordsThai}},
			want:   &FilterStop{stopWords: []string{"_thai_"}},
		},
		{
			name:   "stop_words_swedish",
			args:   args{stopWords: []string{StopWordsSwedish}},
			want:   &FilterStop{stopWords: []string{"_swedish_"}},
		},
		{
			name:   "stop_words_spanish",
			args:   args{stopWords: []string{StopWordsSpanish}},
			want:   &FilterStop{stopWords: []string{"_spanish_"}},
		},
		{
			name:   "stop_words_sorani",
			args:   args{stopWords: []string{StopWordsSorani}},
			want:   &FilterStop{stopWords: []string{"_sorani_"}},
		},
		{
			name:   "stop_words_russian",
			args:   args{stopWords: []string{StopWordsRussian}},
			want:   &FilterStop{stopWords: []string{"_russian_"}},
		},
		{
			name:   "stop_words_romanian",
			args:   args{stopWords: []string{StopWordsRomanian}},
			want:   &FilterStop{stopWords: []string{"_romanian_"}},
		},
		{
			name:   "stop_words_portuguese",
			args:   args{stopWords: []string{StopWordsPortuguese}},
			want:   &FilterStop{stopWords: []string{"_portuguese_"}},
		},
		{
			name:   "stop_words_persian",
			args:   args{stopWords: []string{StopWordsPersian}},
			want:   &FilterStop{stopWords: []string{"_persian_"}},
		},
		{
			name:   "stop_words_norwegian",
			args:   args{stopWords: []string{StopWordsNorwegian}},
			want:   &FilterStop{stopWords: []string{"_norwegian_"}},
		},
		{
			name:   "stop_words_lithuanian",
			args:   args{stopWords: []string{StopWordsLithuanian}},
			want:   &FilterStop{stopWords: []string{"_lithuanian_"}},
		},
		{
			name:   "stop_words_latvian",
			args:   args{stopWords: []string{StopWordsLatvian}},
			want:   &FilterStop{stopWords: []string{"_latvian_"}},
		},
		{
			name:   "stop_words_italian",
			args:   args{stopWords: []string{StopWordsItalian}},
			want:   &FilterStop{stopWords: []string{"_italian_"}},
		},
		{
			name:   "stop_words_irish",
			args:   args{stopWords: []string{StopWordsIrish}},
			want:   &FilterStop{stopWords: []string{"_irish_"}},
		},
		{
			name:   "stop_words_indonesian",
			args:   args{stopWords: []string{StopWordsIndonesian}},
			want:   &FilterStop{stopWords: []string{"_indonesian_"}},
		},
		{
			name:   "stop_words_hungarian",
			args:   args{stopWords: []string{StopWordsHungarian}},
			want:   &FilterStop{stopWords: []string{"_hungarian_"}},
		},
		{
			name:   "stop_words_hindi",
			args:   args{stopWords: []string{StopWordsHindi}},
			want:   &FilterStop{stopWords: []string{"_hindi_"}},
		},
		{
			name:   "stop_words_greek",
			args:   args{stopWords: []string{StopWordsGreek}},
			want:   &FilterStop{stopWords: []string{"_greek_"}},
		},
		{
			name:   "stop_words_german",
			args:   args{stopWords: []string{StopWordsGerman}},
			want:   &FilterStop{stopWords: []string{"_german_"}},
		},
		{
			name:   "stop_words_galician",
			args:   args{stopWords: []string{StopWordsGalician}},
			want:   &FilterStop{stopWords: []string{"_galician_"}},
		},
		{
			name:   "stop_words_french",
			args:   args{stopWords: []string{StopWordsFrench}},
			want:   &FilterStop{stopWords: []string{"_french_"}},
		},
		{
			name:   "stop_words_finnish",
			args:   args{stopWords: []string{StopWordsFinnish}},
			want:   &FilterStop{stopWords: []string{"_finnish_"}},
		},
		{
			name:   "stop_words_estonian",
			args:   args{stopWords: []string{StopWordsEstonian}},
			want:   &FilterStop{stopWords: []string{"_estonian_"}},
		},
		{
			name:   "stop_words_english",
			args:   args{stopWords: []string{StopWordsEnglish}},
			want:   &FilterStop{stopWords: []string{"_english_"}},
		},
		{
			name:   "stop_words_dutch",
			args:   args{stopWords: []string{StopWordsDutch}},
			want:   &FilterStop{stopWords: []string{"_dutch_"}},
		},
		{
			name:   "stop_words_danish",
			args:   args{stopWords: []string{StopWordsDanish}},
			want:   &FilterStop{stopWords: []string{"_danish_"}},
		},
		{
			name:   "stop_words_czech",
			args:   args{stopWords: []string{StopWordsCzech}},
			want:   &FilterStop{stopWords: []string{"_czech_"}},
		},
		{
			name:   "stop_words_cjk",
			args:   args{stopWords: []string{StopWordsCjk}},
			want:   &FilterStop{stopWords: []string{"_cjk_"}},
		},
		{
			name:   "stop_words_catalan",
			args:   args{stopWords: []string{StopWordsCatalan}},
			want:   &FilterStop{stopWords: []string{"_catalan_"}},
		},
		{
			name:   "stop_words_bulgarian",
			args:   args{stopWords: []string{StopWordsBulgarian}},
			want:   &FilterStop{stopWords: []string{"_bulgarian_"}},
		},
		{
			name:   "stop_words_brazilian",
			args:   args{stopWords: []string{StopWordsBrazilian}},
			want:   &FilterStop{stopWords: []string{"_brazilian_"}},
		},
		{
			name:   "stop_words_bengali",
			args:   args{stopWords: []string{StopWordsBengali}},
			want:   &FilterStop{stopWords: []string{"_bengali_"}},
		},
		{
			name:   "stop_words_basque",
			args:   args{stopWords: []string{StopWordsBasque}},
			want:   &FilterStop{stopWords: []string{"_basque_"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			if got := filter.AddStopWords(tt.args.stopWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStop_Name(t *testing.T) {
	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStop_SetIgnoreCase(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	type args struct {
		ignoreCase bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStop
	}{
		{
			name: "set",
			args: args{ignoreCase: true},
			want: &FilterStop{ignoreCase: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{ignoreCase: &boolTrue},
			args:   args{ignoreCase: false},
			want:   &FilterStop{ignoreCase: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			if got := filter.SetIgnoreCase(tt.args.ignoreCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStop_SetRemoveTrailing(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	type args struct {
		removeTrailing bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStop
	}{
		{
			name: "set",
			args: args{removeTrailing: true},
			want: &FilterStop{removeTrailing: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{removeTrailing: &boolTrue},
			args:   args{removeTrailing: false},
			want:   &FilterStop{removeTrailing: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			if got := filter.SetRemoveTrailing(tt.args.removeTrailing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRemoveTrailing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStop_SetStopWordsPath(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	type args struct {
		stopWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterStop
	}{
		{
			name: "set",
			args: args{stopWordsPath: "test"},
			want: &FilterStop{stopWordsPath: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{stopWordsPath: &testStringSet},
			args:   args{stopWordsPath: "test1"},
			want:   &FilterStop{stopWordsPath: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			if got := filter.SetStopWordsPath(tt.args.stopWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStopWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterStop_Source(t *testing.T) {
	testString := "test"
	boolTrue := true

	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeStop,
			},
		},
		{
			name:   "stop_words_path",
			fields: fields{stopWordsPath: &testString},
			want: map[string]interface{}{
				"type":            FilterTypeStop,
				"stop_words_path": "test",
			},
		},
		{
			name:   "ignore_case",
			fields: fields{ignoreCase: &boolTrue},
			want: map[string]interface{}{
				"type":        FilterTypeStop,
				"ignore_case": true,
			},
		},
		{
			name:   "remove_trailing",
			fields: fields{removeTrailing: &boolTrue},
			want: map[string]interface{}{
				"type":            FilterTypeStop,
				"remove_trailing": true,
			},
		},
		{
			name:   "stop_words",
			fields: fields{stopWords: []string{"text"}},
			want: map[string]interface{}{
				"type":       FilterTypeStop,
				"stop_words": []string{"text"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			got, err := filter.Source()
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

func TestFilterStop_Type(t *testing.T) {
	type fields struct {
		name           FilterName
		stopWordsPath  *string
		ignoreCase     *bool
		removeTrailing *bool
		stopWords      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeStop,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterStop{
				name:           tt.fields.name,
				stopWordsPath:  tt.fields.stopWordsPath,
				ignoreCase:     tt.fields.ignoreCase,
				removeTrailing: tt.fields.removeTrailing,
				stopWords:      tt.fields.stopWords,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_AddSynonyms(t *testing.T) {
	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	type args struct {
		synonyms []Synonym
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterSynonym
	}{
		{
			name: "set",
			args: args{synonyms: []Synonym{&SynonymSolr{}}},
			want: &FilterSynonym{synonyms: []Synonym{&SynonymSolr{}}},
		},
		{
			name:   "add",
			fields: fields{synonyms: []Synonym{&SynonymSolr{synonyms: []string{"test"}}}},
			args:   args{synonyms: []Synonym{&SynonymSolr{}}},
			want:   &FilterSynonym{synonyms: []Synonym{&SynonymSolr{synonyms: []string{"test"}}, &SynonymSolr{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.AddSynonyms(tt.args.synonyms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSynonyms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_Name(t *testing.T) {
	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_SetExpand(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	type args struct {
		expand bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterSynonym
	}{
		{
			name: "set",
			args: args{expand: true},
			want: &FilterSynonym{expand: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{expand: &boolTrue},
			args:   args{expand: false},
			want:   &FilterSynonym{expand: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.SetExpand(tt.args.expand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExpand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_SetFormat(t *testing.T) {
	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	type args struct {
		format FilterSynonymFormat
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterSynonym
	}{
		{
			name: "synonym_format_wordnet",
			args: args{format: SynonymFormatWordnet},
			want: &FilterSynonym{format: SynonymFormatWordnet},
		},
		{
			name: "synonym_format_solr",
			args: args{format: SynonymFormatSolr},
			want: &FilterSynonym{format: SynonymFormatSolr},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.SetFormat(tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_SetLenient(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	type args struct {
		lenient bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterSynonym
	}{
		{
			name: "set",
			args: args{lenient: true},
			want: &FilterSynonym{lenient: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{lenient: &boolTrue},
			args:   args{lenient: false},
			want:   &FilterSynonym{lenient: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.SetLenient(tt.args.lenient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLenient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_SetSynonymsPath(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	type args struct {
		synonymsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterSynonym
	}{
		{
			name: "set",
			args: args{synonymsPath: "test"},
			want: &FilterSynonym{synonymsPath: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{synonymsPath: &testStringSet},
			args:   args{synonymsPath: "test1"},
			want:   &FilterSynonym{synonymsPath: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.SetSynonymsPath(tt.args.synonymsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSynonymsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterSynonym_Source(t *testing.T) {
	testString := "test"
	boolTrue := true

	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterType(""),
			},
		},
		{
			name:   "filter_type_synonym",
			fields: fields{_type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type": FilterTypeSynonym,
			},
		},
		{
			name:   "filter_filter_type_synonym_graph",
			fields: fields{_type: FilterTypeSynonymGraph},
			want: map[string]interface{}{
				"type": FilterTypeSynonymGraph,
			},
		},
		{
			name:   "format_synonym_format_wordnet",
			fields: fields{format: SynonymFormatWordnet, _type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type":   FilterTypeSynonym,
				"format": SynonymFormatWordnet,
			},
		},
		{
			name:   "format_synonym_format_solr",
			fields: fields{format: SynonymFormatSolr, _type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type": FilterTypeSynonym,
			},
		},
		{
			name:   "synonyms_path",
			fields: fields{synonymsPath: &testString, _type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type":          FilterTypeSynonym,
				"synonyms_path": "test",
			},
		},
		{
			name:   "expand",
			fields: fields{expand: &boolTrue, _type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type":   FilterTypeSynonym,
				"expand": true,
			},
		},
		{
			name:   "lenient",
			fields: fields{lenient: &boolTrue, _type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type":    FilterTypeSynonym,
				"lenient": true,
			},
		},
		{
			name:   "synonyms",
			fields: fields{synonyms: []Synonym{&SynonymSolr{synonyms: []string{"test"}, words: []string{"test"}}}, _type: FilterTypeSynonym},
			want: map[string]interface{}{
				"type":     FilterTypeSynonym,
				"synonyms": []string{(&SynonymSolr{synonyms: []string{"test"}, words: []string{"test"}}).Synonym()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			got, err := filter.Source()
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

func TestFilterSynonym_Type(t *testing.T) {
	type fields struct {
		_type        FilterType
		name         FilterName
		format       FilterSynonymFormat
		synonymsPath *string
		expand       *bool
		lenient      *bool
		synonyms     []Synonym
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			name:   "filter_type_synonym",
			fields: fields{_type: FilterTypeSynonym},
			want:   FilterTypeSynonym,
		},
		{
			name:   "filter_type_synonym_graph",
			fields: fields{_type: FilterTypeSynonymGraph},
			want:   FilterTypeSynonymGraph,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterSynonym{
				_type:        tt.fields._type,
				name:         tt.fields.name,
				format:       tt.fields.format,
				synonymsPath: tt.fields.synonymsPath,
				expand:       tt.fields.expand,
				lenient:      tt.fields.lenient,
				synonyms:     tt.fields.synonyms,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTruncate_Name(t *testing.T) {
	type fields struct {
		name   FilterName
		length *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterTruncate{
				name:   tt.fields.name,
				length: tt.fields.length,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTruncate_SetLength(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name   FilterName
		length *uint32
	}
	type args struct {
		length uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterTruncate
	}{
		{
			name: "set",
			args: args{length: 1},
			want: &FilterTruncate{length: &numberSet},
		},
		{
			name:   "change",
			fields: fields{length: &numberSet},
			args:   args{length: 2},
			want:   &FilterTruncate{length: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterTruncate{
				name:   tt.fields.name,
				length: tt.fields.length,
			}
			if got := filter.SetLength(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTruncate_Source(t *testing.T) {
	number := uint32(1)

	type fields struct {
		name   FilterName
		length *uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeTruncate,
			},
		},
		{
			name:   "length",
			fields: fields{length: &number},
			want: map[string]interface{}{
				"type":   FilterTypeTruncate,
				"length": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterTruncate{
				name:   tt.fields.name,
				length: tt.fields.length,
			}
			got, err := filter.Source()
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

func TestFilterTruncate_Type(t *testing.T) {
	type fields struct {
		name   FilterName
		length *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeTruncate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterTruncate{
				name:   tt.fields.name,
				length: tt.fields.length,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUnique_Name(t *testing.T) {
	type fields struct {
		name               FilterName
		onlyOnSamePosition *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterUnique{
				name:               tt.fields.name,
				onlyOnSamePosition: tt.fields.onlyOnSamePosition,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUnique_SetOnlyOnSamePosition(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name               FilterName
		onlyOnSamePosition *bool
	}
	type args struct {
		onlyOnSamePosition bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterUnique
	}{
		{
			name: "set",
			args: args{onlyOnSamePosition: true},
			want: &FilterUnique{onlyOnSamePosition: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{onlyOnSamePosition: &boolTrue},
			args:   args{onlyOnSamePosition: false},
			want:   &FilterUnique{onlyOnSamePosition: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterUnique{
				name:               tt.fields.name,
				onlyOnSamePosition: tt.fields.onlyOnSamePosition,
			}
			if got := filter.SetOnlyOnSamePosition(tt.args.onlyOnSamePosition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlyOnSamePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterUnique_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		name               FilterName
		onlyOnSamePosition *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeUnique,
			},
		},
		{
			name:   "only_on_same_position",
			fields: fields{onlyOnSamePosition: &boolTrue},
			want: map[string]interface{}{
				"type":                  FilterTypeUnique,
				"only_on_same_position": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterUnique{
				name:               tt.fields.name,
				onlyOnSamePosition: tt.fields.onlyOnSamePosition,
			}
			got, err := filter.Source()
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

func TestFilterUnique_Type(t *testing.T) {
	type fields struct {
		name               FilterName
		onlyOnSamePosition *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeUnique,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterUnique{
				name:               tt.fields.name,
				onlyOnSamePosition: tt.fields.onlyOnSamePosition,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiterGraph_SetAdjustOffsets(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		FilterWordDelimiter FilterWordDelimiter
		adjustOffsets       *bool
		ignoreKeywords      *bool
	}
	type args struct {
		adjustOffsets bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiterGraph
	}{
		{
			name: "set",
			args: args{adjustOffsets: true},
			want: &FilterWordDelimiterGraph{adjustOffsets: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{adjustOffsets: &boolTrue},
			args:   args{adjustOffsets: false},
			want:   &FilterWordDelimiterGraph{adjustOffsets: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiterGraph{
				FilterWordDelimiter: tt.fields.FilterWordDelimiter,
				adjustOffsets:       tt.fields.adjustOffsets,
				ignoreKeywords:      tt.fields.ignoreKeywords,
			}
			if got := filter.SetAdjustOffsets(tt.args.adjustOffsets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAdjustOffsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiterGraph_SetIgnoreKeywords(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		FilterWordDelimiter FilterWordDelimiter
		adjustOffsets       *bool
		ignoreKeywords      *bool
	}
	type args struct {
		ignoreKeywords bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiterGraph
	}{
		{
			name: "set",
			args: args{ignoreKeywords: true},
			want: &FilterWordDelimiterGraph{ignoreKeywords: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{ignoreKeywords: &boolTrue},
			args:   args{ignoreKeywords: false},
			want:   &FilterWordDelimiterGraph{ignoreKeywords: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiterGraph{
				FilterWordDelimiter: tt.fields.FilterWordDelimiter,
				adjustOffsets:       tt.fields.adjustOffsets,
				ignoreKeywords:      tt.fields.ignoreKeywords,
			}
			if got := filter.SetIgnoreKeywords(tt.args.ignoreKeywords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIgnoreKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiterGraph_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		FilterWordDelimiter FilterWordDelimiter
		adjustOffsets       *bool
		ignoreKeywords      *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeWordDelimiterGraph,
			},
		},
		{
			name:   "adjust_offsets",
			fields: fields{adjustOffsets: &boolTrue},
			want: map[string]interface{}{
				"type":           FilterTypeWordDelimiterGraph,
				"adjust_offsets": true,
			},
		},
		{
			name:   "ignore_keywords",
			fields: fields{ignoreKeywords: &boolTrue},
			want: map[string]interface{}{
				"type":            FilterTypeWordDelimiterGraph,
				"ignore_keywords": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiterGraph{
				FilterWordDelimiter: tt.fields.FilterWordDelimiter,
				adjustOffsets:       tt.fields.adjustOffsets,
				ignoreKeywords:      tt.fields.ignoreKeywords,
			}
			got, err := filter.Source()
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

func TestFilterWordDelimiterGraph_Type(t *testing.T) {
	type fields struct {
		FilterWordDelimiter FilterWordDelimiter
		adjustOffsets       *bool
		ignoreKeywords      *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeWordDelimiterGraph,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiterGraph{
				FilterWordDelimiter: tt.fields.FilterWordDelimiter,
				adjustOffsets:       tt.fields.adjustOffsets,
				ignoreKeywords:      tt.fields.ignoreKeywords,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_AddProtectedWords(t *testing.T) {
	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		protectedWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{protectedWords: []string{"test"}},
			want: &FilterWordDelimiter{protectedWords: []string{"test"}},
		},
		{
			name:   "add",
			fields: fields{protectedWords: []string{"test"}},
			args:   args{protectedWords: []string{"test1"}},
			want:   &FilterWordDelimiter{protectedWords: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.AddProtectedWords(tt.args.protectedWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddProtectedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_AddTypeTable(t *testing.T) {
	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		typeTable []*WordDelimiterType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{typeTable: []*WordDelimiterType{{}}},
			want: &FilterWordDelimiter{typeTable: []*WordDelimiterType{{}}},
		},
		{
			name:   "add",
			fields: fields{typeTable: []*WordDelimiterType{{}}},
			args:   args{typeTable: []*WordDelimiterType{{_type: WordDelimiterAlpha}}},
			want:   &FilterWordDelimiter{typeTable: []*WordDelimiterType{{}, {_type: WordDelimiterAlpha}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.AddTypeTable(tt.args.typeTable...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTypeTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_Name(t *testing.T) {
	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterName
	}{
		{
			fields: fields{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetCatenateAll(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		catenateAll bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{catenateAll: true},
			want: &FilterWordDelimiter{catenateAll: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{catenateAll: &boolTrue},
			args:   args{catenateAll: false},
			want:   &FilterWordDelimiter{catenateAll: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetCatenateAll(tt.args.catenateAll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCatenateAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetCatenateNumbers(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		catenateNumbers bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{catenateNumbers: true},
			want: &FilterWordDelimiter{catenateNumbers: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{catenateNumbers: &boolTrue},
			args:   args{catenateNumbers: false},
			want:   &FilterWordDelimiter{catenateNumbers: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetCatenateNumbers(tt.args.catenateNumbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCatenateNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetCatenateWords(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		catenateWords bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{catenateWords: true},
			want: &FilterWordDelimiter{catenateWords: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{catenateWords: &boolTrue},
			args:   args{catenateWords: false},
			want:   &FilterWordDelimiter{catenateWords: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetCatenateWords(tt.args.catenateWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCatenateWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetGenerateNumberParts(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		generateNumberParts bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{generateNumberParts: true},
			want: &FilterWordDelimiter{generateNumberParts: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{generateNumberParts: &boolTrue},
			args:   args{generateNumberParts: false},
			want:   &FilterWordDelimiter{generateNumberParts: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetGenerateNumberParts(tt.args.generateNumberParts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetGenerateNumberParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetGenerateWordParts(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		generateWordParts bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{generateWordParts: true},
			want: &FilterWordDelimiter{generateWordParts: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{generateWordParts: &boolTrue},
			args:   args{generateWordParts: false},
			want:   &FilterWordDelimiter{generateWordParts: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetGenerateWordParts(tt.args.generateWordParts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetGenerateWordParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetPreserveOriginal(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		preserveOriginal bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{preserveOriginal: true},
			want: &FilterWordDelimiter{preserveOriginal: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{preserveOriginal: &boolTrue},
			args:   args{preserveOriginal: false},
			want:   &FilterWordDelimiter{preserveOriginal: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetPreserveOriginal(tt.args.preserveOriginal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPreserveOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetProtectedWordsPath(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		protectedWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{protectedWordsPath: "test"},
			want: &FilterWordDelimiter{protectedWordsPath: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{protectedWordsPath: &testStringSet},
			args:   args{protectedWordsPath: "test1"},
			want:   &FilterWordDelimiter{protectedWordsPath: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetProtectedWordsPath(tt.args.protectedWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProtectedWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetSplitOnCaseChange(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		splitOnCaseChange bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{splitOnCaseChange: true},
			want: &FilterWordDelimiter{splitOnCaseChange: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{splitOnCaseChange: &boolTrue},
			args:   args{splitOnCaseChange: false},
			want:   &FilterWordDelimiter{splitOnCaseChange: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetSplitOnCaseChange(tt.args.splitOnCaseChange); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSplitOnCaseChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetSplitOnNumerics(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		splitOnNumerics bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{splitOnNumerics: true},
			want: &FilterWordDelimiter{splitOnNumerics: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{splitOnNumerics: &boolTrue},
			args:   args{splitOnNumerics: false},
			want:   &FilterWordDelimiter{splitOnNumerics: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetSplitOnNumerics(tt.args.splitOnNumerics); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSplitOnNumerics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetStemEnglishPossessive(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		stemEnglishPossessive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{stemEnglishPossessive: true},
			want: &FilterWordDelimiter{stemEnglishPossessive: &boolTrue},
		},
		{
			name:   "change",
			fields: fields{stemEnglishPossessive: &boolTrue},
			args:   args{stemEnglishPossessive: false},
			want:   &FilterWordDelimiter{stemEnglishPossessive: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetStemEnglishPossessive(tt.args.stemEnglishPossessive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStemEnglishPossessive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_SetTypeTablePath(t *testing.T) {
	testStringSet := "test"
	testStringChange := "test1"

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	type args struct {
		typeTablePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterWordDelimiter
	}{
		{
			name: "set",
			args: args{typeTablePath: "test"},
			want: &FilterWordDelimiter{typeTablePath: &testStringSet},
		},
		{
			name:   "change",
			fields: fields{typeTablePath: &testStringSet},
			args:   args{typeTablePath: "test1"},
			want:   &FilterWordDelimiter{typeTablePath: &testStringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.SetTypeTablePath(tt.args.typeTablePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTypeTablePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWordDelimiter_Source(t *testing.T) {
	testString := "test"
	boolTrue := true

	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{
				"type": FilterTypeWordDelimiter,
			},
		},
		{
			name:   "catenate_all",
			fields: fields{catenateAll: &boolTrue},
			want: map[string]interface{}{
				"type":         FilterTypeWordDelimiter,
				"catenate_all": true,
			},
		},
		{
			name:   "catenate_numbers",
			fields: fields{catenateNumbers: &boolTrue},
			want: map[string]interface{}{
				"type":             FilterTypeWordDelimiter,
				"catenate_numbers": true,
			},
		},
		{
			name:   "catenate_words",
			fields: fields{catenateWords: &boolTrue},
			want: map[string]interface{}{
				"type":           FilterTypeWordDelimiter,
				"catenate_words": true,
			},
		},
		{
			name:   "generate_number_parts",
			fields: fields{generateNumberParts: &boolTrue},
			want: map[string]interface{}{
				"type":                  FilterTypeWordDelimiter,
				"generate_number_parts": true,
			},
		},
		{
			name:   "generate_word_parts",
			fields: fields{generateWordParts: &boolTrue},
			want: map[string]interface{}{
				"type":                FilterTypeWordDelimiter,
				"generate_word_parts": true,
			},
		},
		{
			name:   "preserve_original",
			fields: fields{preserveOriginal: &boolTrue},
			want: map[string]interface{}{
				"type":              FilterTypeWordDelimiter,
				"preserve_original": true,
			},
		},
		{
			name:   "split_on_case_change",
			fields: fields{splitOnCaseChange: &boolTrue},
			want: map[string]interface{}{
				"type":                 FilterTypeWordDelimiter,
				"split_on_case_change": true,
			},
		},
		{
			name:   "split_on_numerics",
			fields: fields{splitOnNumerics: &boolTrue},
			want: map[string]interface{}{
				"type":              FilterTypeWordDelimiter,
				"split_on_numerics": true,
			},
		},
		{
			name:   "stem_english_possessive",
			fields: fields{stemEnglishPossessive: &boolTrue},
			want: map[string]interface{}{
				"type":                    FilterTypeWordDelimiter,
				"stem_english_possessive": true,
			},
		},
		{
			name:   "protected_words_path",
			fields: fields{protectedWordsPath: &testString},
			want: map[string]interface{}{
				"type":                 FilterTypeWordDelimiter,
				"protected_words_path": "test",
			},
		},
		{
			name:   "type_table_path",
			fields: fields{typeTablePath: &testString},
			want: map[string]interface{}{
				"type":            FilterTypeWordDelimiter,
				"type_table_path": "test",
			},
		},
		{
			name:   "protected_words",
			fields: fields{protectedWords: []string{"test"}},
			want: map[string]interface{}{
				"type":            FilterTypeWordDelimiter,
				"protected_words": []string{"test"},
			},
		},
		{
			name:   "type_table",
			fields: fields{typeTable: []*WordDelimiterType{{word: "test", _type: WordDelimiterAlpha}}},
			want: map[string]interface{}{
				"type": FilterTypeWordDelimiter,
				"type_table": []string{
					(&WordDelimiterType{word: "test", _type: WordDelimiterAlpha}).String(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			got, err := filter.Source()
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

func TestFilterWordDelimiter_Type(t *testing.T) {
	type fields struct {
		name                  FilterName
		catenateAll           *bool
		catenateNumbers       *bool
		catenateWords         *bool
		generateNumberParts   *bool
		generateWordParts     *bool
		preserveOriginal      *bool
		splitOnCaseChange     *bool
		splitOnNumerics       *bool
		stemEnglishPossessive *bool
		protectedWordsPath    *string
		typeTablePath         *string
		protectedWords        []string
		typeTable             []*WordDelimiterType
	}
	tests := []struct {
		name   string
		fields fields
		want   FilterType
	}{
		{
			want: FilterTypeWordDelimiter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &FilterWordDelimiter{
				name:                  tt.fields.name,
				catenateAll:           tt.fields.catenateAll,
				catenateNumbers:       tt.fields.catenateNumbers,
				catenateWords:         tt.fields.catenateWords,
				generateNumberParts:   tt.fields.generateNumberParts,
				generateWordParts:     tt.fields.generateWordParts,
				preserveOriginal:      tt.fields.preserveOriginal,
				splitOnCaseChange:     tt.fields.splitOnCaseChange,
				splitOnNumerics:       tt.fields.splitOnNumerics,
				stemEnglishPossessive: tt.fields.stemEnglishPossessive,
				protectedWordsPath:    tt.fields.protectedWordsPath,
				typeTablePath:         tt.fields.typeTablePath,
				protectedWords:        tt.fields.protectedWords,
				typeTable:             tt.fields.typeTable,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterAsciiFolding(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterAsciiFolding
	}{
		{
			args: args{name: "test"},
			want: &FilterAsciiFolding{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterAsciiFolding(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterAsciiFolding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterCjkBigRam(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterCjkBigRam
	}{
		{
			args: args{name: "test"},
			want: &FilterCjkBigRam{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterCjkBigRam(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterCjkBigRam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterCommonGrams(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterCommonGrams
	}{
		{
			args: args{name: "test"},
			want: &FilterCommonGrams{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterCommonGrams(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterCommonGrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterCondition(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterCondition
	}{
		{
			args: args{name: "test"},
			want: &FilterCondition{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterCondition(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterDelimitedPayload(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterDelimitedPayload
	}{
		{
			args: args{name: "test"},
			want: &FilterDelimitedPayload{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterDelimitedPayload(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterDelimitedPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterDictionaryDecompounder(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterDictionaryDecompounder
	}{
		{
			args: args{name: "test"},
			want: &FilterDictionaryDecompounder{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterDictionaryDecompounder(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterDictionaryDecompounder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterEdgeNgram(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterEdgeNgram
	}{
		{
			args: args{name: "test"},
			want: &FilterEdgeNgram{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterEdgeNgram(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterEdgeNgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterElision(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterElision
	}{
		{
			args: args{name: "test"},
			want: &FilterElision{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterElision(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterElision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterFingerprint(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterFingerprint
	}{
		{
			args: args{name: "test"},
			want: &FilterFingerprint{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterFingerprint(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterFingerprint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterHyphenationDecompounder(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterHyphenationDecompounder
	}{
		{
			args: args{name: "test"},
			want: &FilterHyphenationDecompounder{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterHyphenationDecompounder(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterHyphenationDecompounder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterKeep(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterKeep
	}{
		{
			args: args{name: "test"},
			want: &FilterKeep{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterKeep(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterKeep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterKeepTypes(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterKeepTypes
	}{
		{
			args: args{name: "test"},
			want: &FilterKeepTypes{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterKeepTypes(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterKeepTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterKeywordMarker(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterKeywordMarker
	}{
		{
			args: args{name: "test"},
			want: &FilterKeywordMarker{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterKeywordMarker(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterKeywordMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterLength(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterLength
	}{
		{
			args: args{name: "test"},
			want: &FilterLength{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterLength(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterLimit(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterLimit
	}{
		{
			args: args{name: "test"},
			want: &FilterLimit{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterLimit(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterLowercase(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterLowercase
	}{
		{
			args: args{name: "test"},
			want: &FilterLowercase{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterLowercase(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterLowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterMinHash(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterMinHash
	}{
		{
			args: args{name: "test"},
			want: &FilterMinHash{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterMinHash(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterMinHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterMultiplexer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterMultiplexer
	}{
		{
			args: args{name: "test"},
			want: &FilterMultiplexer{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterMultiplexer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterMultiplexer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterNgram(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterNgram
	}{
		{
			args: args{name: "test"},
			want: &FilterNgram{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterNgram(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterNgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterPatternCapture(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterPatternCapture
	}{
		{
			args: args{name: "test"},
			want: &FilterPatternCapture{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterPatternCapture(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterPatternCapture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterPatternReplace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterPatternReplace
	}{
		{
			args: args{name: "test"},
			want: &FilterPatternReplace{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterPatternReplace(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterPatternReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterPhonetic(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterPhonetic
	}{
		{
			args: args{name: "test"},
			want: &FilterPhonetic{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterPhonetic(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterPhonetic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterPredicateTokenFilter(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterPredicateTokenFilter
	}{
		{
			args: args{name: "test"},
			want: &FilterPredicateTokenFilter{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterPredicateTokenFilter(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterPredicateTokenFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterShingle(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterShingle
	}{
		{
			args: args{name: "test"},
			want: &FilterShingle{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterShingle(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterShingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterSnowball(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterSnowball
	}{
		{
			args: args{name: "test"},
			want: &FilterSnowball{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterSnowball(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterSnowball() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterStemmer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterStemmer
	}{
		{
			args: args{name: "test"},
			want: &FilterStemmer{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterStemmer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterStemmer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterStemmerOverride(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterStemmerOverride
	}{
		{
			args: args{name: "test"},
			want: &FilterStemmerOverride{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterStemmerOverride(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterStemmerOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterStop(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterStop
	}{
		{
			args: args{name: "test"},
			want: &FilterStop{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterStop(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterStop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterSynonym(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterSynonym
	}{
		{
			args: args{name: "test"},
			want: &FilterSynonym{name: "test", _type: FilterTypeSynonym},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterSynonym(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterSynonym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterSynonymGraph(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterSynonym
	}{
		{
			args: args{name: "test"},
			want: &FilterSynonym{name: "test", _type: FilterTypeSynonymGraph},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterSynonymGraph(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterSynonymGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterTruncate(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterTruncate
	}{
		{
			args: args{name: "test"},
			want: &FilterTruncate{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterTruncate(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterTruncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterUnique(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterUnique
	}{
		{
			args: args{name: "test"},
			want: &FilterUnique{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterUnique(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterWordDelimiter(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterWordDelimiter
	}{
		{
			args: args{name: "test"},
			want: &FilterWordDelimiter{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterWordDelimiter(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterWordDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilterWordDelimiterGraph(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *FilterWordDelimiterGraph
	}{
		{
			args: args{name: "test"},
			want: &FilterWordDelimiterGraph{FilterWordDelimiter: FilterWordDelimiter{name: "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilterWordDelimiterGraph(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterWordDelimiterGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStemmerOverrideRule(t *testing.T) {
	type args struct {
		stemmer string
		worlds  []string
	}
	tests := []struct {
		name string
		args args
		want *StemmerOverrideRule
	}{
		{
			args: args{stemmer: "|", worlds: []string{"test"}},
			want: &StemmerOverrideRule{stemmer: "|", worlds: []string{"test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStemmerOverrideRule(tt.args.stemmer, tt.args.worlds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStemmerOverrideRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSynonymSolr(t *testing.T) {
	tests := []struct {
		name string
		want *SynonymSolr
	}{
		{
			want: &SynonymSolr{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSynonymSolr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSynonymSolr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSynonymWordnet(t *testing.T) {
	tests := []struct {
		name string
		want *SynonymWordnet
	}{
		{
			want: &SynonymWordnet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSynonymWordnet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSynonymWordnet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWordDelimiterType(t *testing.T) {
	type args struct {
		word  string
		_type WordDelimiter
	}
	tests := []struct {
		name string
		args args
		want *WordDelimiterType
	}{
		{
			name: "word_delimiter_alpha",
			args: args{_type: WordDelimiterAlpha, word: "test"},
			want: &WordDelimiterType{_type: WordDelimiterAlpha, word: "test"},
		},
		{
			name: "word_delimiter_alphanum",
			args: args{_type: WordDelimiterAlphanum, word: "test"},
			want: &WordDelimiterType{_type: WordDelimiterAlphanum, word: "test"},
		},
		{
			name: "word_delimiter_digit",
			args: args{_type: WordDelimiterDigit, word: "test"},
			want: &WordDelimiterType{_type: WordDelimiterDigit, word: "test"},
		},
		{
			name: "word_delimiter_lower",
			args: args{_type: WordDelimiterLower, word: "test"},
			want: &WordDelimiterType{_type: WordDelimiterLower, word: "test"},
		},
		{
			name: "word_delimiter_subword_delim",
			args: args{_type: WordDelimiterSubwordDelim, word: "test"},
			want: &WordDelimiterType{_type: WordDelimiterSubwordDelim, word: "test"},
		},
		{
			name: "word_delimiter_upper",
			args: args{_type: WordDelimiterUpper, word: "test"},
			want: &WordDelimiterType{_type: WordDelimiterUpper, word: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWordDelimiterType(tt.args.word, tt.args._type); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWordDelimiterType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStemmerOverrideRule_String(t *testing.T) {
	type fields struct {
		stemmer string
		worlds  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{stemmer: "|", worlds: []string{"test", "test1"}},
			want: "test,test1 => |",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := &StemmerOverrideRule{
				stemmer: tt.fields.stemmer,
				worlds:  tt.fields.worlds,
			}
			if got := rule.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymSolr_AddSynonyms(t *testing.T) {
	type fields struct {
		synonyms []string
		words    []string
	}
	type args struct {
		synonyms []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymSolr
	}{
		{
			name: "set",
			args: args{synonyms: []string{"test"}},
			want: &SynonymSolr{synonyms: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{synonyms: []string{"test"}},
			args:   args{synonyms: []string{"test1"}},
			want:   &SynonymSolr{synonyms: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymSolr{
				synonyms: tt.fields.synonyms,
				words:    tt.fields.words,
			}
			if got := synonym.AddSynonyms(tt.args.synonyms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSynonyms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymSolr_AddWords(t *testing.T) {
	type fields struct {
		synonyms []string
		words    []string
	}
	type args struct {
		words []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymSolr
	}{
		{
			name: "set",
			args: args{words: []string{"test"}},
			want: &SynonymSolr{words: []string{"test"}},
		},
		{
			name:   "change",
			fields: fields{words: []string{"test"}},
			args:   args{words: []string{"test1"}},
			want:   &SynonymSolr{words: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymSolr{
				synonyms: tt.fields.synonyms,
				words:    tt.fields.words,
			}
			if got := synonym.AddWords(tt.args.words...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymSolr_Synonym(t *testing.T) {
	type fields struct {
		synonyms []string
		words    []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{synonyms: []string{"test"}, words: []string{"test1"}},
			want: "test1 => test",
		},
		{
			fields: fields{synonyms: []string{"test"}, words: []string{"test1", "test2"}},
			want: "test1,test2 => test",
		},
		{
			fields: fields{synonyms: []string{"test", "test3"}, words: []string{"test1", "test2"}},
			want: "test1,test2 => test,test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymSolr{
				synonyms: tt.fields.synonyms,
				words:    tt.fields.words,
			}
			if got := synonym.Synonym(); got != tt.want {
				t.Errorf("Synonym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_SetSenseNumber(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	type args struct {
		senseNumber uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymWordnet
	}{
		{
			name: "set",
			args: args{senseNumber: 1},
			want: &SynonymWordnet{senseNumber: 1},
		},
		{
			name:   "change",
			fields: fields{senseNumber: 1},
			args:   args{senseNumber: 2},
			want:   &SynonymWordnet{senseNumber: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.SetSenseNumber(tt.args.senseNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSenseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_SetSynsetId(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	type args struct {
		synsetId uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymWordnet
	}{
		{
			name: "set",
			args: args{synsetId: 1},
			want: &SynonymWordnet{synsetId: 1},
		},
		{
			name: "change",
			fields: fields{synsetId: 1},
			args: args{synsetId: 2},
			want: &SynonymWordnet{synsetId: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.SetSynsetId(tt.args.synsetId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSynsetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_SetSynsetType(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	type args struct {
		synsetType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymWordnet
	}{
		{
			name: "set",
			args: args{synsetType: "test"},
			want: &SynonymWordnet{synsetType: "test"},
		},
		{
			name: "change",
			fields: fields{synsetType: "test"},
			args: args{synsetType: "test1"},
			want: &SynonymWordnet{synsetType: "test1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.SetSynsetType(tt.args.synsetType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSynsetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_SetTagCount(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	type args struct {
		tagCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymWordnet
	}{
		{
			name: "set",
			args: args{tagCount: 1},
			want: &SynonymWordnet{tagCount: 1},
		},
		{
			name: "change",
			fields: fields{tagCount: 1},
			args: args{tagCount: 2},
			want: &SynonymWordnet{tagCount: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.SetTagCount(tt.args.tagCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTagCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_SetWord(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymWordnet
	}{
		{
			name: "set",
			args: args{word: "test"},
			want: &SynonymWordnet{word: "test"},
		},
		{
			name: "change",
			fields: fields{word: "test"},
			args: args{word: "test1"},
			want: &SynonymWordnet{word: "test1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.SetWord(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_SetWordNumber(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	type args struct {
		wordNumber uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SynonymWordnet
	}{
		{
			name: "set",
			args: args{wordNumber: 1},
			want: &SynonymWordnet{wordNumber: 1},
		},
		{
			name: "change",
			fields: fields{wordNumber: 1},
			args: args{wordNumber: 2},
			want: &SynonymWordnet{wordNumber: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.SetWordNumber(tt.args.wordNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWordNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynonymWordnet_Synonym(t *testing.T) {
	type fields struct {
		synsetId    uint32
		wordNumber  uint32
		word        string
		synsetType  string
		senseNumber uint32
		tagCount    uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			want: "s(100000001,3,'desist',v,1,0).",
			fields: fields{
				synsetId:    100000001,
				wordNumber:  3,
				word:        "desist",
				synsetType:  "v",
				senseNumber: 1,
				tagCount:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synonym := &SynonymWordnet{
				synsetId:    tt.fields.synsetId,
				wordNumber:  tt.fields.wordNumber,
				word:        tt.fields.word,
				synsetType:  tt.fields.synsetType,
				senseNumber: tt.fields.senseNumber,
				tagCount:    tt.fields.tagCount,
			}
			if got := synonym.Synonym(); got != tt.want {
				t.Errorf("Synonym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordDelimiterType_String(t *testing.T) {
	type fields struct {
		word  string
		_type WordDelimiter
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{word: "test", _type: WordDelimiterUpper},
			want: "test => UPPER",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wordDelimiter := &WordDelimiterType{
				word:  tt.fields.word,
				_type: tt.fields._type,
			}
			if got := wordDelimiter.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
