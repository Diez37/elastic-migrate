package settings

import (
	"reflect"
	"testing"
)

func TestMapping_SetCoerce(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		coerce bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{coerce: true},
			want: &Mapping{coerce: &boolTrue},
		},
		{
			name: "change",
			fields: fields{coerce: &boolTrue},
			args: args{coerce: false},
			want: &Mapping{coerce: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetCoerce(tt.args.coerce); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCoerce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_SetDepth(t *testing.T) {
	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		depth *Limit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{depth: NewLimit(1)},
			want: &Mapping{depth: NewLimit(1)},
		},
		{
			name: "change",
			fields: fields{depth: NewLimit(1)},
			args: args{depth: NewLimit(2)},
			want: &Mapping{depth: NewLimit(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetDepth(tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_SetFieldNameLength(t *testing.T) {
	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		fieldNameLength *Limit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{fieldNameLength: NewLimit(1)},
			want: &Mapping{fieldNameLength: NewLimit(1)},
		},
		{
			name: "change",
			fields: fields{fieldNameLength: NewLimit(1)},
			args: args{fieldNameLength: NewLimit(2)},
			want: &Mapping{fieldNameLength: NewLimit(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetFieldNameLength(tt.args.fieldNameLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFieldNameLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_SetIgnoreMalformed(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		ignoreMalformed bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{ignoreMalformed: true},
			want: &Mapping{ignoreMalformed: &boolTrue},
		},
		{
			name: "change",
			fields: fields{ignoreMalformed: &boolTrue},
			args: args{ignoreMalformed: false},
			want: &Mapping{ignoreMalformed: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetIgnoreMalformed(tt.args.ignoreMalformed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIgnoreMalformed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_SetNestedFields(t *testing.T) {
	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		nestedFields *Limit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{nestedFields: NewLimit(1)},
			want: &Mapping{nestedFields: NewLimit(1)},
		},
		{
			name: "change",
			fields: fields{nestedFields: NewLimit(1)},
			args: args{nestedFields: NewLimit(2)},
			want: &Mapping{nestedFields: NewLimit(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetNestedFields(tt.args.nestedFields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNestedFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_SetNestedObjects(t *testing.T) {
	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		nestedObjects *Limit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{nestedObjects: NewLimit(1)},
			want: &Mapping{nestedObjects: NewLimit(1)},
		},
		{
			name: "change",
			fields: fields{nestedObjects: NewLimit(1)},
			args: args{nestedObjects: NewLimit(2)},
			want: &Mapping{nestedObjects: NewLimit(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetNestedObjects(tt.args.nestedObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNestedObjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_SetTotalFields(t *testing.T) {
	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	type args struct {
		totalFields *Limit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "set",
			args: args{totalFields: NewLimit(1)},
			want: &Mapping{totalFields: NewLimit(1)},
		},
		{
			name: "change",
			fields: fields{totalFields: NewLimit(1)},
			args: args{totalFields: NewLimit(2)},
			want: &Mapping{totalFields: NewLimit(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			if got := mapping.SetTotalFields(tt.args.totalFields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTotalFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_Source(t *testing.T) {
	boolTrue := true
	limitSrc, _ := NewLimit(1).Source()

	type fields struct {
		coerce          *bool
		ignoreMalformed *bool
		totalFields     *Limit
		depth           *Limit
		nestedFields    *Limit
		nestedObjects   *Limit
		fieldNameLength *Limit
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{},
		},
		{
			name: "coerce",
			fields: fields{coerce: &boolTrue},
			want: map[string]interface{}{
				"coerce": true,
			},
		},
		{
			name: "ignore_malformed",
			fields: fields{ignoreMalformed: &boolTrue},
			want: map[string]interface{}{
				"ignore_malformed": true,
			},
		},
		{
			name: "total_fields",
			fields: fields{totalFields: NewLimit(1)},
			want: map[string]interface{}{
				"total_fields": limitSrc,
			},
		},
		{
			name: "depth",
			fields: fields{depth: NewLimit(1)},
			want: map[string]interface{}{
				"depth": limitSrc,
			},
		},
		{
			name: "nested_fields",
			fields: fields{nestedFields: NewLimit(1)},
			want: map[string]interface{}{
				"nested_fields": limitSrc,
			},
		},
		{
			name: "nested_objects",
			fields: fields{nestedObjects: NewLimit(1)},
			want: map[string]interface{}{
				"nested_objects": limitSrc,
			},
		},
		{
			name: "field_name_length",
			fields: fields{fieldNameLength: NewLimit(1)},
			want: map[string]interface{}{
				"field_name_length": limitSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapping := &Mapping{
				coerce:          tt.fields.coerce,
				ignoreMalformed: tt.fields.ignoreMalformed,
				totalFields:     tt.fields.totalFields,
				depth:           tt.fields.depth,
				nestedFields:    tt.fields.nestedFields,
				nestedObjects:   tt.fields.nestedObjects,
				fieldNameLength: tt.fields.fieldNameLength,
			}
			got, err := mapping.Source()
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

func TestNewMapping(t *testing.T) {
	tests := []struct {
		name string
		want *Mapping
	}{
		{
			want: &Mapping{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMapping(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
