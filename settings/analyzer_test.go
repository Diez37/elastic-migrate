package settings

import (
	"reflect"
	"testing"
)

func TestAnalyzerFingerprint_AddStopWords(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
	}
	type args struct {
		stopWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerFingerprint
	}{
		{
			name: "set",
			args: args{stopWords: []string{"test"}},
			want: &AnalyzerFingerprint{stopWords: []string{"test"}},
		},
		{
			name: "add",
			fields: fields{stopWords: []string{"test"}},
			args: args{stopWords: []string{"test2"}},
			want: &AnalyzerFingerprint{stopWords: []string{"test", "test2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.AddStopWords(tt.args.stopWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerFingerprint_Name(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerFingerprint_SetMaxOutputSize(t *testing.T) {
	maxOutputSizeSet := uint8(1)
	maxOutputSizeChange := uint8(2)

	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
	}
	type args struct {
		maxOutputSize uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerFingerprint
	}{
		{
			name: "set",
			args: args{maxOutputSize: 1},
			want: &AnalyzerFingerprint{maxOutputSize: &maxOutputSizeSet},
		},
		{
			name: "change",
			fields: fields{maxOutputSize: &maxOutputSizeSet},
			args: args{maxOutputSize: 2},
			want: &AnalyzerFingerprint{maxOutputSize: &maxOutputSizeChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.SetMaxOutputSize(tt.args.maxOutputSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxOutputSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerFingerprint_SetSeparator(t *testing.T) {
	separatorSet := "|"
	separatorChange := "!"

	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
	}
	type args struct {
		separator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerFingerprint
	}{
		{
			name: "set",
			args: args{separator: "|"},
			want: &AnalyzerFingerprint{separator: &separatorSet},
		},
		{
			name: "change",
			fields: fields{separator: &separatorSet},
			args: args{separator: "!"},
			want: &AnalyzerFingerprint{separator: &separatorChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.SetSeparator(tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerFingerprint_SetStopWordsPath(t *testing.T) {
	stopWordsPathSet := "/temp/stopWords.txt"
	stopWordsPathChange := "/temp/stopWords.v1.txt"

	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
	}
	type args struct {
		stopWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerFingerprint
	}{
		{
			name: "set",
			args: args{stopWordsPath: "/temp/stopWords.txt"},
			want: &AnalyzerFingerprint{stopWordsPath: &stopWordsPathSet},
		},
		{
			name: "change",
			fields: fields{stopWordsPath: &stopWordsPathSet},
			args: args{stopWordsPath: "/temp/stopWords.v1.txt"},
			want: &AnalyzerFingerprint{stopWordsPath: &stopWordsPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.SetStopWordsPath(tt.args.stopWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStopWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerFingerprint_Source(t *testing.T) {
	separator := "|"
	maxOutputSize := uint8(1)
	stopWordsPath := "/temp/stopWords.txt"

	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
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
				"type": AnalyzerTypeFingerprint,
			},
		},
		{
			name: "separator",
			fields: fields{separator: &separator},
			want: map[string]interface{}{
				"type": AnalyzerTypeFingerprint,
				"separator": "|",
			},
		},
		{
			name: "separator",
			fields: fields{maxOutputSize: &maxOutputSize},
			want: map[string]interface{}{
				"type": AnalyzerTypeFingerprint,
				"max_output_size": maxOutputSize,
			},
		},
		{
			name: "stop_words_path",
			fields: fields{stopWordsPath: &stopWordsPath},
			want: map[string]interface{}{
				"type": AnalyzerTypeFingerprint,
				"stopwords_path": stopWordsPath,
			},
		},
		{
			name: "stopwords",
			fields: fields{stopWords: []string{"test", "test1"}},
			want: map[string]interface{}{
				"type": AnalyzerTypeFingerprint,
				"stopwords": []string{"test", "test1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			got, err := analyzer.Source()
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

func TestAnalyzerFingerprint_Type(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		maxOutputSize *uint8
		separator     *string
		stopWordsPath *string
		stopWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerType
	}{
		{
			want: "fingerprint",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerFingerprint{
				name:          tt.fields.name,
				maxOutputSize: tt.fields.maxOutputSize,
				separator:     tt.fields.separator,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_AddFlags(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	type args struct {
		flags []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerPattern
	}{
		{
			name: "RegularFlagCanonEq",
			args: args{flags: []JavaRegularFlag{RegularFlagCanonEq}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagCanonEq}},
		},
		{
			name: "RegularFlagCaseInsensitive",
			args: args{flags: []JavaRegularFlag{RegularFlagCaseInsensitive}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagCaseInsensitive}},
		},
		{
			name: "RegularFlagComments",
			args: args{flags: []JavaRegularFlag{RegularFlagComments}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagComments}},
		},
		{
			name: "RegularFlagDotAll",
			args: args{flags: []JavaRegularFlag{RegularFlagDotAll}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagDotAll}},
		},
		{
			name: "RegularFlagLiteral",
			args: args{flags: []JavaRegularFlag{RegularFlagLiteral}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagLiteral}},
		},
		{
			name: "RegularFlagLiteral",
			args: args{flags: []JavaRegularFlag{RegularFlagMultiline}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagMultiline}},
		},
		{
			name: "RegularFlagUnicodeCase",
			args: args{flags: []JavaRegularFlag{RegularFlagUnicodeCase}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagUnicodeCase}},
		},
		{
			name: "RegularFlagUnicodeCharacterClass",
			args: args{flags: []JavaRegularFlag{RegularFlagUnicodeCharacterClass}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagUnicodeCharacterClass}},
		},
		{
			name: "RegularFlagUnixLines",
			args: args{flags: []JavaRegularFlag{RegularFlagUnixLines}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagUnixLines}},
		},
		{
			name: "many",
			args: args{flags: []JavaRegularFlag{RegularFlagUnixLines, RegularFlagCanonEq}},
			want: &AnalyzerPattern{flags: []JavaRegularFlag{RegularFlagUnixLines, RegularFlagCanonEq}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.AddFlags(tt.args.flags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_AddStopWords(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	type args struct {
		stopWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerPattern
	}{
		{
			name: "set",
			args: args{stopWords: []string{"test"}},
			want: &AnalyzerPattern{stopWords: []string{"test"}},
		},
		{
			name: "add",
			fields: fields{stopWords: []string{"test"}},
			args: args{stopWords: []string{"test2"}},
			want: &AnalyzerPattern{stopWords: []string{"test", "test2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.AddStopWords(tt.args.stopWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_Name(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_SetLowercase(t *testing.T) {
	lowercaseTrue := true
	lowercaseFalse := false

	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	type args struct {
		lowercase bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerPattern
	}{
		{
			name: "false",
			args: args{lowercase: false},
			want: &AnalyzerPattern{lowercase: &lowercaseFalse},
		},
		{
			name: "true",
			args: args{lowercase: true},
			want: &AnalyzerPattern{lowercase: &lowercaseTrue},
		},
		{
			name: "change",
			fields: fields{lowercase: &lowercaseFalse},
			args: args{lowercase: true},
			want: &AnalyzerPattern{lowercase: &lowercaseTrue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.SetLowercase(tt.args.lowercase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_SetPattern(t *testing.T) {
	patternSet := "test"
	patternChange := "test1"

	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerPattern
	}{
		{
			name: "set",
			args: args{pattern: "test"},
			want: &AnalyzerPattern{pattern: &patternSet},
		},
		{
			name: "change",
			fields: fields{pattern: &patternSet},
			args: args{pattern: "test1"},
			want: &AnalyzerPattern{pattern: &patternChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.SetPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_SetStopWordsPath(t *testing.T) {
	stopWordsPathSet := "/temp/stopWords.txt"
	stopWordsPathChange := "/temp/stopWords.v1.txt"

	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	type args struct {
		stopWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerPattern
	}{
		{
			name: "set",
			args: args{stopWordsPath: "/temp/stopWords.txt"},
			want: &AnalyzerPattern{stopWordsPath: &stopWordsPathSet},
		},
		{
			name: "change",
			fields: fields{stopWordsPath: &stopWordsPathSet},
			args: args{stopWordsPath: "/temp/stopWords.v1.txt"},
			want: &AnalyzerPattern{stopWordsPath: &stopWordsPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.SetStopWordsPath(tt.args.stopWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStopWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerPattern_Source(t *testing.T) {
	pattern := "test"
	lowercase := true
	stopWordsPath := "/temp/stopWords.txt"

	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
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
				"type": AnalyzerTypePattern,
			},
		},
		{
			name: "pattern",
			fields: fields{pattern: &pattern},
			want: map[string]interface{}{
				"type": AnalyzerTypePattern,
				"pattern": "test",
			},
		},
		{
			name: "flags",
			fields: fields{flags: []JavaRegularFlag{RegularFlagCanonEq, RegularFlagCaseInsensitive}},
			want: map[string]interface{}{
				"type": AnalyzerTypePattern,
				"flags": "CANON_EQ|CASE_INSENSITIVE",
			},
		},
		{
			name: "lowercase",
			fields: fields{lowercase: &lowercase},
			want: map[string]interface{}{
				"type": AnalyzerTypePattern,
				"lowercase": true,
			},
		},
		{
			name: "stopwords_path",
			fields: fields{stopWordsPath: &stopWordsPath},
			want: map[string]interface{}{
				"type": AnalyzerTypePattern,
				"stopwords_path": "/temp/stopWords.txt",
			},
		},
		{
			name: "stopwords",
			fields: fields{stopWords: []string{"test", "test1"}},
			want: map[string]interface{}{
				"type": AnalyzerTypePattern,
				"stopwords": []string{"test", "test1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			got, err := analyzer.Source()
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

func TestAnalyzerPattern_Type(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		lowercase     *bool
		pattern       *string
		stopWordsPath *string
		stopWords     []string
		flags         []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerType
	}{
		{
			want: AnalyzerTypePattern,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerPattern{
				name:          tt.fields.name,
				lowercase:     tt.fields.lowercase,
				pattern:       tt.fields.pattern,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
				flags:         tt.fields.flags,
			}
			if got := analyzer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStandard_AddStopWords(t *testing.T) {
	type fields struct {
		name           AnalyzerName
		maxTokenLength *uint8
		stopWordsPath  *string
		stopWords      []string
	}
	type args struct {
		stopWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerStandard
	}{
		{
			name: "set",
			args: args{stopWords: []string{"test"}},
			want: &AnalyzerStandard{stopWords: []string{"test"}},
		},
		{
			name: "add",
			fields: fields{stopWords: []string{"test"}},
			args: args{stopWords: []string{"test2"}},
			want: &AnalyzerStandard{stopWords: []string{"test", "test2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStandard{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
				stopWordsPath:  tt.fields.stopWordsPath,
				stopWords:      tt.fields.stopWords,
			}
			if got := analyzer.AddStopWords(tt.args.stopWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStandard_Name(t *testing.T) {
	type fields struct {
		name           AnalyzerName
		maxTokenLength *uint8
		stopWordsPath  *string
		stopWords      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStandard{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
				stopWordsPath:  tt.fields.stopWordsPath,
				stopWords:      tt.fields.stopWords,
			}
			if got := analyzer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStandard_SetMaxTokenLength(t *testing.T) {
	maxTokenLengthSet := uint8(1)
	maxTokenLengthChange := uint8(2)

	type fields struct {
		name           AnalyzerName
		maxTokenLength *uint8
		stopWordsPath  *string
		stopWords      []string
	}
	type args struct {
		maxTokenLength uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerStandard
	}{
		{
			name: "set",
			args: args{maxTokenLength: 1},
			want: &AnalyzerStandard{maxTokenLength: &maxTokenLengthSet},
		},
		{
			name: "change",
			fields: fields{maxTokenLength: &maxTokenLengthSet},
			args: args{maxTokenLength: 2},
			want: &AnalyzerStandard{maxTokenLength: &maxTokenLengthChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStandard{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
				stopWordsPath:  tt.fields.stopWordsPath,
				stopWords:      tt.fields.stopWords,
			}
			if got := analyzer.SetMaxTokenLength(tt.args.maxTokenLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxTokenLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStandard_SetStopWordsPath(t *testing.T) {
	stopWordsPathSet := "/temp/stopWords.txt"
	stopWordsPathChange := "/temp/stopWords.v1.txt"

	type fields struct {
		name           AnalyzerName
		maxTokenLength *uint8
		stopWordsPath  *string
		stopWords      []string
	}
	type args struct {
		stopWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerStandard
	}{
		{
			name: "set",
			args: args{stopWordsPath: "/temp/stopWords.txt"},
			want: &AnalyzerStandard{stopWordsPath: &stopWordsPathSet},
		},
		{
			name: "change",
			fields: fields{stopWordsPath: &stopWordsPathSet},
			args: args{stopWordsPath: "/temp/stopWords.v1.txt"},
			want: &AnalyzerStandard{stopWordsPath: &stopWordsPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStandard{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
				stopWordsPath:  tt.fields.stopWordsPath,
				stopWords:      tt.fields.stopWords,
			}
			if got := analyzer.SetStopWordsPath(tt.args.stopWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStopWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStandard_Source(t *testing.T) {
	maxTokenLength := uint8(1)
	stopWordsPath := "/temp/stopWords.txt"

	type fields struct {
		name           AnalyzerName
		maxTokenLength *uint8
		stopWordsPath  *string
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
				"type": AnalyzerTypeStandard,
			},
		},
		{
			name: "max_token_length",
			fields: fields{maxTokenLength: &maxTokenLength},
			want: map[string]interface{}{
				"type": AnalyzerTypeStandard,
				"max_token_length": uint8(1),
			},
		},
		{
			name: "stopwords_path",
			fields: fields{stopWordsPath: &stopWordsPath},
			want: map[string]interface{}{
				"type": AnalyzerTypeStandard,
				"stopwords_path": "/temp/stopWords.txt",
			},
		},
		{
			name: "stopwords",
			fields: fields{stopWords: []string{"test", "test1"}},
			want: map[string]interface{}{
				"type": AnalyzerTypeStandard,
				"stopwords": []string{"test", "test1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStandard{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
				stopWordsPath:  tt.fields.stopWordsPath,
				stopWords:      tt.fields.stopWords,
			}
			got, err := analyzer.Source()
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

func TestAnalyzerStandard_Type(t *testing.T) {
	type fields struct {
		name           AnalyzerName
		maxTokenLength *uint8
		stopWordsPath  *string
		stopWords      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerType
	}{
		{
			want: AnalyzerTypeStandard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStandard{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
				stopWordsPath:  tt.fields.stopWordsPath,
				stopWords:      tt.fields.stopWords,
			}
			if got := analyzer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStop_AddStopWords(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		stopWordsPath *string
		stopWords     []string
	}
	type args struct {
		stopWords []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerStop
	}{
		{
			name: "set",
			args: args{stopWords: []string{"test"}},
			want: &AnalyzerStop{stopWords: []string{"test"}},
		},
		{
			name: "add",
			fields: fields{stopWords: []string{"test"}},
			args: args{stopWords: []string{"test2"}},
			want: &AnalyzerStop{stopWords: []string{"test", "test2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStop{
				name:          tt.fields.name,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.AddStopWords(tt.args.stopWords...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStopWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStop_Name(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		stopWordsPath *string
		stopWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStop{
				name:          tt.fields.name,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStop_SetStopWordsPath(t *testing.T) {
	stopWordsPathSet := "/temp/stopWords.txt"
	stopWordsPathChange := "/temp/stopWords.v1.txt"

	type fields struct {
		name          AnalyzerName
		stopWordsPath *string
		stopWords     []string
	}
	type args struct {
		stopWordsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AnalyzerStop
	}{
		{
			name: "set",
			args: args{stopWordsPath: "/temp/stopWords.txt"},
			want: &AnalyzerStop{stopWordsPath: &stopWordsPathSet},
		},
		{
			name: "change",
			fields: fields{stopWordsPath: &stopWordsPathSet},
			args: args{stopWordsPath: "/temp/stopWords.v1.txt"},
			want: &AnalyzerStop{stopWordsPath: &stopWordsPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStop{
				name:          tt.fields.name,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.SetStopWordsPath(tt.args.stopWordsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStopWordsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzerStop_Source(t *testing.T) {
	stopWordsPath := "/temp/stopWords.txt"

	type fields struct {
		name          AnalyzerName
		stopWordsPath *string
		stopWords     []string
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
				"type": AnalyzerTypeStop,
			},
		},
		{
			name: "stopwords_path",
			fields: fields{stopWordsPath: &stopWordsPath},
			want: map[string]interface{}{
				"type": AnalyzerTypeStop,
				"stopwords_path": "/temp/stopWords.txt",
			},
		},
		{
			name: "stopwords",
			fields: fields{stopWords: []string{"test", "test1"}},
			want: map[string]interface{}{
				"type": AnalyzerTypeStop,
				"stopwords": []string{"test", "test1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStop{
				name:          tt.fields.name,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			got, err := analyzer.Source()
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

func TestAnalyzerStop_Type(t *testing.T) {
	type fields struct {
		name          AnalyzerName
		stopWordsPath *string
		stopWords     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   AnalyzerType
	}{
		{
			want: AnalyzerTypeStop,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := &AnalyzerStop{
				name:          tt.fields.name,
				stopWordsPath: tt.fields.stopWordsPath,
				stopWords:     tt.fields.stopWords,
			}
			if got := analyzer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAnalyzerFingerprint(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *AnalyzerFingerprint
	}{
		{
			want: &AnalyzerFingerprint{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyzerFingerprint(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyzerFingerprint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAnalyzerPattern(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *AnalyzerPattern
	}{
		{
			want: &AnalyzerPattern{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyzerPattern(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyzerPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAnalyzerStandard(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *AnalyzerStandard
	}{
		{
			want: &AnalyzerStandard{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyzerStandard(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyzerStandard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAnalyzerStop(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *AnalyzerStop
	}{
		{
			want: &AnalyzerStop{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyzerStop(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyzerStop() = %v, want %v", got, tt.want)
			}
		})
	}
}
