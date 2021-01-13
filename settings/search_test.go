package settings

import (
	"reflect"
	"testing"
)

func TestNewSearch(t *testing.T) {
	tests := []struct {
		name string
		want *Search
	}{
		{
			want: &Search{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_SetIdle(t *testing.T) {
	type fields struct {
		throttled *bool
		slowlog   *Slowlog
		idle      *Idle
	}
	type args struct {
		idle *Idle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Search
	}{
		{
			name: "set",
			args: args{idle: &Idle{}},
			want: &Search{idle: &Idle{}},
		},
		{
			name: "change",
			fields: fields{idle: &Idle{}},
			args: args{idle: &Idle{}},
			want: &Search{idle: &Idle{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search := &Search{
				throttled: tt.fields.throttled,
				slowlog:   tt.fields.slowlog,
				idle:      tt.fields.idle,
			}
			if got := search.SetIdle(tt.args.idle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIdle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_SetSlowlog(t *testing.T) {
	type fields struct {
		throttled *bool
		slowlog   *Slowlog
		idle      *Idle
	}
	type args struct {
		slowlog *Slowlog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Search
	}{
		{
			name: "set",
			args: args{slowlog: &Slowlog{}},
			want: &Search{slowlog: &Slowlog{}},
		},
		{
			name: "change",
			fields: fields{slowlog: &Slowlog{threshold: &Threshold{}}},
			args: args{slowlog: &Slowlog{}},
			want: &Search{slowlog: &Slowlog{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search := &Search{
				throttled: tt.fields.throttled,
				slowlog:   tt.fields.slowlog,
				idle:      tt.fields.idle,
			}
			if got := search.SetSlowlog(tt.args.slowlog); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSlowlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_SetThrottled(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		throttled *bool
		slowlog   *Slowlog
		idle      *Idle
	}
	type args struct {
		throttled bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Search
	}{
		{
			name: "set",
			args: args{throttled: true},
			want: &Search{throttled: &boolTrue},
		},
		{
			name: "change",
			fields: fields{throttled: &boolTrue},
			args: args{throttled: false},
			want: &Search{throttled: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search := &Search{
				throttled: tt.fields.throttled,
				slowlog:   tt.fields.slowlog,
				idle:      tt.fields.idle,
			}
			if got := search.SetThrottled(tt.args.throttled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetThrottled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch_Source(t *testing.T) {
	boolTrue := true
	slowlogSrc, _ := (&Slowlog{}).Source()
	idleSrc, _ := (&Idle{&Interval{value: "1s"}}).Source()

	type fields struct {
		throttled *bool
		slowlog   *Slowlog
		idle      *Idle
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
			name: "throttled",
			fields: fields{throttled: &boolTrue},
			want: map[string]interface{}{
				"throttled": true,
			},
		},
		{
			name: "slowlog",
			fields: fields{slowlog: &Slowlog{}},
			want: map[string]interface{}{
				"slowlog": slowlogSrc,
			},
		},
		{
			name: "idle",
			fields: fields{idle: &Idle{&Interval{value: "1s"}}},
			want: map[string]interface{}{
				"idle": idleSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search := &Search{
				throttled: tt.fields.throttled,
				slowlog:   tt.fields.slowlog,
				idle:      tt.fields.idle,
			}
			got, err := search.Source()
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
