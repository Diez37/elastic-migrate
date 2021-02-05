package fields

import (
	"reflect"
	"testing"
)

func TestField_Source(t *testing.T) {
	boolFalse := false

	type fields struct {
		name  FieldName
		field Fielder
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "alias",
			fields: fields{
				name:  "product",
				field: &Alias{path: "goods"},
			},
			want: map[string]interface{}{
				"type": TypeAlias,
				"path": "goods",
			},
			wantErr: false,
		},
		{
			name: "binary",
			fields: fields{
				name:  "blob",
				field: &Binary{docValues: &boolFalse},
			},
			want: map[string]interface{}{
				"type":       TypeBinary,
				"doc_values": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Field{
				name:  tt.fields.name,
				field: tt.fields.field,
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

func TestNewField(t *testing.T) {
	type args struct {
		name  string
		field Fielder
	}
	tests := []struct {
		name string
		args args
		want *Field
	}{
		{
			name: "alias",
			args: args{
				name:  "product",
				field: &Alias{path: "goods"},
			},
			want: &Field{
				name:  "product",
				field: &Alias{path: "goods"},
			},
		},
		{
			name: "binary",
			args: args{
				name:  "blob",
				field: &Binary{},
			},
			want: &Field{
				name:  "blob",
				field: &Binary{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewField(tt.args.name, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewField() = %v, want %v", got, tt.want)
			}
		})
	}
}
