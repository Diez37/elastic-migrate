package settings

import (
	"reflect"
	"testing"
)

func TestNewTranslog(t *testing.T) {
	tests := []struct {
		name string
		want *Translog
	}{
		{
			want: &Translog{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTranslog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTranslog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslog_SetDurability(t *testing.T) {
	translogDurabilityRequest := TranslogDurabilityRequest
	translogDurabilityAsync := TranslogDurabilityAsync

	type fields struct {
		syncInterval       *Interval
		durability         *TranslogDurability
		flushThresholdSize *Size
	}
	type args struct {
		durability TranslogDurability
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Translog
	}{
		{
			name: "translog_durability_request",
			args: args{durability: TranslogDurabilityRequest},
			want: &Translog{durability: &translogDurabilityRequest},
		},
		{
			name: "translog_durability_async",
			args: args{durability: TranslogDurabilityAsync},
			want: &Translog{durability: &translogDurabilityAsync},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translog := &Translog{
				syncInterval:       tt.fields.syncInterval,
				durability:         tt.fields.durability,
				flushThresholdSize: tt.fields.flushThresholdSize,
			}
			if got := translog.SetDurability(tt.args.durability); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDurability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslog_SetFlushThresholdSize(t *testing.T) {
	type fields struct {
		syncInterval       *Interval
		durability         *TranslogDurability
		flushThresholdSize *Size
	}
	type args struct {
		flushThresholdSize *Size
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Translog
	}{
		{
			args: args{flushThresholdSize: &Size{}},
			want: &Translog{flushThresholdSize: &Size{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translog := &Translog{
				syncInterval:       tt.fields.syncInterval,
				durability:         tt.fields.durability,
				flushThresholdSize: tt.fields.flushThresholdSize,
			}
			if got := translog.SetFlushThresholdSize(tt.args.flushThresholdSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFlushThresholdSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslog_SetSyncInterval(t *testing.T) {
	type fields struct {
		syncInterval       *Interval
		durability         *TranslogDurability
		flushThresholdSize *Size
	}
	type args struct {
		syncInterval *Interval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Translog
	}{
		{
			args: args{syncInterval: &Interval{}},
			want: &Translog{syncInterval: &Interval{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translog := &Translog{
				syncInterval:       tt.fields.syncInterval,
				durability:         tt.fields.durability,
				flushThresholdSize: tt.fields.flushThresholdSize,
			}
			if got := translog.SetSyncInterval(tt.args.syncInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSyncInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslog_Source(t *testing.T) {
	translogDurabilityRequest := TranslogDurabilityRequest
	sizeSrc, _ := (&Size{value: "1mb"}).Source()
	intervalSrc, _ := (&Interval{value: "1s"}).Source()

	type fields struct {
		syncInterval       *Interval
		durability         *TranslogDurability
		flushThresholdSize *Size
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "empty",
			want: map[string]interface{}{},
		},
		{
			name: "durability",
			fields: fields{durability: &translogDurabilityRequest},
			want: map[string]interface{}{
				"durability": TranslogDurabilityRequest,
			},
		},
		{
			name: "sync_interval",
			fields: fields{syncInterval: &Interval{value: "1s"}},
			want: map[string]interface{}{
				"sync_interval": intervalSrc,
			},
		},
		{
			name: "flush_threshold_size",
			fields: fields{flushThresholdSize: &Size{value: "1mb"}},
			want: map[string]interface{}{
				"flush_threshold_size": sizeSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translog := &Translog{
				syncInterval:       tt.fields.syncInterval,
				durability:         tt.fields.durability,
				flushThresholdSize: tt.fields.flushThresholdSize,
			}
			got, err := translog.Source()
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
