package settings

import (
	"reflect"
	"testing"
)

func TestLimit_Source(t *testing.T) {
	type fields struct {
		limit uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{limit: uint32(1)},
			want: map[string]interface{}{
				"limit": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit := &Limit{
				limit: tt.fields.limit,
			}
			got, err := limit.Source()
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

func TestNewLimit(t *testing.T) {
	type args struct {
		limit uint32
	}
	tests := []struct {
		name string
		args args
		want *Limit
	}{
		{
			args: args{limit: 1},
			want: &Limit{limit: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLimit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
