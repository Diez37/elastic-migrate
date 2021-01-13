package settings

import (
	"reflect"
	"testing"
)

func TestNewScheduler(t *testing.T) {
	tests := []struct {
		name string
		want *Scheduler
	}{
		{
			want: &Scheduler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheduler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_SetAutoThrottle(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		autoThrottle   *bool
		maxMergeCount  *uint32
		maxThreadCount *uint32
	}
	type args struct {
		autoThrottle bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Scheduler
	}{
		{
			name: "set",
			args: args{autoThrottle: true},
			want: &Scheduler{autoThrottle: &boolTrue},
		},
		{
			name: "change",
			fields: fields{autoThrottle: &boolTrue},
			args: args{autoThrottle: false},
			want: &Scheduler{autoThrottle: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduler := &Scheduler{
				autoThrottle:   tt.fields.autoThrottle,
				maxMergeCount:  tt.fields.maxMergeCount,
				maxThreadCount: tt.fields.maxThreadCount,
			}
			if got := scheduler.SetAutoThrottle(tt.args.autoThrottle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAutoThrottle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_SetMaxMergeCount(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		autoThrottle   *bool
		maxMergeCount  *uint32
		maxThreadCount *uint32
	}
	type args struct {
		maxMergeCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Scheduler
	}{
		{
			name: "set",
			args: args{maxMergeCount: 1},
			want: &Scheduler{maxMergeCount: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxMergeCount: &numberSet},
			args: args{maxMergeCount: 2},
			want: &Scheduler{maxMergeCount: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduler := &Scheduler{
				autoThrottle:   tt.fields.autoThrottle,
				maxMergeCount:  tt.fields.maxMergeCount,
				maxThreadCount: tt.fields.maxThreadCount,
			}
			if got := scheduler.SetMaxMergeCount(tt.args.maxMergeCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxMergeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_SetMaxThreadCount(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		autoThrottle   *bool
		maxMergeCount  *uint32
		maxThreadCount *uint32
	}
	type args struct {
		maxThreadCount uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Scheduler
	}{
		{
			name: "set",
			args: args{maxThreadCount: 1},
			want: &Scheduler{maxThreadCount: &numberSet},
		},
		{
			name: "change",
			fields: fields{maxThreadCount: &numberSet},
			args: args{maxThreadCount: 2},
			want: &Scheduler{maxThreadCount: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduler := &Scheduler{
				autoThrottle:   tt.fields.autoThrottle,
				maxMergeCount:  tt.fields.maxMergeCount,
				maxThreadCount: tt.fields.maxThreadCount,
			}
			if got := scheduler.SetMaxThreadCount(tt.args.maxThreadCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMaxThreadCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_Source(t *testing.T) {
	number := uint32(1)
	boolTrue := true

	type fields struct {
		autoThrottle   *bool
		maxMergeCount  *uint32
		maxThreadCount *uint32
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
			name: "auto_throttle",
			fields: fields{autoThrottle: &boolTrue},
			want: map[string]interface{}{
				"auto_throttle": true,
			},
		},
		{
			name: "max_merge_count",
			fields: fields{maxMergeCount: &number},
			want: map[string]interface{}{
				"max_merge_count": uint32(1),
			},
		},
		{
			name: "max_thread_count",
			fields: fields{maxThreadCount: &number},
			want: map[string]interface{}{
				"max_thread_count": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduler := &Scheduler{
				autoThrottle:   tt.fields.autoThrottle,
				maxMergeCount:  tt.fields.maxMergeCount,
				maxThreadCount: tt.fields.maxThreadCount,
			}
			got, err := scheduler.Source()
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
