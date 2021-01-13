package settings

import (
	"reflect"
	"testing"
)

func TestNewRetention(t *testing.T) {
	type args struct {
		operations uint32
	}
	tests := []struct {
		name string
		args args
		want *Retention
	}{
		{
			args: args{operations: 1},
			want: &Retention{operations: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRetention(tt.args.operations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRetention() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetention_Source(t *testing.T) {
	type fields struct {
		operations uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{operations: 1},
			want: map[string]interface{}{
				"operations": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retention := &Retention{
				operations: tt.fields.operations,
			}
			got, err := retention.Source()
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
