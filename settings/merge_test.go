package settings

import (
	"reflect"
	"testing"
)

func TestMerge_SetPolicy(t *testing.T) {
	type fields struct {
		scheduler *Scheduler
		policy    *Policy
	}
	type args struct {
		policy *Policy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Merge
	}{
		{
			name: "set",
			args: args{policy: NewPolicy()},
			want: &Merge{policy: NewPolicy()},
		},
		{
			name: "change",
			fields: fields{policy: NewPolicy()},
			args: args{policy: NewPolicy().SetMaxMergeAtOnce(1)},
			want: &Merge{policy: NewPolicy().SetMaxMergeAtOnce(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge := &Merge{
				scheduler: tt.fields.scheduler,
				policy:    tt.fields.policy,
			}
			if got := merge.SetPolicy(tt.args.policy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge_SetScheduler(t *testing.T) {
	type fields struct {
		scheduler *Scheduler
		policy    *Policy
	}
	type args struct {
		scheduler *Scheduler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Merge
	}{
		{
			name: "set",
			args: args{scheduler: NewScheduler()},
			want: &Merge{scheduler: NewScheduler()},
		},
		{
			name: "change",
			fields: fields{scheduler: NewScheduler()},
			args: args{scheduler: NewScheduler().SetMaxMergeCount(1)},
			want: &Merge{scheduler: NewScheduler().SetMaxMergeCount(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge := &Merge{
				scheduler: tt.fields.scheduler,
				policy:    tt.fields.policy,
			}
			if got := merge.SetScheduler(tt.args.scheduler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScheduler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge_Source(t *testing.T) {
	schedulerSrc, _ := (&Scheduler{}).Source()
	policySrc, _ := (&Policy{}).Source()

	type fields struct {
		scheduler *Scheduler
		policy    *Policy
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
			name: "scheduler",
			fields: fields{scheduler: &Scheduler{}},
			want: map[string]interface{}{
				"scheduler": schedulerSrc,
			},
		},
		{
			name: "policy",
			fields: fields{policy: &Policy{}},
			want: map[string]interface{}{
				"policy": policySrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge := &Merge{
				scheduler: tt.fields.scheduler,
				policy:    tt.fields.policy,
			}
			got, err := merge.Source()
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

func TestNewMerge(t *testing.T) {
	tests := []struct {
		name string
		want *Merge
	}{
		{
			want: &Merge{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMerge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
