package settings

import (
	"reflect"
	"testing"
)

func TestNewTokenizerCharGroup(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerCharGroup
	}{
		{
			args: args{name: "test"},
			want: &TokenizerCharGroup{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerCharGroup(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerCharGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerClassic(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerClassic
	}{
		{
			args: args{name: "test"},
			want: &TokenizerClassic{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerClassic(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerClassic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerEdgeNgram(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerNgram
	}{
		{
			args: args{name: "test"},
			want: &TokenizerNgram{name: "test", _type: TokenizerTypeEdgeNgram},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerEdgeNgram(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerEdgeNgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerKeyword(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerKeyword
	}{
		{
			args: args{name: "test"},
			want: &TokenizerKeyword{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerKeyword(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerNgram(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerNgram
	}{
		{
			args: args{name: "test"},
			want: &TokenizerNgram{name: "test", _type: TokenizerTypeNgram},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerNgram(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerNgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerPathHierarchy(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerPathHierarchy
	}{
		{
			args: args{name: "test"},
			want: &TokenizerPathHierarchy{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerPathHierarchy(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerPathHierarchy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerPattern(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerPattern
	}{
		{
			args: args{name: "test"},
			want: &TokenizerPattern{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerPattern(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerSimplePattern(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerSimplePattern
	}{
		{
			args: args{name: "test"},
			want: &TokenizerSimplePattern{name: "test", _type: TokenizerTypeSimplePattern},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerSimplePattern(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerSimplePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerSimplePatternSplit(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerSimplePattern
	}{
		{
			args: args{name: "test"},
			want: &TokenizerSimplePattern{name: "test", _type: TokenizerTypeSimplePatternSplit},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerSimplePatternSplit(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerSimplePatternSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerStandard(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerStandard
	}{
		{
			args: args{name: "test"},
			want: &TokenizerStandard{name: "test", _type: TokenizerTypeStandard},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerStandard(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerStandard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerUaxUrlEmail(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerStandard
	}{
		{
			args: args{name: "test"},
			want: &TokenizerStandard{name: "test", _type: TokenizerTypeUaxUrlEmail},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerUaxUrlEmail(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerUaxUrlEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenizerWhitespace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *TokenizerStandard
	}{
		{
			args: args{name: "test"},
			want: &TokenizerStandard{name: "test", _type: TokenizerTypeWhitespace},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizerWhitespace(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenizerWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerCharGroup_AddTokenizeOnChars(t *testing.T) {
	type fields struct {
		name            TokenizerName
		maxTokenLength  *uint32
		tokenizeOnChars []string
	}
	type args struct {
		tokenizeOnChars []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerCharGroup
	}{
		{
			name: "set",
			args: args{tokenizeOnChars: []string{"test"}},
			want: &TokenizerCharGroup{tokenizeOnChars: []string{"test"}},
		},
		{
			name: "change",
			fields: fields{tokenizeOnChars: []string{"test"}},
			args: args{tokenizeOnChars: []string{"test1"}},
			want: &TokenizerCharGroup{tokenizeOnChars: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerCharGroup{
				name:            tt.fields.name,
				maxTokenLength:  tt.fields.maxTokenLength,
				tokenizeOnChars: tt.fields.tokenizeOnChars,
			}
			if got := tokenizer.AddTokenizeOnChars(tt.args.tokenizeOnChars...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTokenizeOnChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerCharGroup_Name(t *testing.T) {
	type fields struct {
		name            TokenizerName
		maxTokenLength  *uint32
		tokenizeOnChars []string
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerCharGroup{
				name:            tt.fields.name,
				maxTokenLength:  tt.fields.maxTokenLength,
				tokenizeOnChars: tt.fields.tokenizeOnChars,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerCharGroup_SetMaxTokenLength(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name            TokenizerName
		maxTokenLength  *uint32
		tokenizeOnChars []string
	}
	type args struct {
		maxTokenLength uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerCharGroup
	}{
		{
			name: "set",
			args: args{maxTokenLength: 1},
			want: &TokenizerCharGroup{maxTokenLength: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxTokenLength: &numberSet},
			args: args{maxTokenLength: 2},
			want: &TokenizerCharGroup{maxTokenLength: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerCharGroup{
				name:            tt.fields.name,
				maxTokenLength:  tt.fields.maxTokenLength,
				tokenizeOnChars: tt.fields.tokenizeOnChars,
			}
			if got := tokenizer.SetMaxTokenLength(tt.args.maxTokenLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxTokenLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerCharGroup_Source(t *testing.T) {
	number := uint32(1)

	type fields struct {
		name            TokenizerName
		maxTokenLength  *uint32
		tokenizeOnChars []string
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
				"type": TokenizerTypeCharGroup,
			},
		},
		{
			name: "maxTokenLength",
			fields: fields{maxTokenLength: &number},
			want: map[string]interface{}{
				"type": TokenizerTypeCharGroup,
				"max_token_length": uint32(1),
			},
		},
		{
			name: "tokenize_on_chars",
			fields: fields{tokenizeOnChars: []string{"test"}},
			want: map[string]interface{}{
				"type": TokenizerTypeCharGroup,
				"tokenize_on_chars": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerCharGroup{
				name:            tt.fields.name,
				maxTokenLength:  tt.fields.maxTokenLength,
				tokenizeOnChars: tt.fields.tokenizeOnChars,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerCharGroup_Type(t *testing.T) {
	type fields struct {
		name            TokenizerName
		maxTokenLength  *uint32
		tokenizeOnChars []string
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			want: TokenizerTypeCharGroup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerCharGroup{
				name:            tt.fields.name,
				maxTokenLength:  tt.fields.maxTokenLength,
				tokenizeOnChars: tt.fields.tokenizeOnChars,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerClassic_Name(t *testing.T) {
	type fields struct {
		name           TokenizerName
		maxTokenLength *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerClassic{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerClassic_SetMaxTokenLength(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name           TokenizerName
		maxTokenLength *uint32
	}
	type args struct {
		maxTokenLength uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerClassic
	}{
		{
			name: "set",
			args: args{maxTokenLength: 1},
			want: &TokenizerClassic{maxTokenLength: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxTokenLength: &numberSet},
			args: args{maxTokenLength: 2},
			want: &TokenizerClassic{maxTokenLength: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerClassic{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			if got := tokenizer.SetMaxTokenLength(tt.args.maxTokenLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxTokenLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerClassic_Source(t *testing.T) {
	number := uint32(1)

	type fields struct {
		name           TokenizerName
		maxTokenLength *uint32
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
				"type": TokenizerTypeClassic,
			},
		},
		{
			name: "max_token_length",
			fields: fields{maxTokenLength: &number},
			want: map[string]interface{}{
				"type": TokenizerTypeClassic,
				"max_token_length": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerClassic{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerClassic_Type(t *testing.T) {
	type fields struct {
		name           TokenizerName
		maxTokenLength *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			want: TokenizerTypeClassic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerClassic{
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerKeyword_Name(t *testing.T) {
	type fields struct {
		name       TokenizerName
		bufferSize *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerKeyword{
				name:       tt.fields.name,
				bufferSize: tt.fields.bufferSize,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerKeyword_SetBufferSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name       TokenizerName
		bufferSize *uint32
	}
	type args struct {
		bufferSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerKeyword
	}{
		{
			name: "set",
			args: args{bufferSize: 1},
			want: &TokenizerKeyword{bufferSize: &numberSet},
		},
		{
			name: "change",
			fields: fields{bufferSize: &numberSet},
			args: args{bufferSize: 2},
			want: &TokenizerKeyword{bufferSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerKeyword{
				name:       tt.fields.name,
				bufferSize: tt.fields.bufferSize,
			}
			if got := tokenizer.SetBufferSize(tt.args.bufferSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerKeyword_Source(t *testing.T) {
	number := uint32(1)

	type fields struct {
		name       TokenizerName
		bufferSize *uint32
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
				"type": TokenizerTypeKeyword,
			},
		},
		{
			name: "buffer_size",
			fields: fields{bufferSize: &number},
			want: map[string]interface{}{
				"type": TokenizerTypeKeyword,
				"buffer_size": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerKeyword{
				name:       tt.fields.name,
				bufferSize: tt.fields.bufferSize,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerKeyword_Type(t *testing.T) {
	type fields struct {
		name       TokenizerName
		bufferSize *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			want: TokenizerTypeKeyword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerKeyword{
				name:       tt.fields.name,
				bufferSize: tt.fields.bufferSize,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerNgram_AddCustomTokenChars(t *testing.T) {
	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	type args struct {
		customTokenChars []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerNgram
	}{
		{
			name: "set",
			args: args{customTokenChars: []string{"test"}},
			want: &TokenizerNgram{customTokenChars: []string{"test"}},
		},
		{
			name: "change",
			fields: fields{customTokenChars: []string{"test"}},
			args: args{customTokenChars: []string{"test1"}},
			want: &TokenizerNgram{customTokenChars: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			if got := tokenizer.AddCustomTokenChars(tt.args.customTokenChars...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCustomTokenChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerNgram_AddTokenChars(t *testing.T) {
	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	type args struct {
		tokenChars []EdgeNgramTokenChars
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerNgram
	}{
		{
			name: "edge_ngram_token_chars_letter",
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsLetter}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsLetter}},
		},
		{
			name: "edge_ngram_token_chars_digit",
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsDigit}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsDigit}},
		},
		{
			name: "edge_ngram_token_chars_whitespace",
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsWhitespace}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsWhitespace}},
		},
		{
			name: "edge_ngram_token_chars_punctuation",
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsPunctuation}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsPunctuation}},
		},
		{
			name: "edge_ngram_token_chars_symbol",
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsSymbol}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsSymbol}},
		},
		{
			name: "edge_ngram_token_chars_custom",
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsCustom}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsCustom}},
		},
		{
			name: "add",
			fields: fields{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsCustom}},
			args: args{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsSymbol}},
			want: &TokenizerNgram{tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsCustom, EdgeNgramTokenCharsSymbol}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			if got := tokenizer.AddTokenChars(tt.args.tokenChars...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTokenChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerNgram_Name(t *testing.T) {
	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerNgram_SetMaxGram(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	type args struct {
		maxGram uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerNgram
	}{
		{
			name: "set",
			args: args{maxGram: 1},
			want: &TokenizerNgram{maxGram: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxGram: &numberSet},
			args: args{maxGram: 2},
			want: &TokenizerNgram{maxGram: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			if got := tokenizer.SetMaxGram(tt.args.maxGram); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerNgram_SetMinGram(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	type args struct {
		minGram uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerNgram
	}{
		{
			name: "set",
			args: args{minGram: 1},
			want: &TokenizerNgram{minGram: &numberSet},
		},
		{
			name: "change",
			fields: fields{minGram: &numberSet},
			args: args{minGram: 2},
			want: &TokenizerNgram{minGram: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			if got := tokenizer.SetMinGram(tt.args.minGram); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMinGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerNgram_Source(t *testing.T) {
	number := uint32(1)

	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "tokenizer_type_ngram",
			fields: fields{_type: TokenizerTypeNgram},
			want: map[string]interface{}{
				"type": TokenizerTypeNgram,
			},
		},
		{
			name: "tokenizer_type_edge_ngram",
			fields: fields{_type: TokenizerTypeEdgeNgram},
			want: map[string]interface{}{
				"type": TokenizerTypeEdgeNgram,
			},
		},
		{
			name: "min_gram",
			fields: fields{_type: TokenizerTypeEdgeNgram, minGram: &number},
			want: map[string]interface{}{
				"type": TokenizerTypeEdgeNgram,
				"min_gram": uint32(1),
			},
		},
		{
			name: "max_gram",
			fields: fields{_type: TokenizerTypeEdgeNgram, maxGram: &number},
			want: map[string]interface{}{
				"type": TokenizerTypeEdgeNgram,
				"max_gram": uint32(1),
			},
		},
		{
			name: "token_chars",
			fields: fields{_type: TokenizerTypeEdgeNgram, tokenChars: []EdgeNgramTokenChars{EdgeNgramTokenCharsCustom}},
			want: map[string]interface{}{
				"type": TokenizerTypeEdgeNgram,
				"token_chars": []EdgeNgramTokenChars{EdgeNgramTokenCharsCustom},
			},
		},
		{
			name: "custom_token_chars",
			fields: fields{_type: TokenizerTypeEdgeNgram, customTokenChars: []string{"test"}},
			want: map[string]interface{}{
				"type": TokenizerTypeEdgeNgram,
				"custom_token_chars": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerNgram_Type(t *testing.T) {
	type fields struct {
		_type            TokenizerType
		name             TokenizerName
		minGram          *uint32
		maxGram          *uint32
		tokenChars       []EdgeNgramTokenChars
		customTokenChars []string
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			name: "tokenizer_type_ngram",
			fields: fields{_type: TokenizerTypeNgram},
			want: TokenizerTypeNgram,
		},
		{
			name: "tokenizer_type_edge_ngram",
			fields: fields{_type: TokenizerTypeEdgeNgram},
			want: TokenizerTypeEdgeNgram,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerNgram{
				_type:            tt.fields._type,
				name:             tt.fields.name,
				minGram:          tt.fields.minGram,
				maxGram:          tt.fields.maxGram,
				tokenChars:       tt.fields.tokenChars,
				customTokenChars: tt.fields.customTokenChars,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_Name(t *testing.T) {
	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_SetBufferSize(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	type args struct {
		bufferSize uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPathHierarchy
	}{
		{
			name: "set",
			args: args{bufferSize: 1},
			want: &TokenizerPathHierarchy{bufferSize: &numberSet},
		},
		{
			name: "change",
			fields: fields{bufferSize: &numberSet},
			args: args{bufferSize: 2},
			want: &TokenizerPathHierarchy{bufferSize: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.SetBufferSize(tt.args.bufferSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_SetDelimiter(t *testing.T) {
	stringSet := "set"
	stringChange := "change"

	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	type args struct {
		delimiter string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPathHierarchy
	}{
		{
			name: "set",
			args: args{delimiter: "set"},
			want: &TokenizerPathHierarchy{delimiter: &stringSet},
		},
		{
			name: "change",
			fields: fields{delimiter: &stringSet},
			args: args{delimiter: "change"},
			want: &TokenizerPathHierarchy{delimiter: &stringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.SetDelimiter(tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_SetReplacement(t *testing.T) {
	stringSet := "set"
	stringChange := "change"

	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	type args struct {
		replacement string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPathHierarchy
	}{
		{
			name: "set",
			args: args{replacement: "set"},
			want: &TokenizerPathHierarchy{replacement: &stringSet},
		},
		{
			name: "change",
			fields: fields{replacement: &stringSet},
			args: args{replacement: "change"},
			want: &TokenizerPathHierarchy{replacement: &stringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.SetReplacement(tt.args.replacement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_SetReverse(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	type args struct {
		reverse bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPathHierarchy
	}{
		{
			name: "set",
			args: args{reverse: true},
			want: &TokenizerPathHierarchy{reverse: &boolTrue},
		},
		{
			name: "change",
			fields: fields{reverse: &boolTrue},
			args: args{reverse: false},
			want: &TokenizerPathHierarchy{reverse: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.SetReverse(tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_SetSkip(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	type args struct {
		skip uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPathHierarchy
	}{
		{
			name: "set",
			args: args{skip: 1},
			want: &TokenizerPathHierarchy{skip: &numberSet},
		},
		{
			name: "change",
			fields: fields{skip: &numberSet},
			args: args{skip: 2},
			want: &TokenizerPathHierarchy{skip: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.SetSkip(tt.args.skip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPathHierarchy_Source(t *testing.T) {
	number := uint32(1)
	boolTrue := true
	stringTest := "test"

	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
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
				"type": TokenizerTypePathHierarchy,
			},
		},
		{
			name: "reverse",
			fields: fields{reverse: &boolTrue},
			want: map[string]interface{}{
				"type": TokenizerTypePathHierarchy,
				"reverse": true,
			},
		},
		{
			name: "delimiter",
			fields: fields{delimiter: &stringTest},
			want: map[string]interface{}{
				"type": TokenizerTypePathHierarchy,
				"delimiter": "test",
			},
		},
		{
			name: "replacement",
			fields: fields{replacement: &stringTest},
			want: map[string]interface{}{
				"type": TokenizerTypePathHierarchy,
				"replacement": "test",
			},
		},
		{
			name: "buffer_size",
			fields: fields{bufferSize: &number},
			want: map[string]interface{}{
				"type": TokenizerTypePathHierarchy,
				"buffer_size": uint32(1),
			},
		},
		{
			name: "skip",
			fields: fields{skip: &number},
			want: map[string]interface{}{
				"type": TokenizerTypePathHierarchy,
				"skip": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerPathHierarchy_Type(t *testing.T) {
	type fields struct {
		name        TokenizerName
		reverse     *bool
		delimiter   *string
		replacement *string
		bufferSize  *uint32
		skip        *uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			want: TokenizerTypePathHierarchy,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPathHierarchy{
				name:        tt.fields.name,
				reverse:     tt.fields.reverse,
				delimiter:   tt.fields.delimiter,
				replacement: tt.fields.replacement,
				bufferSize:  tt.fields.bufferSize,
				skip:        tt.fields.skip,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPattern_AddFlags(t *testing.T) {
	type fields struct {
		name    TokenizerName
		pattern *string
		group   *uint8
		flags   []JavaRegularFlag
	}
	type args struct {
		flags []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPattern
	}{
		{
			name: "set",
			args: args{flags: []JavaRegularFlag{RegularFlagCanonEq}},
			want: &TokenizerPattern{flags: []JavaRegularFlag{RegularFlagCanonEq}},
		},
		{
			name: "add",
			fields: fields{flags: []JavaRegularFlag{RegularFlagCaseInsensitive}},
			args: args{flags: []JavaRegularFlag{RegularFlagCanonEq}},
			want: &TokenizerPattern{flags: []JavaRegularFlag{RegularFlagCaseInsensitive, RegularFlagCanonEq}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPattern{
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
				group:   tt.fields.group,
				flags:   tt.fields.flags,
			}
			if got := tokenizer.AddFlags(tt.args.flags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPattern_Name(t *testing.T) {
	type fields struct {
		name    TokenizerName
		pattern *string
		group   *uint8
		flags   []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPattern{
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
				group:   tt.fields.group,
				flags:   tt.fields.flags,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPattern_SetGroup(t *testing.T) {
	numberSet := uint8(1)
	numberChange := uint8(2)

	type fields struct {
		name    TokenizerName
		pattern *string
		group   *uint8
		flags   []JavaRegularFlag
	}
	type args struct {
		group uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPattern
	}{
		{
			name: "set",
			args: args{group: 1},
			want: &TokenizerPattern{group: &numberSet},
		},
		{
			name: "change",
			fields: fields{group: &numberSet},
			args: args{group: 2},
			want: &TokenizerPattern{group: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPattern{
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
				group:   tt.fields.group,
				flags:   tt.fields.flags,
			}
			if got := tokenizer.SetGroup(tt.args.group); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPattern_SetPattern(t *testing.T) {
	stringSet := "set"
	stringChange := "change"

	type fields struct {
		name    TokenizerName
		pattern *string
		group   *uint8
		flags   []JavaRegularFlag
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerPattern
	}{
		{
			name: "set",
			args: args{pattern: "set"},
			want: &TokenizerPattern{pattern: &stringSet},
		},
		{
			name: "change",
			fields: fields{pattern: &stringSet},
			args: args{pattern: "change"},
			want: &TokenizerPattern{pattern: &stringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPattern{
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
				group:   tt.fields.group,
				flags:   tt.fields.flags,
			}
			if got := tokenizer.SetPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerPattern_Source(t *testing.T) {
	stringTest := "test"
	number := uint8(1)

	type fields struct {
		name    TokenizerName
		pattern *string
		group   *uint8
		flags   []JavaRegularFlag
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
				"type": TokenizerTypePattern,
			},
		},
		{
			name: "pattern",
			fields: fields{pattern: &stringTest},
			want: map[string]interface{}{
				"type": TokenizerTypePattern,
				"pattern": "test",
			},
		},
		{
			name: "group",
			fields: fields{group: &number},
			want: map[string]interface{}{
				"type": TokenizerTypePattern,
				"group": uint8(1),
			},
		},
		{
			name: "flags",
			fields: fields{flags: []JavaRegularFlag{RegularFlagCanonEq,RegularFlagCaseInsensitive}},
			want: map[string]interface{}{
				"type": TokenizerTypePattern,
				"flags": "CANON_EQ|CASE_INSENSITIVE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPattern{
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
				group:   tt.fields.group,
				flags:   tt.fields.flags,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerPattern_Type(t *testing.T) {
	type fields struct {
		name    TokenizerName
		pattern *string
		group   *uint8
		flags   []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			want: TokenizerTypePattern,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerPattern{
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
				group:   tt.fields.group,
				flags:   tt.fields.flags,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerSimplePattern_Name(t *testing.T) {
	type fields struct {
		_type   TokenizerType
		name    TokenizerName
		pattern *string
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerSimplePattern{
				_type:   tt.fields._type,
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerSimplePattern_SetPattern(t *testing.T) {
	stringSet := "set"
	stringChange := "change"

	type fields struct {
		_type   TokenizerType
		name    TokenizerName
		pattern *string
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerSimplePattern
	}{
		{
			name: "set",
			args: args{pattern: "set"},
			want: &TokenizerSimplePattern{pattern: &stringSet},
		},
		{
			name: "change",
			fields: fields{pattern: &stringSet},
			args: args{pattern: "change"},
			want: &TokenizerSimplePattern{pattern: &stringChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerSimplePattern{
				_type:   tt.fields._type,
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
			}
			if got := tokenizer.SetPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerSimplePattern_Source(t *testing.T) {
	type fields struct {
		_type   TokenizerType
		name    TokenizerName
		pattern *string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "tokenizer_type_simple_pattern",
			fields: fields{_type: TokenizerTypeSimplePattern},
			want: map[string]interface{}{
				"type": TokenizerTypeSimplePattern,
			},
		},
		{
			name: "tokenizer_type_simple_pattern_split",
			fields: fields{_type: TokenizerTypeSimplePatternSplit},
			want: map[string]interface{}{
				"type": TokenizerTypeSimplePatternSplit,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerSimplePattern{
				_type:   tt.fields._type,
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerSimplePattern_Type(t *testing.T) {
	type fields struct {
		_type   TokenizerType
		name    TokenizerName
		pattern *string
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			name: "tokenizer_type_simple_pattern",
			fields: fields{_type: TokenizerTypeSimplePattern},
			want: TokenizerTypeSimplePattern,
		},
		{
			name: "tokenizer_type_simple_pattern_split",
			fields: fields{_type: TokenizerTypeSimplePatternSplit},
			want: TokenizerTypeSimplePatternSplit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerSimplePattern{
				_type:   tt.fields._type,
				name:    tt.fields.name,
				pattern: tt.fields.pattern,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerStandard_Name(t *testing.T) {
	type fields struct {
		_type          TokenizerType
		name           TokenizerName
		maxTokenLength *uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerStandard{
				_type:          tt.fields._type,
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			if got := tokenizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerStandard_SetMaxTokenLength(t *testing.T) {
	numberSet := uint16(1)
	numberChange := uint16(2)

	type fields struct {
		_type          TokenizerType
		name           TokenizerName
		maxTokenLength *uint16
	}
	type args struct {
		maxTokenLength uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TokenizerStandard
	}{
		{
			name: "set",
			args: args{maxTokenLength: 1},
			want: &TokenizerStandard{maxTokenLength: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxTokenLength: &numberSet},
			args: args{maxTokenLength: 2},
			want: &TokenizerStandard{maxTokenLength: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerStandard{
				_type:          tt.fields._type,
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			if got := tokenizer.SetMaxTokenLength(tt.args.maxTokenLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxTokenLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizerStandard_Source(t *testing.T) {
	number := uint16(1)

	type fields struct {
		_type          TokenizerType
		name           TokenizerName
		maxTokenLength *uint16
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "tokenizer_type_standard",
			fields: fields{_type: TokenizerTypeStandard},
			want: map[string]interface{}{
				"type": TokenizerTypeStandard,
			},
		},
		{
			name: "tokenizer_type_uax_url_email",
			fields: fields{_type: TokenizerTypeUaxUrlEmail},
			want: map[string]interface{}{
				"type": TokenizerTypeUaxUrlEmail,
			},
		},
		{
			name: "tokenizer_type_whitespace",
			fields: fields{_type: TokenizerTypeWhitespace},
			want: map[string]interface{}{
				"type": TokenizerTypeWhitespace,
			},
		},
		{
			name: "max_token_length",
			fields: fields{maxTokenLength: &number, _type: TokenizerTypeWhitespace},
			want: map[string]interface{}{
				"type": TokenizerTypeWhitespace,
				"max_token_length": uint16(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerStandard{
				_type:          tt.fields._type,
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			got, err := tokenizer.Source()
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

func TestTokenizerStandard_Type(t *testing.T) {
	type fields struct {
		_type          TokenizerType
		name           TokenizerName
		maxTokenLength *uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   TokenizerType
	}{
		{
			name: "tokenizer_type_standard",
			fields: fields{_type: TokenizerTypeStandard},
			want: TokenizerTypeStandard,
		},
		{
			name: "tokenizer_type_uax_url_email",
			fields: fields{_type: TokenizerTypeUaxUrlEmail},
			want: TokenizerTypeUaxUrlEmail,
		},
		{
			name: "tokenizer_type_whitespace",
			fields: fields{_type: TokenizerTypeWhitespace},
			want: TokenizerTypeWhitespace,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &TokenizerStandard{
				_type:          tt.fields._type,
				name:           tt.fields.name,
				maxTokenLength: tt.fields.maxTokenLength,
			}
			if got := tokenizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
