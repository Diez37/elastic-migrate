package settings

import (
	"reflect"
	"testing"
)

func TestNewSlowlog(t *testing.T) {
	tests := []struct {
		name string
		want *Slowlog
	}{
		{
			want: &Slowlog{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlowlog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlowlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlowlog_SetLevel(t *testing.T) {
	slowLogWarn := SlowLogWarn
	slowLogTrace := SlowLogTrace
	slowLogDebug := SlowLogDebug
	slowLogInfo := SlowLogInfo

	type fields struct {
		reformat  *bool
		source    *uint32
		level     *SlowLogLevel
		threshold *Threshold
	}
	type args struct {
		level SlowLogLevel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Slowlog
	}{
		{
			name: "slow_log_warn",
			args: args{level: SlowLogWarn},
			want: &Slowlog{level: &slowLogWarn},
		},
		{
			name: "slow_log_trace",
			args: args{level: SlowLogTrace},
			want: &Slowlog{level: &slowLogTrace},
		},
		{
			name: "slow_log_debug",
			args: args{level: SlowLogDebug},
			want: &Slowlog{level: &slowLogDebug},
		},
		{
			name: "slow_log_info",
			args: args{level: SlowLogInfo},
			want: &Slowlog{level: &slowLogInfo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowlog := &Slowlog{
				reformat:  tt.fields.reformat,
				source:    tt.fields.source,
				level:     tt.fields.level,
				threshold: tt.fields.threshold,
			}
			if got := slowlog.SetLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlowlog_SetReformat(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		reformat  *bool
		source    *uint32
		level     *SlowLogLevel
		threshold *Threshold
	}
	type args struct {
		reformat bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Slowlog
	}{
		{
			name: "set",
			args: args{reformat: true},
			want: &Slowlog{reformat: &boolTrue},
		},
		{
			name: "change",
			fields: fields{reformat: &boolTrue},
			args: args{reformat: false},
			want: &Slowlog{reformat: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowlog := &Slowlog{
				reformat:  tt.fields.reformat,
				source:    tt.fields.source,
				level:     tt.fields.level,
				threshold: tt.fields.threshold,
			}
			if got := slowlog.SetReformat(tt.args.reformat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReformat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlowlog_SetSource(t *testing.T) {
	numberSet := uint32(1)
	numberChange := uint32(2)

	type fields struct {
		reformat  *bool
		source    *uint32
		level     *SlowLogLevel
		threshold *Threshold
	}
	type args struct {
		source uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Slowlog
	}{
		{
			name: "set",
			args: args{source: 1},
			want: &Slowlog{source: &numberSet},
		},
		{
			name: "change",
			fields: fields{source: &numberSet},
			args: args{source: 2},
			want: &Slowlog{source: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowlog := &Slowlog{
				reformat:  tt.fields.reformat,
				source:    tt.fields.source,
				level:     tt.fields.level,
				threshold: tt.fields.threshold,
			}
			if got := slowlog.SetSource(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlowlog_SetThreshold(t *testing.T) {
	type fields struct {
		reformat  *bool
		source    *uint32
		level     *SlowLogLevel
		threshold *Threshold
	}
	type args struct {
		threshold *Threshold
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Slowlog
	}{
		{
			args: args{threshold: &Threshold{}},
			want: &Slowlog{threshold: &Threshold{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowlog := &Slowlog{
				reformat:  tt.fields.reformat,
				source:    tt.fields.source,
				level:     tt.fields.level,
				threshold: tt.fields.threshold,
			}
			if got := slowlog.SetThreshold(tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlowlog_Source(t *testing.T) {
	slowLogWarn := SlowLogWarn
	boolTrue := true
	number := uint32(1)
	thresholdSrc, _ := (&Threshold{}).Source()

	type fields struct {
		reformat  *bool
		source    *uint32
		level     *SlowLogLevel
		threshold *Threshold
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
			name: "reformat",
			fields: fields{reformat: &boolTrue},
			want: map[string]interface{}{
				"reformat": true,
			},
		},
		{
			name: "source",
			fields: fields{source: &number},
			want: map[string]interface{}{
				"source": uint32(1),
			},
		},
		{
			name: "level",
			fields: fields{level: &slowLogWarn},
			want: map[string]interface{}{
				"level": SlowLogWarn,
			},
		},
		{
			name: "threshold",
			fields: fields{threshold: &Threshold{}},
			want: map[string]interface{}{
				"threshold": thresholdSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowlog := &Slowlog{
				reformat:  tt.fields.reformat,
				source:    tt.fields.source,
				level:     tt.fields.level,
				threshold: tt.fields.threshold,
			}
			got, err := slowlog.Source()
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
