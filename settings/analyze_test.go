package settings

import (
	"reflect"
	"testing"
)

func TestAnalyze_Source(t *testing.T) {
	type fields struct {
		maxTokenCount uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{maxTokenCount: 2},
			want: map[string]interface{}{
				"max_token_count": uint32(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyze := &Analyze{
				maxTokenCount: tt.fields.maxTokenCount,
			}
			got, err := analyze.Source()
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

func TestNewAnalyze(t *testing.T) {
	maxTokenCount := uint32(1)

	type args struct {
		maxTokenCount uint32
	}
	tests := []struct {
		name string
		args args
		want *Analyze
	}{
		{
			args: args{maxTokenCount: maxTokenCount},
			want: &Analyze{maxTokenCount: maxTokenCount},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnalyze(tt.args.maxTokenCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
