package settings

import (
	"reflect"
	"testing"
)

func TestInterval_Source(t *testing.T) {
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
			fields: fields{value: "1s"},
			want: "1s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interval := &Interval{
				value: tt.fields.value,
			}
			got, err := interval.Source()
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

func TestNewInterval(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *Interval
	}{
		{
			args: args{value: "1s"},
			want: &Interval{value: "1s"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterval(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
