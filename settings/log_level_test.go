package settings

import (
	"reflect"
	"testing"
)

func TestLogLevel_SetDebug(t *testing.T) {
	numberSet := int8(1)
	numberChange := int8(2)

	type fields struct {
		warn  *int8
		trace *int8
		debug *int8
		info  *int8
	}
	type args struct {
		debug int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogLevel
	}{
		{
			name: "set",
			args: args{debug: 1},
			want: &LogLevel{debug: &numberSet},
		},
		{
			name: "change",
			fields: fields{debug: &numberSet},
			args: args{debug: 2},
			want: &LogLevel{debug: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel := &LogLevel{
				warn:  tt.fields.warn,
				trace: tt.fields.trace,
				debug: tt.fields.debug,
				info:  tt.fields.info,
			}
			if got := logLevel.SetDebug(tt.args.debug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel_SetInfo(t *testing.T) {
	numberSet := int8(1)
	numberChange := int8(2)

	type fields struct {
		warn  *int8
		trace *int8
		debug *int8
		info  *int8
	}
	type args struct {
		info int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogLevel
	}{
		{
			name: "set",
			args: args{info: 1},
			want: &LogLevel{info: &numberSet},
		},
		{
			name: "change",
			fields: fields{info: &numberSet},
			args: args{info: 2},
			want: &LogLevel{info: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel := &LogLevel{
				warn:  tt.fields.warn,
				trace: tt.fields.trace,
				debug: tt.fields.debug,
				info:  tt.fields.info,
			}
			if got := logLevel.SetInfo(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel_SetTrace(t *testing.T) {
	numberSet := int8(1)
	numberChange := int8(2)

	type fields struct {
		warn  *int8
		trace *int8
		debug *int8
		info  *int8
	}
	type args struct {
		trace int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogLevel
	}{
		{
			name: "set",
			args: args{trace: 1},
			want: &LogLevel{trace: &numberSet},
		},
		{
			name: "change",
			fields: fields{trace: &numberSet},
			args: args{trace: 2},
			want: &LogLevel{trace: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel := &LogLevel{
				warn:  tt.fields.warn,
				trace: tt.fields.trace,
				debug: tt.fields.debug,
				info:  tt.fields.info,
			}
			if got := logLevel.SetTrace(tt.args.trace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTrace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel_SetWarn(t *testing.T) {
	numberSet := int8(1)
	numberChange := int8(2)

	type fields struct {
		warn  *int8
		trace *int8
		debug *int8
		info  *int8
	}
	type args struct {
		warn int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogLevel
	}{
		{
			name: "set",
			args: args{warn: 1},
			want: &LogLevel{warn: &numberSet},
		},
		{
			name: "change",
			fields: fields{warn: &numberSet},
			args: args{warn: 2},
			want: &LogLevel{warn: &numberChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel := &LogLevel{
				warn:  tt.fields.warn,
				trace: tt.fields.trace,
				debug: tt.fields.debug,
				info:  tt.fields.info,
			}
			if got := logLevel.SetWarn(tt.args.warn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWarn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel_Source(t *testing.T) {
	number := int8(1)

	type fields struct {
		warn  *int8
		trace *int8
		debug *int8
		info  *int8
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
			name: "warn",
			fields: fields{warn: &number},
			want: map[string]interface{}{
				"warn": int8(1),
			},
		},
		{
			name: "trace",
			fields: fields{trace: &number},
			want: map[string]interface{}{
				"trace": int8(1),
			},
		},
		{
			name: "debug",
			fields: fields{debug: &number},
			want: map[string]interface{}{
				"debug": int8(1),
			},
		},
		{
			name: "info",
			fields: fields{info: &number},
			want: map[string]interface{}{
				"info": int8(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel := &LogLevel{
				warn:  tt.fields.warn,
				trace: tt.fields.trace,
				debug: tt.fields.debug,
				info:  tt.fields.info,
			}
			got, err := logLevel.Source()
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

func TestNewLogLevel(t *testing.T) {
	tests := []struct {
		name string
		want *LogLevel
	}{
		{
			want: &LogLevel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
