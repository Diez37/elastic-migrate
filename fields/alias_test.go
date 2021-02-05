package fields

import (
	"reflect"
	"testing"
)

func TestAlias_GetType(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "type alias",
			fields: fields{path: "alis"},
			want:   TypeAlias,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Alias{
				path: tt.fields.path,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlias_Source(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "no error",
			fields: fields{path: "test"},
			want: map[string]interface{}{
				"type": TypeAlias,
				"path": "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &Alias{
				path: tt.fields.path,
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

func TestNewAlias(t *testing.T) {
	type args struct {
		path string
		name string
	}
	tests := []struct {
		name string
		args args
		want *Alias
	}{
		{
			name: "new",
			args: args{path: "test", name: "test"},
			want: &Alias{path: "test", name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlias(tt.args.name, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}
