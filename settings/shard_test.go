package settings

import (
	"reflect"
	"testing"
)

func TestNewShard(t *testing.T) {
	type args struct {
		checkOnStartup CheckOnStartup
	}
	tests := []struct {
		name string
		args args
		want *Shard
	}{
		{
			name: "check_on_startup_false",
			args: args{checkOnStartup: CheckOnStartupFalse},
			want: &Shard{checkOnStartup: CheckOnStartupFalse},
		},
		{
			name: "check_on_startup_true",
			args: args{checkOnStartup: CheckOnStartupTrue},
			want: &Shard{checkOnStartup: CheckOnStartupTrue},
		},
		{
			name: "check_on_startup_checksum",
			args: args{checkOnStartup: CheckOnStartupChecksum},
			want: &Shard{checkOnStartup: CheckOnStartupChecksum},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShard(tt.args.checkOnStartup); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShard_Source(t *testing.T) {
	type fields struct {
		checkOnStartup CheckOnStartup
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "check_on_startup_false",
			fields: fields{checkOnStartup: CheckOnStartupFalse},
			want: map[string]interface{}{
				"check_on_startup": CheckOnStartupFalse,
			},
		},
		{
			name: "check_on_startup_true",
			fields: fields{checkOnStartup: CheckOnStartupTrue},
			want: map[string]interface{}{
				"check_on_startup": CheckOnStartupTrue,
			},
		},
		{
			name: "check_on_startup_checksum",
			fields: fields{checkOnStartup: CheckOnStartupChecksum},
			want: map[string]interface{}{
				"check_on_startup": CheckOnStartupChecksum,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shard := &Shard{
				checkOnStartup: tt.fields.checkOnStartup,
			}
			got, err := shard.Source()
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
