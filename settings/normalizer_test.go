package settings

import (
	"reflect"
	"testing"
)

func TestNewNormalizerCustom(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *NormalizerCustom
	}{
		{
			args: args{name: "test"},
			want: &NormalizerCustom{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNormalizerCustom(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNormalizerCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizerCustom_AddCharFilter(t *testing.T) {
	type fields struct {
		name       NormalizerName
		charFilter []CharFilterName
		filter     []FilterName
	}
	type args struct {
		charFilter []CharFilterName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *NormalizerCustom
	}{
		{
			args: args{charFilter: []CharFilterName{CharFilterNameHtmlStrip}},
			want: &NormalizerCustom{charFilter: []CharFilterName{CharFilterNameHtmlStrip}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizer := &NormalizerCustom{
				name:       tt.fields.name,
				charFilter: tt.fields.charFilter,
				filter:     tt.fields.filter,
			}
			if got := normalizer.AddCharFilter(tt.args.charFilter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCharFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizerCustom_AddFilter(t *testing.T) {
	type fields struct {
		name       NormalizerName
		charFilter []CharFilterName
		filter     []FilterName
	}
	type args struct {
		filter []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *NormalizerCustom
	}{
		{
			args: args{filter: []FilterName{FilterNameArabic}},
			want: &NormalizerCustom{filter: []FilterName{FilterNameArabic}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizer := &NormalizerCustom{
				name:       tt.fields.name,
				charFilter: tt.fields.charFilter,
				filter:     tt.fields.filter,
			}
			if got := normalizer.AddFilter(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizerCustom_Name(t *testing.T) {
	type fields struct {
		name       NormalizerName
		charFilter []CharFilterName
		filter     []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		want   NormalizerName
	}{
		{
			fields: fields{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizer := &NormalizerCustom{
				name:       tt.fields.name,
				charFilter: tt.fields.charFilter,
				filter:     tt.fields.filter,
			}
			if got := normalizer.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizerCustom_Source(t *testing.T) {
	type fields struct {
		name       NormalizerName
		charFilter []CharFilterName
		filter     []FilterName
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
				"type": NormalizerTypeCustom,
			},
		},
		{
			name: "char_filter",
			fields: fields{charFilter: []CharFilterName{CharFilterNameHtmlStrip}},
			want: map[string]interface{}{
				"type": NormalizerTypeCustom,
				"char_filter": []CharFilterName{CharFilterNameHtmlStrip},
			},
		},
		{
			name: "filter",
			fields: fields{filter: []FilterName{FilterNameArabic}},
			want: map[string]interface{}{
				"type": NormalizerTypeCustom,
				"filter": []FilterName{FilterNameArabic},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizer := &NormalizerCustom{
				name:       tt.fields.name,
				charFilter: tt.fields.charFilter,
				filter:     tt.fields.filter,
			}
			got, err := normalizer.Source()
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

func TestNormalizerCustom_Type(t *testing.T) {
	type fields struct {
		name       NormalizerName
		charFilter []CharFilterName
		filter     []FilterName
	}
	tests := []struct {
		name   string
		fields fields
		want   NormalizerType
	}{
		{
			want: NormalizerTypeCustom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizer := &NormalizerCustom{
				name:       tt.fields.name,
				charFilter: tt.fields.charFilter,
				filter:     tt.fields.filter,
			}
			if got := normalizer.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
