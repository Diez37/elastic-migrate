package settings

import (
	"reflect"
	"testing"
)

func TestNewThreshold(t *testing.T) {
	tests := []struct {
		name string
		want *Threshold
	}{
		{
			want: &Threshold{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewThreshold(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreshold_SetFetch(t *testing.T) {
	type fields struct {
		fetch *LogLevel
		query *LogLevel
	}
	type args struct {
		fetch *LogLevel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Threshold
	}{
		{
			args: args{fetch: &LogLevel{}},
			want: &Threshold{fetch: &LogLevel{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := &Threshold{
				fetch: tt.fields.fetch,
				query: tt.fields.query,
			}
			if got := threshold.SetFetch(tt.args.fetch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreshold_SetQuery(t *testing.T) {
	type fields struct {
		fetch *LogLevel
		query *LogLevel
	}
	type args struct {
		query *LogLevel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Threshold
	}{
		{
			args: args{query: &LogLevel{}},
			want: &Threshold{query: &LogLevel{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := &Threshold{
				fetch: tt.fields.fetch,
				query: tt.fields.query,
			}
			if got := threshold.SetQuery(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreshold_Source(t *testing.T) {
	logLevelSrc, _ := (&LogLevel{}).Source()

	type fields struct {
		fetch *LogLevel
		query *LogLevel
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
			name: "fetch",
			fields: fields{fetch: &LogLevel{}},
			want: map[string]interface{}{
				"fetch": logLevelSrc,
			},
		},
		{
			name: "query",
			fields: fields{query: &LogLevel{}},
			want: map[string]interface{}{
				"query": logLevelSrc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := &Threshold{
				fetch: tt.fields.fetch,
				query: tt.fields.query,
			}
			got, err := threshold.Source()
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
