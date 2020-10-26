package fields

import (
    "reflect"
    "testing"
)

func TestNewRankFeatures(t *testing.T) {
    tests := []struct {
        name string
        want *RankFeatures
    }{
        {
            name: "new",
            want: &RankFeatures{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewRankFeatures(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewRankFeatures() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRankFeatures_GetType(t *testing.T) {
    tests := []struct {
        name string
        want Type
    }{
        {
            name: "type rank features",
            want: TypeRankFeatures,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &RankFeatures{}
            if got := field.GetType(); got != tt.want {
                t.Errorf("GetType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRankFeatures_Source(t *testing.T) {
    tests := []struct {
        name    string
        want    interface{}
        wantErr bool
    }{
        {
            name:    "source",
            want:    map[string]interface{}{
                "type": TypeRankFeatures,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            field := &RankFeatures{}
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