package settings

import (
	"reflect"
	"testing"
)

func TestCharFilterHtmlStrip_AddEscapedTags(t *testing.T) {
	type fields struct {
		name        CharFilterName
		escapedTags []string
	}
	type args struct {
		escapedTags []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CharFilterHtmlStrip
	}{
		{
			name: "set",
			args: args{escapedTags: []string{"test"}},
			want: &CharFilterHtmlStrip{escapedTags: []string{"test"}},
		},
		{
			name: "add",
			fields: fields{escapedTags: []string{"test"}},
			args: args{escapedTags: []string{"test1"}},
			want: &CharFilterHtmlStrip{escapedTags: []string{"test", "test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterHtmlStrip{
				name:        tt.fields.name,
				escapedTags: tt.fields.escapedTags,
			}
			if got := filter.AddEscapedTags(tt.args.escapedTags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEscapedTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterHtmlStrip_Name(t *testing.T) {
	type fields struct {
		name        CharFilterName
		escapedTags []string
	}
	tests := []struct {
		name   string
		fields fields
		want   CharFilterName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterHtmlStrip{
				name:        tt.fields.name,
				escapedTags: tt.fields.escapedTags,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterHtmlStrip_Source(t *testing.T) {
	type fields struct {
		name        CharFilterName
		escapedTags []string
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
				"type": CharFilterTypeHtmlStrip,
			},
		},
		{
			name: "escaped_tags",
			fields: fields{escapedTags: []string{"test"}},
			want: map[string]interface{}{
				"type": CharFilterTypeHtmlStrip,
				"escaped_tags": []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterHtmlStrip{
				name:        tt.fields.name,
				escapedTags: tt.fields.escapedTags,
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

func TestCharFilterHtmlStrip_Type(t *testing.T) {
	type fields struct {
		name        CharFilterName
		escapedTags []string
	}
	tests := []struct {
		name   string
		fields fields
		want   CharFilterType
	}{
		{
			want: CharFilterTypeHtmlStrip,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterHtmlStrip{
				name:        tt.fields.name,
				escapedTags: tt.fields.escapedTags,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterMapping_AddMappings(t *testing.T) {
	type fields struct {
		name         CharFilterName
		mappingsPath *string
		mappings     []*CharMapping
	}
	type args struct {
		mappings []*CharMapping
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CharFilterMapping
	}{
		{
			name: "set",
			args: args{mappings: []*CharMapping{{from: "й", to: "и"}, {from: "ё", to: "е"}}},
			want: &CharFilterMapping{mappings: []*CharMapping{{from: "й", to: "и"}, {from: "ё", to: "е"}}},
		},
		{
			name: "add",
			fields: fields{mappings: []*CharMapping{{from: "й", to: "и"}}},
			args: args{mappings: []*CharMapping{{from: "ё", to: "е"}}},
			want: &CharFilterMapping{mappings: []*CharMapping{{from: "й", to: "и"}, {from: "ё", to: "е"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterMapping{
				name:         tt.fields.name,
				mappingsPath: tt.fields.mappingsPath,
				mappings:     tt.fields.mappings,
			}
			if got := filter.AddMappings(tt.args.mappings...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMappings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterMapping_Name(t *testing.T) {
	type fields struct {
		name         CharFilterName
		mappingsPath *string
		mappings     []*CharMapping
	}
	tests := []struct {
		name   string
		fields fields
		want   CharFilterName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterMapping{
				name:         tt.fields.name,
				mappingsPath: tt.fields.mappingsPath,
				mappings:     tt.fields.mappings,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterMapping_SetMappingsPath(t *testing.T) {
	mappingsPathSet := "/tmp/mapping.txt"
	mappingsPathChange := "/tmp/mapping.v1.txt"

	type fields struct {
		name         CharFilterName
		mappingsPath *string
		mappings     []*CharMapping
	}
	type args struct {
		mappingsPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CharFilterMapping
	}{
		{
			name: "set",
			args: args{mappingsPath: "/tmp/mapping.txt"},
			want: &CharFilterMapping{mappingsPath: &mappingsPathSet},
		},
		{
			name: "change",
			fields: fields{mappingsPath: &mappingsPathSet},
			args: args{mappingsPath: "/tmp/mapping.v1.txt"},
			want: &CharFilterMapping{mappingsPath: &mappingsPathChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterMapping{
				name:         tt.fields.name,
				mappingsPath: tt.fields.mappingsPath,
				mappings:     tt.fields.mappings,
			}
			if got := filter.SetMappingsPath(tt.args.mappingsPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMappingsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterMapping_Source(t *testing.T) {
	mappingsPath := "/tmp/mapping.txt"

	type fields struct {
		name         CharFilterName
		mappingsPath *string
		mappings     []*CharMapping
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
				"type": CharFilterTypeMapping,
			},
		},
		{
			name: "mappings_path",
			fields: fields{mappingsPath: &mappingsPath},
			want: map[string]interface{}{
				"type": CharFilterTypeMapping,
				"mappings_path": "/tmp/mapping.txt",
			},
		},
		{
			name: "mappings",
			fields: fields{mappings: []*CharMapping{{from: "й", to: "и"}}},
			want: map[string]interface{}{
				"type": CharFilterTypeMapping,
				"mappings": []string{"й => и"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterMapping{
				name:         tt.fields.name,
				mappingsPath: tt.fields.mappingsPath,
				mappings:     tt.fields.mappings,
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

func TestCharFilterMapping_Type(t *testing.T) {
	type fields struct {
		name         CharFilterName
		mappingsPath *string
		mappings     []*CharMapping
	}
	tests := []struct {
		name   string
		fields fields
		want   CharFilterType
	}{
		{
			want: CharFilterTypeMapping,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterMapping{
				name:         tt.fields.name,
				mappingsPath: tt.fields.mappingsPath,
				mappings:     tt.fields.mappings,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterPatternReplace_AddFlags(t *testing.T) {
	type fields struct {
		name        CharFilterName
		pattern     *string
		replacement *string
		flags       []JavaRegularFlag
	}
	type args struct {
		flags []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CharFilterPatternReplace
	}{
		{
			name: "RegularFlagCanonEq",
			args: args{flags: []JavaRegularFlag{RegularFlagCanonEq}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagCanonEq}},
		},
		{
			name: "RegularFlagCaseInsensitive",
			args: args{flags: []JavaRegularFlag{RegularFlagCaseInsensitive}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagCaseInsensitive}},
		},
		{
			name: "RegularFlagComments",
			args: args{flags: []JavaRegularFlag{RegularFlagComments}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagComments}},
		},
		{
			name: "RegularFlagDotAll",
			args: args{flags: []JavaRegularFlag{RegularFlagDotAll}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagDotAll}},
		},
		{
			name: "RegularFlagLiteral",
			args: args{flags: []JavaRegularFlag{RegularFlagLiteral}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagLiteral}},
		},
		{
			name: "RegularFlagLiteral",
			args: args{flags: []JavaRegularFlag{RegularFlagMultiline}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagMultiline}},
		},
		{
			name: "RegularFlagUnicodeCase",
			args: args{flags: []JavaRegularFlag{RegularFlagUnicodeCase}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagUnicodeCase}},
		},
		{
			name: "RegularFlagUnicodeCharacterClass",
			args: args{flags: []JavaRegularFlag{RegularFlagUnicodeCharacterClass}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagUnicodeCharacterClass}},
		},
		{
			name: "RegularFlagUnixLines",
			args: args{flags: []JavaRegularFlag{RegularFlagUnixLines}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagUnixLines}},
		},
		{
			name: "many",
			args: args{flags: []JavaRegularFlag{RegularFlagUnixLines, RegularFlagCanonEq}},
			want: &CharFilterPatternReplace{flags: []JavaRegularFlag{RegularFlagUnixLines, RegularFlagCanonEq}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterPatternReplace{
				name:        tt.fields.name,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
				flags:       tt.fields.flags,
			}
			if got := filter.AddFlags(tt.args.flags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterPatternReplace_Name(t *testing.T) {
	type fields struct {
		name        CharFilterName
		pattern     *string
		replacement *string
		flags       []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		want   CharFilterName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterPatternReplace{
				name:        tt.fields.name,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
				flags:       tt.fields.flags,
			}
			if got := filter.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterPatternReplace_SetPattern(t *testing.T) {
	patternSet := "test"
	patternChange := "test1"

	type fields struct {
		name        CharFilterName
		pattern     *string
		replacement *string
		flags       []JavaRegularFlag
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CharFilterPatternReplace
	}{
		{
			name: "set",
			args: args{pattern: "test"},
			want: &CharFilterPatternReplace{pattern: &patternSet},
		},
		{
			name: "change",
			fields: fields{pattern: &patternSet},
			args: args{pattern: "test1"},
			want: &CharFilterPatternReplace{pattern: &patternChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterPatternReplace{
				name:        tt.fields.name,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
				flags:       tt.fields.flags,
			}
			if got := filter.SetPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterPatternReplace_SetReplacement(t *testing.T) {
	replacementSet := "{1}"
	replacementChange := "{1} - {2}"

	type fields struct {
		name        CharFilterName
		pattern     *string
		replacement *string
		flags       []JavaRegularFlag
	}
	type args struct {
		replacement string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CharFilterPatternReplace
	}{
		{
			name: "set",
			args: args{replacement: "{1}"},
			want: &CharFilterPatternReplace{replacement: &replacementSet},
		},
		{
			name: "change",
			fields: fields{replacement: &replacementSet},
			args: args{replacement: "{1} - {2}"},
			want: &CharFilterPatternReplace{replacement: &replacementChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterPatternReplace{
				name:        tt.fields.name,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
				flags:       tt.fields.flags,
			}
			if got := filter.SetReplacement(tt.args.replacement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharFilterPatternReplace_Source(t *testing.T) {
	pattern := "test"
	replacement := "{1} - {2}"

	type fields struct {
		name        CharFilterName
		pattern     *string
		replacement *string
		flags       []JavaRegularFlag
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
				"type": CharFilterTypePatternReplace,
			},
		},
		{
			name: "pattern",
			fields: fields{pattern: &pattern},
			want: map[string]interface{}{
				"type": CharFilterTypePatternReplace,
				"pattern": "test",
			},
		},
		{
			name: "replacement",
			fields: fields{replacement: &replacement},
			want: map[string]interface{}{
				"type": CharFilterTypePatternReplace,
				"replacement": "{1} - {2}",
			},
		},
		{
			name: "flags",
			fields: fields{flags: []JavaRegularFlag{RegularFlagCanonEq, RegularFlagCaseInsensitive}},
			want: map[string]interface{}{
				"type": CharFilterTypePatternReplace,
				"flags": "CANON_EQ|CASE_INSENSITIVE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterPatternReplace{
				name:        tt.fields.name,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
				flags:       tt.fields.flags,
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

func TestCharFilterPatternReplace_Type(t *testing.T) {
	type fields struct {
		name        CharFilterName
		pattern     *string
		replacement *string
		flags       []JavaRegularFlag
	}
	tests := []struct {
		name   string
		fields fields
		want   CharFilterType
	}{
		{
			want: CharFilterTypePatternReplace,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &CharFilterPatternReplace{
				name:        tt.fields.name,
				pattern:     tt.fields.pattern,
				replacement: tt.fields.replacement,
				flags:       tt.fields.flags,
			}
			if got := filter.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharMapping_String(t *testing.T) {
	type fields struct {
		from string
		to   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{from: "ё", to: "е"},
			want: "ё => е",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charMapping := &CharMapping{
				from: tt.fields.from,
				to:   tt.fields.to,
			}
			if got := charMapping.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCharFilterHtmlStrip(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *CharFilterHtmlStrip
	}{
		{
			args: args{name: "test"},
			want: &CharFilterHtmlStrip{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCharFilterHtmlStrip(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharFilterHtmlStrip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCharFilterMapping(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *CharFilterMapping
	}{
		{
			args: args{name: "test"},
			want: &CharFilterMapping{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCharFilterMapping(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharFilterMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCharFilterPatternReplace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *CharFilterPatternReplace
	}{
		{
			args: args{name: "test"},
			want: &CharFilterPatternReplace{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCharFilterPatternReplace(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharFilterPatternReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCharMapping(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name string
		args args
		want *CharMapping
	}{
		{
			args: args{from: "q", to: "r"},
			want: &CharMapping{from: "q", to: "r"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCharMapping(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
