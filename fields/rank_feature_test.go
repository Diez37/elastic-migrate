package fields

import (
	"reflect"
	"testing"
)

func TestNewRankFeature(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *RankFeature
	}{
		{
			args: args{name: "test"},
			want: &RankFeature{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRankFeature(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRankFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankFeature_GetType(t *testing.T) {
	type fields struct {
		positiveScoreImpact *bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "type rank feature",
			fields: fields{},
			want:   TypeRankFeature,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &RankFeature{
				positiveScoreImpact: tt.fields.positiveScoreImpact,
			}
			if got := field.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankFeature_SetPositiveScoreImpact(t *testing.T) {
	boolTrue := true
	boolFalse := false

	type fields struct {
		positiveScoreImpact *bool
	}
	type args struct {
		positiveScoreImpact bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RankFeature
	}{
		{
			name:   "true",
			fields: fields{},
			args:   args{positiveScoreImpact: true},
			want:   &RankFeature{positiveScoreImpact: &boolTrue},
		},
		{
			name:   "false",
			fields: fields{},
			args:   args{positiveScoreImpact: false},
			want:   &RankFeature{positiveScoreImpact: &boolFalse},
		},
		{
			name:   "change",
			fields: fields{positiveScoreImpact: &boolTrue},
			args:   args{positiveScoreImpact: false},
			want:   &RankFeature{positiveScoreImpact: &boolFalse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &RankFeature{
				positiveScoreImpact: tt.fields.positiveScoreImpact,
			}
			if got := field.SetPositiveScoreImpact(tt.args.positiveScoreImpact); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPositiveScoreImpact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankFeature_Source(t *testing.T) {
	boolTrue := true

	type fields struct {
		positiveScoreImpact *bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name:   "empty",
			fields: fields{},
			want: map[string]interface{}{
				"type": TypeRankFeature,
			},
			wantErr: false,
		},
		{
			name:   "positiveScoreImpact",
			fields: fields{positiveScoreImpact: &boolTrue},
			want: map[string]interface{}{
				"type":                  TypeRankFeature,
				"positive_score_impact": true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &RankFeature{
				positiveScoreImpact: tt.fields.positiveScoreImpact,
			}
			got, err := field.Source()
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
