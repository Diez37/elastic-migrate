package settings

import (
	"reflect"
	"testing"
)

func TestIdle_Source(t *testing.T) {
	type fields struct {
		after *Interval
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{after: &Interval{value: "1s"}},
			want: map[string]interface{}{
				"after": "1s",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idle := &Idle{
				after: tt.fields.after,
			}
			got, err := idle.Source()
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

func TestNewIdle(t *testing.T) {
	type args struct {
		after *Interval
	}
	tests := []struct {
		name string
		args args
		want *Idle
	}{
		{
			args: args{after: &Interval{value: "1s"}},
			want: &Idle{after: &Interval{value: "1s"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdle(tt.args.after); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdle() = %v, want %v", got, tt.want)
			}
		})
	}
}
