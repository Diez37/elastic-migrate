package settings

import (
	"reflect"
	"testing"
)

func TestNewRetentionLease(t *testing.T) {
	type args struct {
		period *Interval
	}
	tests := []struct {
		name string
		args args
		want *RetentionLease
	}{
		{
			args: args{period: &Interval{value: "1s"}},
			want: &RetentionLease{period: &Interval{value: "1s"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRetentionLease(tt.args.period); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRetentionLease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetentionLease_Source(t *testing.T) {
	type fields struct {
		period *Interval
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{period: &Interval{value: "1s"}},
			want: map[string]interface{}{
				"period": "1s",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retentionLease := &RetentionLease{
				period: tt.fields.period,
			}
			got, err := retentionLease.Source()
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
