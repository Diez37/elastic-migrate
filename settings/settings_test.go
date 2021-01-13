package settings

import (
	"reflect"
	"testing"
)

func TestNewSettings(t *testing.T) {
	type args struct {
		index *Index
	}
	tests := []struct {
		name string
		args args
		want *Settings
	}{
		{
			want: &Settings{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSettings(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettings_Source(t *testing.T) {
	indexSrc, _ := (&Index{}).Source()

	type fields struct {
		index *Index
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{index: &Index{}},
			want: map[string]interface{}{
				"index": indexSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := &Settings{
				index: tt.fields.index,
			}
			got, err := settings.Source()
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
