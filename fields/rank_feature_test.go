package fields

import (
    "reflect"
    "testing"
)

func TestNewRankFeature(t *testing.T) {
    tests := []struct {
        name string
        want *RankFeature
    }{
        {
            name: "new",
            want: &RankFeature{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewRankFeature(); !reflect.DeepEqual(got, tt.want) {
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
            want:   &RankFeature{positiveScoreImpact: &testTrue},
        },
        {
            name:   "false",
            fields: fields{},
            args:   args{positiveScoreImpact: false},
            want:   &RankFeature{positiveScoreImpact: &testFalse},
        },
        {
            name:   "change",
            fields: fields{positiveScoreImpact: &testTrue},
            args:   args{positiveScoreImpact: false},
            want:   &RankFeature{positiveScoreImpact: &testFalse},
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
            name:    "empty",
            fields:  fields{},
            want:    map[string]interface{}{
                "type": TypeRankFeature,
            },
            wantErr: false,
        },
        {
            name:    "positiveScoreImpact",
            fields:  fields{positiveScoreImpact: &testTrue},
            want:    map[string]interface{}{
                "type": TypeRankFeature,
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