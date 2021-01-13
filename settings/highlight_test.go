package settings

import (
	"reflect"
	"testing"
)

func TestHighlight_Source(t *testing.T) {
	type fields struct {
		maxAnalyzedOffset uint32
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			fields: fields{maxAnalyzedOffset: 1},
			want: map[string]interface{}{
				"max_analyzed_offset": uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			highlight := &Highlight{
				maxAnalyzedOffset: tt.fields.maxAnalyzedOffset,
			}
			got, err := highlight.Source()
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

func TestNewHighlight(t *testing.T) {
	type args struct {
		maxAnalyzedOffset uint32
	}
	tests := []struct {
		name string
		args args
		want *Highlight
	}{
		{
			args: args{maxAnalyzedOffset: 1},
			want: &Highlight{maxAnalyzedOffset: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHighlight(tt.args.maxAnalyzedOffset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHighlight() = %v, want %v", got, tt.want)
			}
		})
	}
}
