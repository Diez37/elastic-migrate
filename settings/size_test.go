package settings

import (
	"reflect"
	"testing"
)

func TestNewSize(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *Size
	}{
		{
			args: args{value: "1mb"},
			want: &Size{value: "1mb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSize(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize_Source(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{value: "1mb"},
			want: "1mb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := &Size{
				value: tt.fields.value,
			}
			got, err := size.Source()
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
