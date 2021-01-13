package settings

import (
	"reflect"
	"testing"
)

func TestNewWrite(t *testing.T) {
	type args struct {
		waitForActiveShards uint32
	}
	tests := []struct {
		name string
		args args
		want *Write
	}{
		{
			args: args{waitForActiveShards: 1},
			want: &Write{waitForActiveShards: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWrite(tt.args.waitForActiveShards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrite_Source(t *testing.T) {
	type fields struct {
		waitForActiveShards uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{waitForActiveShards: 1},
			want: map[string]interface{}{
				"wait_for_active_shards": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			write := &Write{
				waitForActiveShards: tt.fields.waitForActiveShards,
			}
			got, err := write.Source()
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
